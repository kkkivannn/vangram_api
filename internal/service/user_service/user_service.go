package user_service

import (
	"context"
	"vangram_api/internal/service"
	"vangram_api/internal/storage/user"
)

type Storage interface {
	CreateUser(ctx context.Context, user service.RequestUser) (int, error)
	ReadUser(ctx context.Context, id int) (service.User, error)
	UpdateUser(ctx context.Context, user service.User) ([]service.User, error)
	DeleteUser(ctx context.Context, id int) (string, error)
	GetAllUsers(ctx context.Context) ([]service.User, error)
}

type Service struct {
	repository *user.Storage
}

func New(repository *user.Storage) *Service {
	return &Service{repository}
}

func (as *Service) CreateUser(ctx context.Context, user service.RequestUser) (int, error) {
	return as.repository.CreateUser(ctx, user)
}
func (as *Service) UpdateUser(ctx context.Context, user service.User) ([]service.User, error) {
	return as.repository.UpdateUser(ctx, user)
}

func (as *Service) DeleteUser(ctx context.Context, userId int) (string, error) {
	return as.repository.DeleteUser(ctx, userId)
}

func (as *Service) GetUser(ctx context.Context, userId int) (service.User, error) {
	return as.repository.ReadUser(ctx, userId)
}

func (as *Service) GetAllUsers(ctx context.Context) ([]service.User, error) {
	return as.repository.GetAllUsers(ctx)
}
