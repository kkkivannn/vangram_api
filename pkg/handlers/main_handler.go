package handlers

import (
	"github.com/gin-gonic/gin"
	"vangram_api/pkg/service"
)

type MainHandlers struct {
	services service.AuthServiceInterface
}

func NewMainHandlers(services service.AuthServiceInterface) *MainHandlers {
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
