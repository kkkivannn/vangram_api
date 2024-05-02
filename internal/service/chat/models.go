package chat

import "time"

type CreateChatModel struct {
	Name      string    `json:"name" db:"name"`
	IDUser    int       `json:"id_user" db:"id_user"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}
