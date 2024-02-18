package post

import "github.com/jackc/pgx/v4/pgxpool"

type Storage struct {
	db *pgxpool.Pool
}

func New(db *pgxpool.Pool) *Storage {
	return &Storage{db: db}
}
