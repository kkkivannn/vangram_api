package post

import (
	"mime/multipart"
	"time"
	"vangram_api/internal/service/user"
)

type CreatePostModel struct {
	Photo     *multipart.FileHeader `form:"photo" db:"photo"`
	Body      string                `form:"body" db:"body"`
	CreatedAt time.Time             `form:"created_at" db:"created_at"`
	UserID    int                   `form:"user_id" db:"id_user"`
}

type SavePost struct {
	Photo     *string   `json:"photo" db:"photo, omitempty"`
	Body      string    `json:"body" db:"body,omitempty"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UserID    int       `json:"user_id" db:"id_user"`
}

type Post struct {
	ID         int        `json:"id" db:"id"`
	Photo      string     `json:"photo" db:"photo, omitempty"`
	CountLikes int        `json:"count_likes" db:"count_likes"`
	Body       string     `json:"body" db:"body,omitempty"`
	CreatedAt  time.Time  `json:"created_at" db:"created_at"`
	UploadedAt *time.Time `json:"uploaded_at" db:"uploaded_at"`
	User       user.User  `json:"user" db:"id_user"`
}
