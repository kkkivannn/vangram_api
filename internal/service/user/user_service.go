package userService

import (
	"context"
	"vangram_api/internal/handlers"
	"vangram_api/internal/service/response"
)

type Storage interface {
	CreateUser(ctx context.Context, user handlers.RequestCreateUser) (int, error)
	ReadUser(ctx context.Context, id int) (response.UserResponse, error)
	UpdateUser(ctx context.Context, user handlers.RequestUpdateUser) ([]handlers.RequestUpdateUser, error)
	DeleteUser(ctx context.Context, id int) (string, error)
	GetAllUsers(ctx context.Context) ([]response.UserResponse, error)
}

type Service struct {
	repository *userStorage.Storage
}

func New(repository *userStorage.Storage) *Service {
	return &Service{repository}
}

func (as *Service) CreateUser(ctx context.Context, user handlers.RequestCreateUser) (int, error) {
	return as.repository.CreateUser(ctx, user)
}
func (as *Service) UpdateUser(ctx context.Context, user handlers.RequestUpdateUser) ([]handlers.RequestUpdateUser, error) {
	return as.repository.UpdateUser(ctx, user)
}

func (as *Service) DeleteUser(ctx context.Context, userId int) (string, error) {
	return as.repository.DeleteUser(ctx, userId)
}

func (as *Service) GetUser(ctx context.Context, userId int) (response.UserResponse, error) {
	return as.repository.ReadUser(ctx, userId)
}

func (as *Service) GetAllUsers(ctx context.Context) ([]response.UserResponse, error) {
	return as.repository.GetAllUsers(ctx)
}
