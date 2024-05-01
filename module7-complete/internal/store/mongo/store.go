package mongo

import (
	"context"
	"errors"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"

	"github.com/cshep4/grpc-course/module7/internal/chat"
)

const (
	db         = "chat"
	collection = "messages"
)

type store struct {
	client     *mongo.Client
	collection *mongo.Collection
}

func New(ctx context.Context, client *mongo.Client) (*store, error) {
	if client == nil {
		return nil, errors.New("client cannot be nil")
	}

	s := &store{
		client:     client,
		collection: client.Database(db).Collection(collection),
	}

	ctx, cancel := context.WithTimeout(ctx, 1*time.Minute)
	defer cancel()

	if err := client.Ping(ctx, nil); err != nil {
		return nil, fmt.Errorf("failed to ping mongo: %w", err)
	}

	return s, nil
}

func (s *store) StoreMessage(ctx context.Context, msg chat.Message) error {
	_, err := s.collection.InsertOne(ctx, message{
		ID:        msg.ID,
		Message:   msg.Message,
		UserID:    msg.User.ID,
		UserName:  msg.User.Name,
		Timestamp: msg.Timestamp,
	})
	if err != nil {
		return fmt.Errorf("failed to insert message: %w", err)
	}

	return nil
}

func (s *store) Close(ctx context.Context) error {
	return s.client.Disconnect(ctx)
}
