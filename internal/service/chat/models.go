package chat

import (
	"time"
)

type CreateChatModel struct {
	Name      string    `json:"name" db:"name"`
	IDUser    int       `json:"id_user" db:"id_user"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

type Chat struct {
	ID        int              `json:"id" db:"chat_id"`
	Name      string           `json:"name" db:"chat_name"`
	Photo     *string          `json:"photo" db:"chat_photo"`
	CreatedAt time.Time        `json:"created_at" db:"created_at"`
	User      *SendLastUser    `json:"user"`
	Message   *SendLastMessage `json:"message"`
}

type SendLastUser struct {
	ID      *int    `json:"id,omitempty" db:"user_id"`
	Name    *string `json:"name,omitempty" db:"user_name"`
	Surname *string `json:"surname,omitempty" db:"user_surname"`
	Photo   *string `json:"photo,omitempty" db:"user_photo"`
}

type SendLastMessage struct {
	ID   *int    `json:"id,omitempty" db:"message_id"`
	Body *string `json:"body,omitempty" db:"message_body"`
}
