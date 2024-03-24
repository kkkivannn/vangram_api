package handlers

import (
	"context"
	"time"
	"vangram_api/internal/http/middleware"
	"vangram_api/internal/service/post"
	"vangram_api/internal/service/user"

	"github.com/gin-gonic/gin"
)

type UserService interface {
	CreateUser(ctx context.Context, user user.RequestUser) (int, error)
	GetUser(ctx context.Context, id int) (user.User, error)
	UpdateUser(ctx context.Context, user user.RequestUser, userId int) error
	DeleteUser(ctx context.Context, id int) (string, error)
	GetAllUsers(ctx context.Context) ([]user.User, error)
	GetUserByNumber(ctx context.Context, number string) (user.User, error)
	GenerateTokens(ctx context.Context, number string) (user.Tokens, error)
	RefreshTokens(ctx context.Context, refreshToken string) (user.Tokens, error)
	RemoveUserSession(ctx context.Context, sessionId string, userID int) error
}

type PostService interface {
	CreateUserPost(ctx context.Context, post post.CreatePostModel) (int, error)
	GetPost(ctx context.Context, postID int) (post.Post, error)
	GetAllPosts(ctx context.Context) ([]post.Post, error)
	SetLikeToPost(ctx context.Context, postID int) error
	AddLikesPost(ctx context.Context, postID, userID, userPostID int, likedAt time.Time) error
	GetLikesUsersPosts(ctx context.Context, userID int) ([]post.Post, error)
}

type Handler struct {
	userService UserService
	postService PostService
}

func New(userService UserService, postService PostService) *Handler {
	return &Handler{userService: userService, postService: postService}

}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	router.Static("/image", "./")
	api := router.Group("/api")
	{
		auth := api.Group("/auth")
		{
			auth.POST("/send_number", h.sendNumber)
			auth.POST("/send_code", h.sendCode)
			auth.POST("/refresh", h.refreshToken)
			auth.Use(middleware.Identity)
			auth.POST("/signOut", h.signOut)
		}
		api.Use(middleware.Identity)
		u := api.Group("/user")
		{
			u.GET("", h.getUser)
			u.DELETE("", h.deleteUser)
			u.PATCH("", h.updateUser)

		}
		api.GET("/users", h.getAllUsers)
		post := api.Group("/post")
		{
			post.POST("", h.createPost)
			post.GET("", h.getPost)
			post.POST("/like", h.setLike)
			post.GET("/likes", h.getLikesUserPosts)
		}
		api.GET("/posts", h.getAllPosts)

	}
	return router

}
