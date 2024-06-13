package storage

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"vangram_api/internal/service/message"
	"vangram_api/internal/storage/postgres"
)

type Message struct {
	db *pgxpool.Pool
}

func NewMessageStorage(db *pgxpool.Pool) *Message {
	return &Message{db: db}
}

func (m *Message) CreateMessage(ctx context.Context, message message.CreateMessage, senderID int) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (id_chat, id_user, body, created_at) VALUES ($1, $2, $3, $4) RETURNING id", postgres.Message)
	row := m.db.QueryRow(ctx, query, message.IDChat, senderID, message.Body, message.CreatedAt)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (m *Message) ReadMessages(ctx context.Context, chatID int, userID int) ([]message.MessagesChat, error) {
	var messages []message.MessagesChat
	query := fmt.Sprintf("SELECT u.id AS user_id, u.name, u.surname, u.photo, m.body, m.created_at, m.uploaded_at, m.id, m.id_user AS message_id FROM %s m INNER JOIN %s u ON m.id_user = u.id AND m.id_chat = $1", postgres.Message, postgres.User)
	rows, err := m.db.Query(ctx, query, chatID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var messagesChat message.MessagesChat
		var mess message.Message
		var user message.UserMessage
		if err := rows.Scan(&user.ID, &user.Name, &user.Surname, &user.Photo, &mess.Body, &mess.CreatedAt, &mess.UploadedAt, &mess.ID, &mess.IDUser); err != nil {
			return nil, err
		}
		if user.Photo != nil {
			photo := fmt.Sprintf("%s%s", url, *user.Photo)
			user.Photo = &photo
		}
		mess.IsMine = userID == mess.IDUser
		messagesChat.Messages = mess
		messagesChat.User = user
		messages = append(messages, messagesChat)
	}
	return messages, nil
}
