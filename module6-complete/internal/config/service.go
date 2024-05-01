package config

import (
	"context"
	"errors"
	"github.com/cshep4/grpc-course/module6/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"math/rand"
	"time"
)

type service struct {
	proto.UnimplementedConfigServiceServer
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

func (s service) GetServerName(ctx context.Context, request *proto.GetServerNameRequest) (*proto.GetServerNameResponse, error) {
	log.Printf("request received on server: %s", s.name)

	return &proto.GetServerNameResponse{
		Name: s.name,
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

func (s service) Flaky(ctx context.Context, request *proto.FlakyRequest) (*proto.FlakyResponse, error) {
	// Generate a random number between 0 and 2
	if rand.Intn(3) != 0 { // approximately 2 in 3 chance to be true
		log.Println("error response returned")
		return nil, status.Error(codes.Internal, "flaky error occurred") // Return an error 2 in 3 times
	}

	log.Println("successful response returned")

	return &proto.FlakyResponse{}, nil
}
