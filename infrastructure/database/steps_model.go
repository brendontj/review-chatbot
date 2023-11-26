package database

import (
	"github.com/google/uuid"
)

type StepModel struct {
	ID      uuid.UUID
	OrderF  int
	AnswerF string
}

func (s StepModel) Id() uuid.UUID {
	return s.ID
}

func (s StepModel) Order() int {
	return s.OrderF
}

func (s StepModel) Answer() string {
	return s.AnswerF
}
