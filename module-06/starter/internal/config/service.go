package config

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/cshep4/grpc-course/module6/proto"
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

func (s service) LongRunning(ctx context.Context, req *proto.LongRunningRequest) (*proto.LongRunningResponse, error) {
	select {
	case <-time.Tick(time.Second * 5):
		log.Println("finish request")
	case <-ctx.Done():
		log.Println("context done")
	}

	return &proto.LongRunningResponse{}, nil
}
