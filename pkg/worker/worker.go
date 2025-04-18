package worker

import (
	"fmt"
	"runionow/gork/pkg/task"

	"github.com/golang-collections/collections/queue"
	"github.com/google/uuid"
)

// Responsible for managing the task
type Worker struct {
	Name      string
	Queue     queue.Queue
	Db        map[uuid.UUID]*task.Task
	TaskCount int
}

func (w *Worker) CollectStats() {
	fmt.Println("Collect the stats")
}

func (w *Worker) RunTask() {
	fmt.Println("Run the task")
}

func (w *Worker) StartTask() {
	fmt.Println("Start the task")
}

func (w *Worker) StopTask() {
	fmt.Println("Stop the task")
}
