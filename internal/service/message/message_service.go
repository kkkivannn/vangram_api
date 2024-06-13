package message

import (
	"context"
	"time"
)

type Storage interface {
	CreateMessage(ctx context.Context, message CreateMessage, senderID int) (int, error)
	ReadMessages(ctx context.Context, chatID int, userID int) ([]MessagesChat, error)
}

type Service struct {
	storage Storage
}

func NewMessageService(storage Storage) *Service {
	return &Service{storage: storage}
}

func (s *Service) AddNewMessage(ctx context.Context, message CreateMessage, senderID int) (int, error) {
	message.CreatedAt = time.Now()
	return s.storage.CreateMessage(ctx, message, senderID)
}

func (s *Service) GetChatMessages(ctx context.Context, chatID int, userID int) ([]MessagesChat, error) {
	return s.storage.ReadMessages(ctx, chatID, userID)
}
