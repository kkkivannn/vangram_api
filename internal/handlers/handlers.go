package handlers

import (
	"context"
	"github.com/gin-gonic/gin"
	"vangram_api/internal/service/response"
)

type AuthService interface {
	CreateUser(ctx context.Context, user RequestCreateUser) (int, error)
	GetUser(ctx context.Context, id int) (response.UserResponse, error)
	UpdateUser(ctx context.Context, user RequestUpdateUser) ([]RequestUpdateUser, error)
	DeleteUser(ctx context.Context, id int) (string, error)
	GetAllUsers(ctx context.Context) ([]response.UserResponse, error)
}

type Handler struct {
	service AuthService
}

func New(services AuthService) *Handler {
	return &Handler{service: services}
}

func (h *Handler) Init() *gin.Engine {
	router := gin.New()

	api := router.Group("/api")
	{
		auth := api.Group("/auth")
		{
			auth.POST("/signUp", h.signUp)
		}
		user := api.Group("/user")
		{
			user.GET("", h.getUser)
			user.DELETE("", h.deleteUser)
		}
		api.GET("/users", h.getAllUsers)
	}
	return router
}
