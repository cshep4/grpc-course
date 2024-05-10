package interceptor

import (
	"context"
	"fmt"
	"log"
	"time"

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

func (s Service) LongRunning(ctx context.Context, req *proto.LongRunningRequest) (*proto.LongRunningResponse, error) {
	select {
	case <-time.Tick(time.Second * 5):
		log.Println("finished waiting, not end request successfully")
	case <-ctx.Done():
		log.Println("context cancelled")
		return nil, ctx.Err()
	}

	return &proto.LongRunningResponse{}, nil
}
