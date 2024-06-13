package handlers

import (
	"context"
	"github.com/gin-gonic/gin"
	"time"
	"vangram_api/internal/http/middleware"
	"vangram_api/internal/service/chat"
	"vangram_api/internal/service/friends"
	"vangram_api/internal/service/message"
	"vangram_api/internal/service/post"
	"vangram_api/internal/service/user"
)

type UserService interface {
	CreateUser(ctx context.Context, user user.RequestUser) (int, error)
	GetUser(ctx context.Context, id int) (user.User, error)
	UpdateUser(ctx context.Context, user user.RequestUser, userId int) error
	DeleteUser(ctx context.Context, id int) (string, error)
	GetAllUsers(ctx context.Context, userID int) ([]user.User, error)
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
	GetUserPosts(ctx context.Context, userID int) ([]post.Post, error)
}

type MessageService interface {
	AddNewMessage(ctx context.Context, message message.CreateMessage, senderID int) (int, error)
	GetChatMessages(ctx context.Context, chatID int, userID int) ([]message.MessagesChat, error)
}

type ChatService interface {
	AddNewChat(ctx context.Context, chat chat.CreateChatModel) (int, error)
	GetAllChats(ctx context.Context) ([]chat.Chat, error)
}

type FriendService interface {
	GetAllFriends(ctx context.Context, userID int) ([]friends.Friend, error)
	AddNewFriend(ctx context.Context, userID, friendID int) error
}

type Handler struct {
	userService    UserService
	postService    PostService
	messageService MessageService
	chatService    ChatService
	friendService  FriendService
}

func NewHandler(messageService MessageService, userService UserService, postService PostService, chatService ChatService, friendService FriendService) *Handler {
	return &Handler{messageService: messageService, userService: userService, postService: postService, chatService: chatService, friendService: friendService}
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
		api.GET("/ws", h.connectToSocket)
		u := api.Group("/user")
		{
			u.GET("", h.getUser)
			u.DELETE("", h.deleteUser)
			u.PATCH("", h.updateUser)
			u.GET("/profile", h.getProfile)
			u.GET("/friends", h.getAllFriends)

		}
		users := api.Group("/users")
		{
			users.GET("", h.getAllUsers)
			users.POST("", h.createFriend)
		}
		p := api.Group("/post")
		{
			p.GET("", h.getPost)
			p.POST("", h.createPost)
			p.POST("/like", h.setLike)
			p.GET("/likes", h.getLikesUserPosts)
		}
		api.GET("/posts", h.getAllPosts)
		api.GET("/users_posts", h.getUsersPosts)

		c := api.Group("/chat")
		{
			c.POST("", h.createChat)
			c.GET("/chats", h.getAllChats)
			c.GET("/messages", h.getMessages)
		}

	}
	return router

}
