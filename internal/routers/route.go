package routers

import (
	"context"
	"github.com/gin-gonic/gin"
	"vangram_api/internal/service"
)

type UserService interface {
	CreateUser(ctx context.Context, user service.RequestUser) (int, error)
	GetUser(ctx context.Context, id int) (service.User, error)
	UpdateUser(ctx context.Context, user service.User) ([]service.User, error)
	DeleteUser(ctx context.Context, id int) (string, error)
	GetAllUsers(ctx context.Context) ([]service.User, error)
}

type Route struct {
	service UserService
}

func New(service UserService) *Route {
	return &Route{service: service}

}

func (r *Route) InitRoutes() *gin.Engine {
	router := gin.New()
	api := router.Group("/api")
	{
		auth := api.Group("/auth")
		{
			auth.POST("/signUp", r.signUp)
		}
		user := api.Group("/user")
		{
			user.GET("", r.getUser)
			user.DELETE("", r.deleteUser)
		}
		api.GET("/users", r.getAllUsers)
	}
	return router

}
