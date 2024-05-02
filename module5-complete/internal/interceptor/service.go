package interceptor

import (
	"context"
	"errors"
	"fmt"
	"github.com/cshep4/grpc-course/module5/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"log"
	"time"
)

type service struct {
	proto.UnimplementedInterceptorServiceServer
	name string
}

func NewService(name string) (*service, error) {
	if name == "" {
		return nil, errors.New("name cannot be empty")
	}

	return &service{
		name: name,
	}, nil
}

func (s service) SayHello(ctx context.Context, request *proto.SayHelloRequest) (*proto.SayHelloResponse, error) {
	start := time.Now()

	meta, ok := metadata.FromIncomingContext(ctx)
	if !ok || len(meta["x-request-id"]) == 0 {
		return nil, status.Error(codes.InvalidArgument, "missing x-request-id")
	}
	requestID := meta["x-request-id"][0]

	log.Printf("Request received with ID: %s", requestID)

	// send headers (before response message data is sent)
	header := metadata.New(map[string]string{"request-start-timestamp": start.String()})
	if err := grpc.SendHeader(ctx, header); err != nil {
		return nil, status.Error(codes.Internal, "failed to send headers")
	}

	// send trailers (special kind of header after response message data is sent)
	//
	// Trailers are a special kind of header that is sent after the message data.
	//
	// They are used internally to communicate the outcome of an RPC.
	//
	// At the application level, custom trailers can be used to communicate things not directly
	// part of the data, such as server utilization and query cost. Trailers are sent only by the server.
	trailer := metadata.Pairs("request-end-timestamp", time.Now().String())
	if err := grpc.SetTrailer(ctx, trailer); err != nil {
		return nil, status.Error(codes.Internal, "failed to send trailers")
	}

	return &proto.SayHelloResponse{
		Message: fmt.Sprintf("Hello: %s", request.GetName()),
	}, nil
}

func (s service) LongRunning(ctx context.Context, request *proto.LongRunningRequest) (*proto.LongRunningResponse, error) {
	select {
	case <-time.Tick(time.Second * 5):
		log.Println("finish request")
	case <-ctx.Done():
		log.Println("context done")
	}

	return &proto.LongRunningResponse{}, nil
}

func (s service) Protected(ctx context.Context, request *proto.ProtectedRequest) (*proto.ProtectedResponse, error) {
	userID, ok := ctx.Value(userIDKey).(string)
	if !ok {
		return nil, status.Error(codes.FailedPrecondition, "user id missing")
	}

	return &proto.ProtectedResponse{
		UserId: userID,
	}, nil
}
