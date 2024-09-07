package auth

import (
	"context"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
)

type service struct {
	secret []byte
}

var ErrInvalidToken = errors.New("invalid token")

const claimID = "ID"

func NewService(secret string) (*service, error) {
	if secret == "" {
		return nil, errors.New("cannot have an empty secret")
	}
	return &service{secret: []byte(secret)}, nil
}

func (s *service) IssueToken(_ context.Context, userID string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		claimID: userID,
	}, nil)

	signed, err := token.SignedString(s.secret)
	if err != nil {
		return "", fmt.Errorf("failed to sign JWT: %w", err)
	}
	return signed, nil
}

func (s *service) ValidateToken(_ context.Context, token string) (string, error) {
	t, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return s.secret, nil
	})
	if err != nil {
		return "", errors.Join(ErrInvalidToken, err)
	}

	if claims, ok := t.Claims.(jwt.MapClaims); ok && t.Valid {
		id, ok := claims[claimID].(string)
		if !ok {
			return "", fmt.Errorf("%w: failed to extract id from Claims", ErrInvalidToken)
		}

		return id, nil
	}

	return "", ErrInvalidToken
}
