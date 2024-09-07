package todo

import (
	"context"
	"errors"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	todo "github.com/cshep4/grpc-course/07-todo-service/internal/store"
	"github.com/cshep4/grpc-course/07-todo-service/proto"
)

type (
	TaskStore interface {
		AddTask(task string) (string, error)
		CompleteTask(taskID string) error
		ListTasks() ([]todo.Task, error)
	}
	service struct {
		proto.UnimplementedTodoServiceServer
		store TaskStore
	}
)

func NewService(store TaskStore) (*service, error) {
	if store == nil {
		return nil, errors.New("store cannot be nil")
	}

	return &service{
		store: store,
	}, nil
}

func (s *service) AddTask(ctx context.Context, request *proto.AddTaskRequest) (*proto.AddTaskResponse, error) {
	if request.GetTask() == "" {
		return nil, status.Error(codes.InvalidArgument, "task cannot be empty")
	}

	id, err := s.store.AddTask(request.Task)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to add task: %v", err)
	}

	return &proto.AddTaskResponse{
		Id: id,
	}, nil
}

func (s *service) CompleteTask(ctx context.Context, request *proto.CompleteTaskRequest) (*proto.CompleteTaskResponse, error) {
	if err := s.store.CompleteTask(request.GetId()); err != nil {
		if errors.Is(err, todo.ErrTaskNotFound) {
			return nil, status.Error(codes.NotFound, "task not found")
		}

		return nil, status.Errorf(codes.Internal, "failed to complete task: %v", err)
	}

	return &proto.CompleteTaskResponse{}, nil
}

func (s *service) ListTasks(ctx context.Context, request *proto.ListTasksRequest) (*proto.ListTasksResponse, error) {
	storedTasks, err := s.store.ListTasks()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to list tasks: %v", err)
	}

	tasks := make([]*proto.Task, 0, len(storedTasks))

	for _, task := range storedTasks {
		tasks = append(tasks, &proto.Task{
			Id:   task.ID,
			Task: task.Task,
		})
	}

	return &proto.ListTasksResponse{
		Tasks: tasks,
	}, nil
}
