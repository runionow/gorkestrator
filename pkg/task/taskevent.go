package task

import (
	"time"

	"github.com/google/uuid"
)

// TaskEvent is internal to the worker and maintains the lifecycle of the Task
type TaskEvent struct {
	ID        uuid.UUID
	State     State
	Timestamp time.Time
	Task      Task
}
