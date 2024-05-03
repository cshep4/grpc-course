package token

import (
	"context"

	"github.com/cshep4/grpc-course/module5-exercise/proto"
)

type Service struct {
	proto.UnimplementedTokenServiceServer
}

func (s Service) Validate(ctx context.Context, _ *proto.ValidateRequest) (*proto.ValidateResponse, error) {
	// your implementation goes here ...
	panic("implement me")
}
