package token

import (
	"context"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

const claimsKey = "claims"

type middleware struct {
	secret []byte
}

func NewMiddleware(secret []byte) (*middleware, error) {
	if secret == nil {
		return nil, errors.New("secret cannot be nil")
	}

	return &middleware{secret: secret}, nil
}

func (m *middleware) UnaryAuthMiddleware(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	token, err := getTokenFromMetadata(ctx)
	if err != nil {
		return nil, status.Error(codes.PermissionDenied, "token must be provided")
	}

	t, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return m.secret, nil
	})
	if err != nil {
		return "", status.Error(codes.PermissionDenied, "invalid token")
	}

	claimsMap := make(map[string]string)

	if claims, ok := t.Claims.(jwt.MapClaims); ok && t.Valid {
		for key, val := range claims {
			claim, ok := val.(string)
			if ok {
				claimsMap[key] = claim
				continue
			}

			ts, ok := val.(float64)
			if ok {
				claimsMap[key] = fmt.Sprintf("%.0f", ts)
			}
		}
	}
	ctx = context.WithValue(ctx, claimsKey, claimsMap)

	return handler(ctx, req)
}

func getTokenFromMetadata(ctx context.Context) (string, error) {
	meta, ok := metadata.FromIncomingContext(ctx)
	if !ok || len(meta["authorization"]) != 1 {
		return "", errors.New("token not found in metadata")
	}

	return meta["authorization"][0], nil
}
