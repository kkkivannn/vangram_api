package storage

import "github.com/jackc/pgx/v4/pgxpool"

type PostStorage struct {
	db *pgxpool.Pool
}

func NewPostStorage(db *pgxpool.Pool) *PostStorage {
	return &PostStorage{db: db}
}
