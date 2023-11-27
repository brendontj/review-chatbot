package gateway

import (
	"github.com/google/uuid"
)

type Repository interface {
	GetLastWorkflowNonFinalizedWithSteps() (WorkflowWithStepsModel, error)
	SaveFinalizedWorkflow(workflowID uuid.UUID) error
	SaveStepAnswer(stepID uuid.UUID, answer string) error
	SaveWorkflow(workflowType string) (WorkflowWithStepsModel, error)
	SaveStep(workflowID uuid.UUID, stepOrder int) error
	SaveReview(WorkflowWithStepsModel, string) error
}

type StepModel interface {
	Id() uuid.UUID
	Order() int
	Answer() string
}

type WorkflowWithStepsModel interface {
	Id() uuid.UUID
	Type() string
	Steps() []StepModel
	StepWithoutAnswer() StepModel
}
