package interceptor

import (
	"context"
	"errors"

	"github.com/cshep4/go-log"
	"github.com/cshep4/grpc-course/module5/internal/auth"
	"github.com/cshep4/grpc-course/module5/proto"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

const userIDKey = "user_id"

type (
	Validator interface {
		ValidateToken(ctx context.Context, token string) (string, error)
	}

	middleware struct {
		validator Validator
	}
)

func NewMiddleware(validator Validator) (*middleware, error) {
	if validator == nil {
		return nil, errors.New("validator cannot be nil")
	}

	return &middleware{validator: validator}, nil
}

func (m *middleware) UnaryAuthMiddleware(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	// check the RPC method and only run if its the Protected RPC
	if info.FullMethod != proto.InterceptorService_Protected_FullMethodName {
		return handler(ctx, req)
	}

	// get the token from the metadata
	token, err := getTokenFromMetadata(ctx)
	if err != nil {
		return nil, status.Error(codes.PermissionDenied, "token must be provided")
	}

	// call validate token
	userID, err := m.validator.ValidateToken(ctx, token)
	if err != nil {
		if errors.Is(err, auth.ErrInvalidToken) {
			return nil, status.Error(codes.PermissionDenied, "invalid token")
		}

		log.Error(ctx, "failed to validate token", zap.Error(err))
		return nil, status.Error(codes.Internal, "error validating token")
	}

	// add the user ID to the context
	ctx = context.WithValue(ctx, userIDKey, userID)

	// call our handler
	return handler(ctx, req)
}

func getTokenFromMetadata(ctx context.Context) (string, error) {
	meta, ok := metadata.FromIncomingContext(ctx)
	if !ok || len(meta["authorization"]) != 1 {
		return "", errors.New("token not found in metadata")
	}

	return meta["authorization"][0], nil
}
