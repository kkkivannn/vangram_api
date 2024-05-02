package chat

import (
	"context"
	"time"
)

type StorageChat interface {
	CreateChat(ctx context.Context, chat CreateChatModel) (int, error)
}
type Service struct {
	storage StorageChat
}

func NewChatService(storage StorageChat) *Service {
	return &Service{storage: storage}
}

func (s *Service) AddNewChat(ctx context.Context, chat CreateChatModel) (int, error) {
	chat.CreatedAt = time.Now()
	return s.storage.CreateChat(ctx, chat)
}
