package handlers

import (
	"context"
	"github.com/gin-gonic/gin"
	"time"
	"vangram_api/internal/http/middleware"
	"vangram_api/internal/service/chat"
	"vangram_api/internal/service/message"
	"vangram_api/internal/service/post"
	"vangram_api/internal/service/user"
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

type MessageService interface {
	AddNewMessage(ctx context.Context, message message.CreateMessage) (int, error)
	GetChatMessages(ctx context.Context, chatID int) ([]message.MessagesChat, error)
}

type ChatService interface {
	AddNewChat(ctx context.Context, chat chat.CreateChatModel) (int, error)
}

type Handler struct {
	userService    UserService
	postService    PostService
	messageService MessageService
	chatService    ChatService
}

func NewHandler(userService UserService, postService PostService, messageService MessageService, chatService ChatService) *Handler {
	return &Handler{userService: userService, postService: postService, messageService: messageService, chatService: chatService}
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
		api.GET("/ws", h.connectToSocket)
		//api.Use(middleware.Identity)
		u := api.Group("/user")
		{
			u.GET("", h.getUser)
			u.DELETE("", h.deleteUser)
			u.PATCH("", h.updateUser)
			u.GET("/profile", h.getProfile)

		}
		api.GET("/users", h.getAllUsers)
		p := api.Group("/post", h.createPost)
		{
			p.GET("", h.getPost)
			p.POST("/like", h.setLike)
			p.GET("/likes", h.getLikesUserPosts)
		}
		api.GET("/posts", h.getAllPosts)

		c := api.Group("/chat")
		{
			c.POST("", h.createChat)
			c.GET("/chats", h.getAllChats)
			c.GET("/messages", h.getMessages)
		}

	}
	return router

}
