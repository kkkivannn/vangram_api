package service

import (
	"context"
	"vangram_api/pkg/repository"
	"vangram_api/utils"
)

type AuthServiceInterface interface {
	CreateUser(ctx context.Context, user *utils.UserDTO) (int, error)
	GetUser(ctx context.Context, id int) (utils.UserDTO, error)
	UpdateUser(ctx context.Context, user *utils.UserDTO) ([]utils.UserDTO, error)
	DeleteUser(ctx context.Context, id int) (string, error)
}

type AuthService struct {
	repository *repository.AuthRepository
}

func NewAuthService(repository *repository.AuthRepository) *AuthService {
	return &AuthService{repository}
}

func (as *AuthService) CreateUser(ctx context.Context, user *utils.UserDTO) (int, error) {
	return as.repository.Create(ctx, user)
}
func (as *AuthService) UpdateUser(ctx context.Context, user *utils.UserDTO) ([]utils.UserDTO, error) {
	return as.repository.Update(ctx, user)
}

func (as *AuthService) DeleteUser(ctx context.Context, userId int) (string, error) {
	return as.repository.Delete(ctx, userId)
}

func (as *AuthService) GetUser(ctx context.Context, userId int) (utils.UserDTO, error) {
	return as.repository.Read(ctx, userId)
}
