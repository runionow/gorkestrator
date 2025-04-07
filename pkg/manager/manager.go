package manager

import (
	"fmt"
	"runionow/gork/pkg/task"

	"github.com/golang-collections/collections/queue"
	"github.com/google/uuid"
)

/*
1. Accepts requests from the customer
2. Schedule tasks onto the worker
3. Keep track of tasks and their state
*/
type Manager struct {
	Pending       queue.Queue
	TaskDb        map[string][]*task.Task
	EventDb       map[string][]*task.TaskEvent
	Workers       []string
	WorkerTaskMap map[string]uuid.UUID
	TaskWorkerMap map[uuid.UUID]string
}

func (m *Manager) SelectWorker() {
	fmt.Println("Select the worker")
}

func (m *Manager) UpdateTasks() {
	fmt.Println("Update the tasks")
}

func (m *Manager) SendWork() {
	fmt.Println("I will send to the worker")
}
