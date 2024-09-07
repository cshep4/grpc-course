package module7

// gRPC Layer
//go:generate mockgen -destination=internal/mocks/chat/mock_chat_service.gen.go -package=chat_mock github.com/cshep4/grpc-course/module7/internal/transport/grpc ChatService
//go:generate mockgen -destination=internal/mocks/grpc/mock_grpc_stream.gen.go -package=grpc_mock github.com/cshep4/grpc-course/module7/proto ChatService_SubscribeServer

// Service Layer
//go:generate mockgen -destination=internal/mocks/chat/mock_id_generator.gen.go -package=chat_mock github.com/cshep4/grpc-course/module7/internal/chat IDGenerator
//go:generate mockgen -destination=internal/mocks/store/mock_store.gen.go -package=store_mock github.com/cshep4/grpc-course/module7/internal/chat Store
