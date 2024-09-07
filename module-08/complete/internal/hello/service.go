package hello

import (
	"context"
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/cshep4/grpc-course/module8/proto"
)

type Service struct {
	proto.UnimplementedHelloServiceServer
}

func (s *Service) SayHello(ctx context.Context, request *proto.SayHelloRequest) (*proto.SayHelloResponse, error) {
	if request.GetName() == "" {
		return nil, status.Error(codes.InvalidArgument, "name cannot be empty")
	}

	return &proto.SayHelloResponse{
		Message: fmt.Sprintf("Hello %s", request.GetName()),
	}, nil
}
