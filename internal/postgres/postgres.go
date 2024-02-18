package postgres

import (
	"context"
	"fmt"
	"log"
	"vangram_api/internal/config"

	"github.com/jackc/pgx/v4/pgxpool"
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
