package service

import (
	"context"
	"vangram_api/internal/repository"
	"vangram_api/utils"
)

type AuthRepository interface {
	Create(ctx context.Context, user *utils.Request) (int, error)
	Read(ctx context.Context, id int) (utils.Request, error)
	Update(ctx context.Context, user *utils.Request) ([]utils.Request, error)
	Delete(ctx context.Context, id int) (string, error)
}

type AuthService struct {
	repository *repository.AuthorizeRepository
}

func NewAuthService(repository *repository.AuthorizeRepository) *AuthService {
	return &AuthService{repository}
}

func (as *AuthService) CreateUser(ctx context.Context, user *utils.Request) (int, error) {
	return as.repository.Create(ctx, user)
}
func (as *AuthService) UpdateUser(ctx context.Context, user *utils.Request) ([]utils.Request, error) {
	return as.repository.Update(ctx, user)
}

func (as *AuthService) DeleteUser(ctx context.Context, userId int) (string, error) {
	return as.repository.Delete(ctx, userId)
}

func (as *AuthService) GetUser(ctx context.Context, userId int) (utils.Request, error) {
	return as.repository.Read(ctx, userId)
}
