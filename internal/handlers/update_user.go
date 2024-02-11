package handlers

import (
	"context"
	"vangram_api/internal/lib/api/response"
)

type UpdateUser interface {
	UpdateUser(ctx context.Context, user RequestCreateUser) ([]response.UserResponse, error)
}

type RequestUpdateUser struct {
	ID      int     `json:"id"`
	Name    *string `json:"name"`
	Surname *string `json:"surname"`
}
