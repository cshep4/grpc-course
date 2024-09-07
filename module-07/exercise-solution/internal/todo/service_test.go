package todo_test

import (
	"context"
	"errors"
	"fmt"
	"testing"

	store_mock "github.com/cshep4/grpc-course/07-todo-service/internal/mocks/store"
	todostore "github.com/cshep4/grpc-course/07-todo-service/internal/store"
	"github.com/cshep4/grpc-course/07-todo-service/internal/todo"
	"github.com/cshep4/grpc-course/07-todo-service/proto"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestNewService(t *testing.T) {
	t.Run("returns error if todo store is nil", func(t *testing.T) {
		service, err := todo.NewService(nil)
		require.Error(t, err)

		assert.Empty(t, service)
	})

	t.Run("successfully initialises todo grpc service", func(t *testing.T) {
		service, err := todo.NewService(store_mock.NewMockTaskStore(gomock.NewController(t)))
		require.NoError(t, err)

		assert.NotEmpty(t, service)
	})
}

func TestService_AddTask(t *testing.T) {
	t.Run("returns INVALID_ARGUMENT status code when task is empty", func(t *testing.T) {
		var (
			ctrl, ctx = gomock.WithContext(context.Background(), t)
			todoStore = store_mock.NewMockTaskStore(ctrl)
		)
		service, err := todo.NewService(todoStore)
		require.NoError(t, err)

		res, err := service.AddTask(ctx, &proto.AddTaskRequest{Task: ""})
		require.Error(t, err)
		require.Empty(t, res)

		statusErr, ok := status.FromError(err)
		require.True(t, ok)

		assert.Equal(t, codes.InvalidArgument, statusErr.Code())
		assert.Equal(t, "task cannot be empty", statusErr.Message())
	})

	t.Run("returns INTERNAL status code when an error is returned from the store", func(t *testing.T) {
		const task = "wake up"
		var (
			ctrl, ctx = gomock.WithContext(context.Background(), t)
			todoStore = store_mock.NewMockTaskStore(ctrl)

			testErr = errors.New("some error")
		)
		service, err := todo.NewService(todoStore)
		require.NoError(t, err)

		todoStore.EXPECT().AddTask(task).Return("", testErr)

		res, err := service.AddTask(ctx, &proto.AddTaskRequest{Task: task})
		require.Error(t, err)
		require.Empty(t, res)

		statusErr, ok := status.FromError(err)
		require.True(t, ok)

		assert.Equal(t, codes.Internal, statusErr.Code())
		assert.Equal(t, fmt.Sprintf("failed to add task: %v", testErr), statusErr.Message())
	})

	t.Run("returns task ID when task is stored successfully", func(t *testing.T) {
		const (
			task   = "wake up"
			taskID = "some task id"
		)
		var (
			ctrl, ctx = gomock.WithContext(context.Background(), t)
			todoStore = store_mock.NewMockTaskStore(ctrl)
		)
		service, err := todo.NewService(todoStore)
		require.NoError(t, err)

		todoStore.EXPECT().AddTask(task).Return(taskID, nil)

		res, err := service.AddTask(ctx, &proto.AddTaskRequest{Task: task})
		require.NoError(t, err)
		require.NotEmpty(t, res)

		assert.Equal(t, taskID, res.GetId())
	})
}

func TestService_CompleteTask(t *testing.T) {
	t.Run("returns NOT_FOUND status code when task is not found", func(t *testing.T) {
		const taskID = "some task id"
		var (
			ctrl, ctx = gomock.WithContext(context.Background(), t)
			todoStore = store_mock.NewMockTaskStore(ctrl)
		)
		service, err := todo.NewService(todoStore)
		require.NoError(t, err)

		todoStore.EXPECT().CompleteTask(taskID).Return(todostore.ErrTaskNotFound)

		res, err := service.CompleteTask(ctx, &proto.CompleteTaskRequest{Id: taskID})
		require.Error(t, err)
		require.Empty(t, res)

		statusErr, ok := status.FromError(err)
		require.True(t, ok)

		assert.Equal(t, codes.NotFound, statusErr.Code())
		assert.Equal(t, "task not found", statusErr.Message())
	})

	t.Run("returns INTERNAL status code when an error is returned from the store", func(t *testing.T) {
		const taskID = "some task id"
		var (
			ctrl, ctx = gomock.WithContext(context.Background(), t)
			todoStore = store_mock.NewMockTaskStore(ctrl)

			testErr = errors.New("some error")
		)
		service, err := todo.NewService(todoStore)
		require.NoError(t, err)

		todoStore.EXPECT().CompleteTask(taskID).Return(testErr)

		res, err := service.CompleteTask(ctx, &proto.CompleteTaskRequest{Id: taskID})
		require.Error(t, err)
		require.Empty(t, res)

		statusErr, ok := status.FromError(err)
		require.True(t, ok)

		assert.Equal(t, codes.Internal, statusErr.Code())
		assert.Equal(t, fmt.Sprintf("failed to complete task: %v", testErr), statusErr.Message())
	})

	t.Run("returns successful response when a task is completed", func(t *testing.T) {
		const taskID = "some task id"
		var (
			ctrl, ctx = gomock.WithContext(context.Background(), t)
			todoStore = store_mock.NewMockTaskStore(ctrl)
		)
		service, err := todo.NewService(todoStore)
		require.NoError(t, err)

		todoStore.EXPECT().CompleteTask(taskID).Return(nil)

		res, err := service.CompleteTask(ctx, &proto.CompleteTaskRequest{Id: taskID})
		require.NoError(t, err)

		assert.NotNil(t, res)
	})
}

func TestService_ListTasks(t *testing.T) {
	t.Run("returns INTERNAL status code when an error is returned from store", func(t *testing.T) {
		var (
			ctrl, ctx = gomock.WithContext(context.Background(), t)
			todoStore = store_mock.NewMockTaskStore(ctrl)

			testErr = errors.New("some error")
		)
		service, err := todo.NewService(todoStore)
		require.NoError(t, err)

		todoStore.EXPECT().ListTasks().Return(nil, testErr)

		res, err := service.ListTasks(ctx, &proto.ListTasksRequest{})
		require.Error(t, err)
		require.Empty(t, res)

		statusErr, ok := status.FromError(err)
		require.True(t, ok)

		assert.Equal(t, codes.Internal, statusErr.Code())
		assert.Equal(t, fmt.Sprintf("failed to list tasks: %v", testErr), statusErr.Message())
	})

	t.Run("returns a list of tasks retrieved from store", func(t *testing.T) {
		const (
			taskID1 = "some task id 1"
			taskID2 = "some task id 2"
			taskID3 = "some task id 3"

			task1 = "wake up"
			task2 = "walk the dog"
			task3 = "have breakfast"
		)
		var (
			ctrl, ctx = gomock.WithContext(context.Background(), t)
			todoStore = store_mock.NewMockTaskStore(ctrl)

			tasks = []todostore.Task{
				{ID: taskID1, Task: task1},
				{ID: taskID2, Task: task2},
				{ID: taskID3, Task: task3},
			}
		)
		service, err := todo.NewService(todoStore)
		require.NoError(t, err)

		todoStore.EXPECT().ListTasks().Return(tasks, nil)

		res, err := service.ListTasks(ctx, &proto.ListTasksRequest{})
		require.NoError(t, err)
		require.Len(t, res.GetTasks(), 3)

		assert.Equal(t, taskID1, res.Tasks[0].Id)
		assert.Equal(t, task1, res.Tasks[0].Task)
		assert.Equal(t, taskID2, res.Tasks[1].Id)
		assert.Equal(t, task2, res.Tasks[1].Task)
		assert.Equal(t, taskID3, res.Tasks[2].Id)
		assert.Equal(t, task3, res.Tasks[2].Task)
	})
}
