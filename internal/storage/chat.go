package storage

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"vangram_api/internal/service/chat"
	"vangram_api/internal/storage/postgres"
)

type Chat struct {
	db *pgxpool.Pool
}

func NewChatStorage(db *pgxpool.Pool) *Chat {
	return &Chat{db: db}
}

func (c *Chat) CreateChat(ctx context.Context, chat chat.CreateChatModel) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (id_user, name, created_at) VALUES ($1, $2, $3) RETURNING id", postgres.Chat)
	row := c.db.QueryRow(ctx, query, chat.IDUser, chat.Name, chat.CreatedAt)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (c *Chat) GetChats(ctx context.Context) ([]chat.Chat, error) {
	var chats []chat.Chat
	query := fmt.Sprintf("SELECT c.id AS chat_id,c.name AS chat_name,c.photo AS chat_photo,c.created_at,u.id AS user_id,u.name AS user_name,u.surname AS user_surname,u.photo AS user_photo,m.id AS message_id,m.body AS message_body FROM %s c  LEFT JOIN LATERAL (SELECT id, id_chat, id_user, body FROM %s WHERE id_chat = c.id ORDER BY created_at DESC LIMIT 1) m  inner join %s u on m.id_user = u.id ON true", postgres.Chat, postgres.Message, postgres.User)
	rows, err := c.db.Query(ctx, query)
	if err != nil {
		return chats, err
	}
	defer rows.Close()
	for rows.Next() {
		var c chat.Chat
		var m = new(chat.SendLastMessage)
		var u = new(chat.SendLastUser)
		err := rows.Scan(&c.ID, &c.Name, &c.Photo, &c.CreatedAt, &u.ID, &u.Name, &u.Surname, &u.Photo, &m.ID, &m.Body)
		if err != nil {
			return chats, err
		}
		if u.ID == nil {
			m = nil
			u = nil
		} else {
			if c.Photo != nil {
				photo := fmt.Sprintf("%s%s", url, *c.Photo)
				c.Photo = &photo
			}
			if u.Photo != nil {
				photo := fmt.Sprintf("%s%s", url, *u.Photo)
				u.Photo = &photo
			}
			c.User = u
			c.Message = m
		}
		chats = append(chats, c)

	}
	return chats, nil
}
