package database

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
	"vangram_api/internal/config"
)

const (
	Client = "Client"
)

func NewPostgresDB(context context.Context, config *config.Config) (pool *pgxpool.Pool, err error) {
	dbUrl := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", config.DBConfig.Username, config.DBConfig.Password, config.DBConfig.DbHost, config.DBConfig.DbPort, config.DBConfig.DbName)
	pool, err = pgxpool.Connect(context, dbUrl)
	if err != nil {
		log.Fatal(err.Error())
	}
	return pool, nil
}
