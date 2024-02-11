package handlers

import (
	"context"
	"github.com/gin-gonic/gin"
	"vangram_api/utils"
)

type AuthService interface {
	CreateUser(ctx context.Context, user utils.Request) (int, error)
	GetUser(ctx context.Context, id int) (utils.Request, error)
	UpdateUser(ctx context.Context, user utils.Request) ([]utils.Request, error)
	DeleteUser(ctx context.Context, id int) (string, error)
}

type MainHandlers struct {
	services AuthService
}

func NewMainHandlers(services AuthService) *MainHandlers {
	return &MainHandlers{services: services}
}

func (h *MainHandlers) InitHandlers() *gin.Engine {
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
	}
	return router
}
