package chat

import (
	"errors"
	"time"
)

type (
	Message struct {
		ID        string
		Message   string
		User      User
		Timestamp time.Time
	}
	User struct {
		ID   string
		Name string
	}
	Connection struct {
		UserID       string
		Subscription chan Message
	}

	Option func(*service) error
)

func WithIDGenerator(generator IDGenerator) Option {
	return func(s *service) error {
		if generator == nil {
			return errors.New("id generator cannot be nil")
		}
		s.idGenerator = generator

		return nil
	}
}
