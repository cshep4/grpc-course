package module7

//go:generate mockgen -destination=internal/mocks/store/mock_store.gen.go -package=store_mock github.com/cshep4/grpc-course/07-todo-service/internal/todo TaskStore
