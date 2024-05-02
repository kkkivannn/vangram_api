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
