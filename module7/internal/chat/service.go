package chat

import (
	"context"
	"errors"
	"fmt"
	"sync"
)

type (
	Store interface {
		StoreMessage(ctx context.Context, message Message) error
	}
	IDGenerator interface {
		Generate() string
	}
	service struct {
		store       Store
		idGenerator IDGenerator

		// connections grouped by chat_id and user_id
		connections map[string]map[string]Connection
		lock        sync.Mutex
	}
)

func NewService(store Store, opts ...Option) (*service, error) {
	if store == nil {
		return nil, errors.New("store cannot be empty")
	}

	svc := &service{
		store:       store,
		idGenerator: idGenerator{},
		connections: make(map[string]map[string]Connection),
		lock:        sync.Mutex{},
	}

	for _, opt := range opts {
		if err := opt(svc); err != nil {
			return nil, err
		}
	}

	return svc, nil
}

func (s *service) SendMessage(ctx context.Context, chatID string, message Message) (string, error) {
	// generate message ID
	id := s.idGenerator.Generate()
	message.ID = id

	// store message in DB
	if err := s.store.StoreMessage(ctx, message); err != nil {
		return "", fmt.Errorf("failed to store message: %w", err)
	}

	s.lock.Lock()
	defer s.lock.Unlock()

	if _, ok := s.connections[chatID]; !ok {
		return id, nil
	}

	// loop through active connections and send messages to clients
	for _, conn := range s.connections[chatID] {
		// prevent sender receiving their own messages
		if conn.UserID == message.User.ID {
			continue
		}

		// push message to connection channel
		conn.Subscription <- message
	}

	return id, nil
}

func (s *service) Subscribe(ctx context.Context, chatID, userID string) (chan Message, error) {
	s.lock.Lock()
	defer s.lock.Unlock()

	if _, ok := s.connections[chatID]; !ok {
		s.connections[chatID] = make(map[string]Connection)
	}

	// initialise a channel for user to receive messages on
	// using a buffered channel to remove chances of deadlocks
	// when a client unsubscribes
	msgChan := make(chan Message, 10)

	s.connections[chatID][userID] = Connection{
		UserID:       userID,
		Subscription: msgChan,
	}

	return msgChan, nil
}

func (s *service) Unsubscribe(ctx context.Context, chatID, userID string) error {
	s.lock.Lock()
	defer s.lock.Unlock()

	if _, ok := s.connections[chatID]; !ok {
		return nil
	}

	// close message chan for user and chatID
	close(s.connections[chatID][userID].Subscription)

	// remove user/chat connection
	delete(s.connections[chatID], userID)

	return nil
}
