package user

import (
	"context"
)

type Storage interface {
	CreateUser(ctx context.Context, user RequestUser) (int, error)
	ReadUser(ctx context.Context, id int) (User, error)
	UpdateUser(ctx context.Context, user User) ([]User, error)
	DeleteUser(ctx context.Context, id int) (string, error)
	GetAllUsers(ctx context.Context) ([]User, error)
}

type Service struct {
	storage Storage
}

func New(storage Storage) *Service {
	return &Service{storage}
}

func (as *Service) CreateUser(ctx context.Context, user RequestUser) (int, error) {
	return as.storage.CreateUser(ctx, user)
}
func (as *Service) UpdateUser(ctx context.Context, user User) ([]User, error) {
	return as.storage.UpdateUser(ctx, user)
}

func (as *Service) DeleteUser(ctx context.Context, userId int) (string, error) {
	return as.storage.DeleteUser(ctx, userId)
}

func (as *Service) GetUser(ctx context.Context, userId int) (User, error) {
	return as.storage.ReadUser(ctx, userId)
}

func (as *Service) GetAllUsers(ctx context.Context) ([]User, error) {
	return as.storage.GetAllUsers(ctx)
}
