package interceptor

import (
	"context"
	"fmt"
	"github.com/cshep4/grpc-course/module5/proto"
)

type Service struct {
	proto.UnimplementedInterceptorServiceServer
}

func (s Service) SayHello(ctx context.Context, request *proto.SayHelloRequest) (*proto.SayHelloResponse, error) {
	return &proto.SayHelloResponse{
		Message: fmt.Sprintf("Hello: %s", request.GetName()),
	}, nil
}
