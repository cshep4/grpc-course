package todo_test

import (
	store_mock "github.com/cshep4/grpc-course/07-todo-service/internal/mocks/store"
	"github.com/cshep4/grpc-course/07-todo-service/internal/todo"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
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

// your tests should go here...
