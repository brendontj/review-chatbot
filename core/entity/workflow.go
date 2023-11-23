package entity

import (
	"github.com/google/uuid"
)

type Workflow struct {
	ID   uuid.UUID
	Type string
}
