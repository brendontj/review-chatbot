package usecase

import (
	"fmt"

	"github.com/brendontj/review-chatbot/core/gateway"
	"github.com/brendontj/review-chatbot/infrastructure/template"
	"github.com/google/uuid"
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
	fmt.Println("Finding workflow template by type: ", msg)

	if workflowTemplate.Type == "" {
		fmt.Println("Workflow template not found")

		fmt.Println("Finding last workflow non finalized with steps")
		w, err := u.Repo.GetLastWorkflowNonFinalizedWithSteps()
		if err != nil {
			fmt.Println("Error getting last workflow non finalized with steps: ", err)
			return ""
		}

		fmt.Println("Saving step answer...")
		if err := u.Repo.SaveStepAnswer(w.StepWithoutAnswer().Id(), msg); err != nil {
			fmt.Println("Error saving step answer: ", err)
			return ""
		}

		return u.executeWorkflow(w, workflowTemplate, w.StepWithoutAnswer().Order()+1, msg)
	}

	return u.executeWorkflow(nil, workflowTemplate, 0, msg)
}

func (u *StartWorkflowUseCase) executeWorkflow(
	wfFromDB gateway.WorkflowWithStepsModel,
	workflowTemplate template.WorkflowTemplate,
	stepIndex int,
	msg string) string {

	if wfFromDB == nil {
		return u.startWorkflow(workflowTemplate)
	}

	if !workflowTemplate.ExistStepWithOrderEqual(stepIndex) {
		return u.generateReview(wfFromDB, msg)
	}

	return u.generateNextStep(wfFromDB, workflowTemplate, stepIndex, msg)
}

func (u *StartWorkflowUseCase) startWorkflow(workflowTemplate template.WorkflowTemplate) string {
	fmt.Println("Starting workflow...")
	if workflowTemplate.ExistStepWithOrderEqual(0) {
		fmt.Println("Workflow template has step with order equal 0")
		fmt.Println("Saving workflow...")
		workflowFromDB, err := u.Repo.SaveWorkflow(workflowTemplate.Type)
		if err != nil {
			fmt.Println("Error saving workflow: ", err)
			return ""
		}
		
		fmt.Println("Generating step...")
		if err := u.generateStep(workflowFromDB.Id(), 0); err != nil {
			return ""
		}
		return workflowTemplate.Steps[0].Text
	}
	return ""
}

func (u *StartWorkflowUseCase) generateNextStep(wfFromDB gateway.WorkflowWithStepsModel, workflowTemplate template.WorkflowTemplate, stepOrder int, msg string) string {
	fmt.Println("Saving step answer...")
	err := u.Repo.SaveStepAnswer(wfFromDB.StepWithoutAnswer().Id(), msg)
	if err != nil {
		return ""
	}

	fmt.Println("Generating next step...")	
	for _, v := range workflowTemplate.Steps {
		if v.Order == stepOrder {
			if err := u.generateStep(wfFromDB.Id(), stepOrder); err != nil {
				return ""
			}
			return v.Text
		}
	}

	return ""
}

func (u *StartWorkflowUseCase) generateReview(wfFromDB gateway.WorkflowWithStepsModel, msg string) string {
	fmt.Println("Generating review...")
	_ = u.Repo.SaveReview(wfFromDB, msg)
	return "Thank you for your review. We really appreciate it."
}

func (u *StartWorkflowUseCase) generateStep(workflowID uuid.UUID, stepIndex int) error {
	fmt.Println("Generating step...")
	return u.Repo.SaveStep(workflowID, stepIndex)
}
