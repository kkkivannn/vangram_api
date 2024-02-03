package main

import (
	"context"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"os"
	api "vangram_api"
	"vangram_api/pkg/database"
	"vangram_api/pkg/handlers"
	"vangram_api/pkg/repository"
	"vangram_api/pkg/service"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	err := godotenv.Load(".env")
	if err != nil {
		logrus.Fatal(err.Error())
	}

	db, err := database.NewPostgresDB(context.Background(), &database.ConfigDB{
		Host:     os.Getenv("HOST"),
		Port:     os.Getenv("PORT"),
		Username: os.Getenv("USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DBNAME"),
		SSLMode:  os.Getenv("SSLMODE"),
	})
	if err != nil {
		logrus.Fatal("Неполучилось инициализировать бд: ", err.Error())
	}

	repositories := repository.NewAuth(db)
	services := service.NewAuthService(repositories)
	mainHandlers := handlers.NewMainHandlers(services)

	var server = new(api.Server)

	if err := server.Run(os.Getenv("DB_PORT"), mainHandlers.InitHandlers()); err != nil {
		logrus.Fatalf("Server not running: %s", err.Error())
	}

}
