package database

import (
	"github.com/brendontj/review-chatbot/core/gateway"
	"github.com/google/uuid"
)

type WorkflowWithStepsModel struct {
	ID     uuid.UUID
	TypeF  string
	StepsF []gateway.StepModel
}

func (w WorkflowWithStepsModel) Id() uuid.UUID {
	return w.ID
}

func (w WorkflowWithStepsModel) Type() string {
	return w.TypeF
}

func (w WorkflowWithStepsModel) Steps() []gateway.StepModel {
	return w.StepsF
}

func (w WorkflowWithStepsModel) StepWithoutAnswer() gateway.StepModel {
	for _, step := range w.StepsF {
		if step.Answer() == "" {
			return step
		}
	}
	return StepModel{}
}
