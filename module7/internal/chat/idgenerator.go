package chat

import "github.com/google/uuid"

type idGenerator struct{}

func (idGenerator) Generate() string {
	return uuid.New().String()
}
