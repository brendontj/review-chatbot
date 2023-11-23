package usecase 

import (
    "github.com/brendontj/review-chatbot/infrastructure/template"
    "github.com/brendontj/review-chatbot/core/gateway"
)


type StartWorkflowUseCase struct {
    WorkflowsTemplate gateway.Template
    Repo gateway.Repository
}

func NewStartWorkflowUseCase(workflowsTemplate gateway.Template, dbRepo gateway.Repository) *StartWorkflowUseCase {
    return &StartWorkflowUseCase{
		WorkflowsTemplate: workflowsTemplate,
		Repo: dbRepo,
	}
}

func (u *StartWorkflowUseCase) Execute(msg string) {
    workflow := u.WorkflowsTemplate.FindByType(msg)
    if workflow.Type == "" {
		w, err := u.Repo.GetLastWorkflowNonFinalizedWithSteps()
		if err != nil {
			return
		}

		if err := u.Repo.SaveStepAnswer(u, wtIdentifier); err != nil {
			return
		}
		return u.executeWorkflow(workflow, w.LastStepIndex + 1)
    }

	return u.executeWorkflow(workflow, 0) 
}

func (u *StartWorkflowUseCase) executeWorkflow(workflow template.Workflow, stepIndex int) {
	ok := workflow.Steps[stepIndex]
	if !ok {
		return u.generateReview(workflow)
	}

	return u.generateNextStep(workflow, stepIndex)
}

func (u *StartWorkflowUseCase) generateReview(workflow template.Workflow) {
 	u.Repo.GetAnswers()
	return u.Repo.SaveReview(workflow)
}

