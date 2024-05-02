package message

import (
	"context"
	"time"
)

type Storage interface {
	CreateMessage(ctx context.Context, message CreateMessage) (int, error)
	ReadMessages(ctx context.Context, chatID int) ([]MessagesChat, error)
}

type Service struct {
	storage Storage
}

func NewMessageService(storage Storage) *Service {
	return &Service{storage: storage}
}

func (s *Service) AddNewMessage(ctx context.Context, message CreateMessage) (int, error) {
	message.CreatedAt = time.Now()
	return s.storage.CreateMessage(ctx, message)
}

func (s *Service) GetChatMessages(ctx context.Context, chatID int) ([]MessagesChat, error) {
	return s.storage.ReadMessages(ctx, chatID)
}
