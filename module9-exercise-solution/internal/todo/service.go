package todo

import (
	"connectrpc.com/connect"
	"context"
	"errors"

	"github.com/cshep4/grpc-course/09-todo-service/proto"
	"github.com/google/uuid"
)

type service struct {
	proto.UnimplementedTodoServiceServer
	tasks map[string]string
}

func NewService() *service {
	return &service{
		tasks: make(map[string]string),
	}
}

func (s *service) AddTask(ctx context.Context, c *connect.Request[proto.AddTaskRequest]) (*connect.Response[proto.AddTaskResponse], error) {
	// validate input
	if c.Msg.GetTask() == "" {
		return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("task cannot be empty"))
	}

	// generate ID for task
	id := uuid.New().String()

	// add task to store
	s.tasks[id] = c.Msg.GetTask()

	// return generated ID
	return connect.NewResponse(&proto.AddTaskResponse{
		Id: id,
	}), nil
}

func (s *service) CompleteTask(ctx context.Context, c *connect.Request[proto.CompleteTaskRequest]) (*connect.Response[proto.CompleteTaskResponse], error) {
	// check if task exists
	if _, ok := s.tasks[c.Msg.GetId()]; !ok {
		return nil, connect.NewError(connect.CodeNotFound, errors.New("task not found"))
	}

	// remove task from store
	delete(s.tasks, c.Msg.GetId())

	// return response
	return connect.NewResponse(&proto.CompleteTaskResponse{}), nil
}

func (s *service) ListTasks(ctx context.Context, c *connect.Request[proto.ListTasksRequest]) (*connect.Response[proto.ListTasksResponse], error) {
	// initialise a slice of tasks
	tasks := make([]*proto.Task, 0, len(s.tasks))

	// iterate through tasks in our store
	for id, task := range s.tasks {
		tasks = append(tasks, &proto.Task{
			Id:   id,
			Task: task,
		})
	}

	// return list of tasks
	return connect.NewResponse(&proto.ListTasksResponse{
		Tasks: tasks,
	}), nil
}
