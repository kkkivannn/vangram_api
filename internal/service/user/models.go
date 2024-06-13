package user

import (
	"mime/multipart"
	"time"
)

type User struct {
	ID        int     `json:"id,omitempty" db:"id"`
	Name      string  `json:"name" db:"name"`
	Surname   string  `json:"surname" db:"surname"`
	Age       int     `json:"age"`
	Phone     string  `json:"phone,omitempty"`
	Photo     *string `json:"photo" db:"photo"`
	InFriends bool    `json:"in_friends"`
}

type RequestUser struct {
	Name       string                `form:"name" db:"name"`
	Surname    string                `form:"surname" db:"surname"`
	Age        int                   `form:"age" db:"age"`
	Phone      string                `form:"phone" db:"phone_number"`
	Photo      *multipart.FileHeader `form:"photo" db:"photo" binding:"omitempty"`
	CreatedAt  time.Time             `form:"createdAt" db:"created_at" binding:"omitempty"`
	UploadedAt *time.Time            `form:"uploadedAt" db:"uploaded_at" binding:"omitempty"`
}

type SaveUser struct {
	Name       string     `json:"name" db:"name"`
	Surname    string     `json:"surname" db:"surname"`
	Age        int        `json:"age" db:"age"`
	Phone      string     `json:"phone" db:"phone_number"`
	Photo      *string    `json:"photo" db:"photo"`
	CreatedAt  time.Time  `json:"createdAt" db:"created_at"`
	UploadedAt *time.Time `json:"uploadedAt" db:"uploaded_at"`
}

type UpdatedUser struct {
	Name       string     `json:"name" db:"name"`
	Surname    string     `json:"surname" db:"surname"`
	Age        int        `json:"age" db:"age"`
	Phone      string     `json:"phone" db:"phone_number"`
	Photo      *string    `json:"photo" db:"photo"`
	CreatedAt  time.Time  `json:"createdAt" db:"created_at"`
	UploadedAt *time.Time `json:"uploadedAt" db:"uploaded_at"`
}

type SignInUserRequest struct {
	Phone string `json:"phone" db:"phone_number"`
	Code  string `json:"code"`
}

type SendNumberRequest struct {
	Phone string `json:"phone" db:"phone_number"`
}

type Tokens struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type Ids struct {
	UserId    int
	SessionId string
}
