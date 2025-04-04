package worker

import (
	"runionow/gork/pkg/task"

	"github.com/golang-collections/collections/queue"

	"github.com/google/uuid"
)

type Worker struct {
	Name      string
	Queue     queue.Queue
	Db        map[uuid.UUID]*task.Task
	TaskCount int
}
