package task

import (
	"github.com/google/uuid"
)

type Task struct {
	id    uuid.UUID
	name  string
	state State
}
