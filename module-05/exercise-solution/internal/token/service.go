package token

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/cshep4/grpc-course/module5-exercise/proto"
)

type Service struct {
	proto.UnimplementedTokenServiceServer
}

func (s Service) Validate(ctx context.Context, _ *proto.ValidateRequest) (*proto.ValidateResponse, error) {
	claims, ok := ctx.Value(claimsKey).(map[string]string)
	if !ok {
		return nil, status.Error(codes.FailedPrecondition, "claims missing from context")
	}

	return &proto.ValidateResponse{
		Claims: claims,
	}, nil
}
