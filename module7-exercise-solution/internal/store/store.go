package todo

import (
	"errors"

	"github.com/google/uuid"
)

type (
	store struct {
		tasks map[string]string
	}
	Task struct {
		ID   string
		Task string
	}
)

var ErrTaskNotFound = errors.New("task not found")

func NewStore() *store {
	return &store{
		tasks: make(map[string]string),
	}
}

// AddTask adds a task to the in-memory store.
func (s *store) AddTask(task string) (string, error) {
	// generate ID for task
	id := uuid.New().String()

	// add task to store
	s.tasks[id] = task

	// return generated ID
	return id, nil
}

// CompleteTask removes a task from the in-memory store.
func (s *store) CompleteTask(taskID string) error {
	// check if task exists
	if _, ok := s.tasks[taskID]; !ok {
		return ErrTaskNotFound
	}

	// remove task from store
	delete(s.tasks, taskID)

	// return response
	return nil
}

// ListTasks lists all outstanding tasks in the in-memory store.
func (s *store) ListTasks() ([]Task, error) {
	// initialise a slice of tasks
	tasks := make([]Task, 0, len(s.tasks))

	// iterate through tasks in our store
	for id, task := range s.tasks {
		tasks = append(tasks, Task{
			ID:   id,
			Task: task,
		})
	}

	// return list of tasks
	return tasks, nil
}
