package handlers

import (
	"context"
	"vangram_api/internal/service/user"

	"github.com/gin-gonic/gin"
)

type UserService interface {
	CreateUser(ctx context.Context, user user.RequestUser) (int, error)
	GetUser(ctx context.Context, id int) (user.User, error)
	UpdateUser(ctx context.Context, user user.User) ([]user.User, error)
	DeleteUser(ctx context.Context, id int) (string, error)
	GetAllUsers(ctx context.Context) ([]user.User, error)
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
