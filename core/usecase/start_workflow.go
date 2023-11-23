package usecase 

import (
    "github.com/brendontj/review-chatbot/infrastructure/template"
    "github.com/brendontj/review-chatbot/core/gateway
)


type StartWorkflowUseCase struct {
    WorkflowTemplate template.WorkflowTemplate
    Repo gateway.Repository
}

func NewStartWorkflowUseCase() *StartWorkflowUseCase {
    return &StartWorkflowUseCase{}
}

func (u *StartWorkflowUseCase) Execute(wfIndentifier string) {
    workflow := u.WorkflowTemplate.GetWorkflow(wfIndentifier)
    if workflow.Type == "" {
        return
    }

    wf, err :=  u.Repo.GetLastWorkflowNonFinishedByType(workflow.Type)
    if err != nil {
        //Should start a new workflow 
        u.Repo.SaveWorkflow(newWorkflow)
    }
    
    //Start the next step of the workflow 

    nextStep := u.WorkflowTemplate.GetStep(wf.Type, wf.CurrentStep+1)
    if nextStep == "" {
        //Workflow is finished 
        return
    }

    u.Sender.Send(u.WorkflowTemplate.GetStep(nextStep)
    return
}
