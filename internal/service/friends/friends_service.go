package friends

import (
	"context"
)

type Storage interface {
	ReadAllFriends(ctx context.Context, userID int) ([]Friend, error)
	CreateFriend(ctx context.Context, userID, friendID int) error
}

type Service struct {
	storage Storage
}

func NewService(storage Storage) *Service {
	return &Service{storage: storage}
}

func (s *Service) GetAllFriends(ctx context.Context, userID int) ([]Friend, error) {
	return s.storage.ReadAllFriends(ctx, userID)
}

func (s *Service) AddNewFriend(ctx context.Context, userID, friendID int) error {
	return s.storage.CreateFriend(ctx, userID, friendID)
}
