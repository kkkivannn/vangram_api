package database

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sirupsen/logrus"
)

const (
	Client = "Client"
)

type ConfigDB struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func NewPostgresDB(context context.Context, config *ConfigDB) (pool *pgxpool.Pool, err error) {
	dbUrl := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", config.Username, config.Password, config.Host, config.Port, config.DBName)
	pool, err = pgxpool.Connect(context, dbUrl)
	if err != nil {
		logrus.Fatal(err.Error())
	}
	return pool, nil
}
