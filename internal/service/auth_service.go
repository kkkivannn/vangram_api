package service

import (
	"context"
	"vangram_api/internal/handlers"
	"vangram_api/internal/lib/api/response"
	"vangram_api/internal/repository"
)

type AuthRepository interface {
	Create(ctx context.Context, user handlers.RequestCreateUser) (int, error)
	Read(ctx context.Context, id int) (response.UserResponse, error)
	Update(ctx context.Context, user handlers.RequestUpdateUser) ([]handlers.RequestUpdateUser, error)
	Delete(ctx context.Context, id int) (string, error)
	GetAll(ctx context.Context) ([]response.UserResponse, error)
}

type AuthService struct {
	repository *repository.AuthorizeRepository
}

func NewAuthService(repository *repository.AuthorizeRepository) *AuthService {
	return &AuthService{repository}
}

func (as *AuthService) CreateUser(ctx context.Context, user handlers.RequestCreateUser) (int, error) {
	return as.repository.Create(ctx, user)
}
func (as *AuthService) UpdateUser(ctx context.Context, user handlers.RequestUpdateUser) ([]handlers.RequestUpdateUser, error) {
	return as.repository.Update(ctx, user)
}

func (as *AuthService) DeleteUser(ctx context.Context, userId int) (string, error) {
	return as.repository.Delete(ctx, userId)
}

func (as *AuthService) GetUser(ctx context.Context, userId int) (response.UserResponse, error) {
	return as.repository.Read(ctx, userId)
}

func (as *AuthService) GetAllUsers(ctx context.Context) ([]response.UserResponse, error) {
	return as.repository.GetAll(ctx)
}
