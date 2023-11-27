package usecase

import (
	"github.com/brendontj/review-chatbot/core/gateway"
	"github.com/brendontj/review-chatbot/infrastructure/template"
	"github.com/google/uuid"
)

const (
	EmptyAnswer = ""
	AfterReviewMessage = "Thank you for your review. We really appreciate it."
)

type StartWorkflowUseCase struct {
	WorkflowsTemplate gateway.Template
	Repo              gateway.Repository
}

func NewStartWorkflowUseCase(workflowsTemplate gateway.Template, dbRepo gateway.Repository) *StartWorkflowUseCase {
	return &StartWorkflowUseCase{
		WorkflowsTemplate: workflowsTemplate,
		Repo:              dbRepo,
	}
}

func (u *StartWorkflowUseCase) Execute(msg string) string {
	workflowTemplate := u.WorkflowsTemplate.FindByType(msg)

	if workflowTemplate.Type == "" {
		w, err := u.Repo.GetLastWorkflowNonFinalizedWithSteps()
		if err != nil {
			return EmptyAnswer
		}

		if err := u.Repo.SaveStepAnswer(w.StepWithoutAnswer().Id(), msg); err != nil {
			return EmptyAnswer
		}

		return u.executeWorkflow(
			w,
			u.WorkflowsTemplate.FindByType(w.Type()),
			w.StepWithoutAnswer().Order()+1,
			msg)
	}

	return u.executeWorkflow(nil, workflowTemplate, 0, msg)
}

func (u *StartWorkflowUseCase) executeWorkflow(
	wfFromDB gateway.WorkflowWithStepsModel,
	workflowTemplate template.WorkflowTemplate,
	stepIndex int,
	msg string) string {
	
	// If workflow is not started yet, start it
	if wfFromDB == nil {
		return u.startWorkflow(workflowTemplate)
	}

	// If workflow is finalized, return empty answer
	if !workflowTemplate.ExistStepWithOrderEqual(stepIndex) {
		return u.generateReview(wfFromDB, msg)
	}

	return u.generateNextStep(wfFromDB, workflowTemplate, stepIndex, msg)
}

func (u *StartWorkflowUseCase) startWorkflow(workflowTemplate template.WorkflowTemplate) string {
	if workflowTemplate.ExistStepWithOrderEqual(0) {
		workflowFromDB, err := u.Repo.SaveWorkflow(workflowTemplate.Type)
		if err != nil {
			return EmptyAnswer
		}

		if err := u.generateStep(workflowFromDB.Id(), 0); err != nil {
			return EmptyAnswer
		}
		return workflowTemplate.Steps[0].Text
	}
	return EmptyAnswer
}

func (u *StartWorkflowUseCase) generateNextStep(wfFromDB gateway.WorkflowWithStepsModel, workflowTemplate template.WorkflowTemplate, stepOrder int, msg string) string {
	for _, v := range workflowTemplate.Steps {
		if v.Order == stepOrder {
			if err := u.generateStep(wfFromDB.Id(), stepOrder); err != nil {
				return EmptyAnswer
			}
			return v.Text
		}
	}

	return EmptyAnswer
}

func (u *StartWorkflowUseCase) generateReview(wfFromDB gateway.WorkflowWithStepsModel, msg string) string {
	_ = u.Repo.SaveReview(wfFromDB, msg)
	_ = u.finalizeWorkflow(wfFromDB.Id())
	return AfterReviewMessage
}

func (u *StartWorkflowUseCase) generateStep(workflowID uuid.UUID, stepIndex int) error {
	return u.Repo.SaveStep(workflowID, stepIndex)
}

func (u *StartWorkflowUseCase) finalizeWorkflow(workflowID uuid.UUID) error {
	return u.Repo.SaveFinalizedWorkflow(workflowID)
}
