package service

import (
	"vangram_api/pkg/repository"
	"vangram_api/utils"
)

type AuthServiceInterface interface {
	CreateUser(dto utils.UserDto) (int, error)
	RemoveUser(userId int) (string, error)
	GetUser(userId int) (utils.UserDto, error)
	UpdateUser(userId int, user utils.UserDto) (utils.UserDto, error)
}

type AuthService struct {
	repository repository.AuthRepositoryInterface
}

func NewAuthService(repository repository.AuthRepositoryInterface) *AuthService {
	return &AuthService{repository: repository}
}

func (as *AuthService) CreateUser(user utils.UserDto) (int, error) {
	return as.repository.Create(user)
}

func (as *AuthService) RemoveUser(userId int) (string, error) {
	return as.repository.Delete(userId)
}

func (as *AuthService) GetUser(userId int) (utils.UserDto, error) {
	return as.repository.Read(userId)
}

func (as *AuthService) UpdateUser(userId int, user utils.UserDto) (utils.UserDto, error) {
	return as.repository.Update(userId, user)
}
