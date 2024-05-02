package message

import "time"

type Message struct {
	ID int `json:"id,omitempty" db:"id"`
	//IDChat     int        `json:"id_chat,omitempty" db:"id_chat"`
	//IDUser     int        `json:"id_user,omitempty" db:"id_user"`
	Body       string     `json:"body,omitempty" db:"body"`
	CreatedAt  time.Time  `json:"created_at,omitempty" db:"created_at"`
	UploadedAt *time.Time `json:"uploaded_at,omitempty" db:"uploaded_at"`
}

type CreateMessage struct {
	IDChat     int        `json:"id_chat,omitempty" db:"id_chat"`
	IDUser     int        `json:"id_user,omitempty" db:"id_user"`
	Body       string     `json:"body,omitempty" db:"body"`
	CreatedAt  time.Time  `json:"created_at,omitempty" db:"created_at"`
	UploadedAt *time.Time `json:"uploaded_at,omitempty" db:"uploaded_at"`
}

type SocketMessage struct {
	IDChat int    `json:"id_chat,omitempty" db:"id_chat"`
	IDUser int    `json:"id_user,omitempty" db:"id_user"`
	Body   string `json:"body,omitempty" db:"body"`
}

type MessagesChat struct {
	Messages Message     `json:"messages"`
	User     UserMessage `json:"user"`
}

type UserMessage struct {
	ID      int     `json:"id" db:"user_id"`
	Name    string  `json:"name" db:"name"`
	Surname string  `json:"surname" db:"surname"`
	Photo   *string `json:"photo" db:"photo, omitempty"`
}
