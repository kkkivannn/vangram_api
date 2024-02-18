package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"
	"vangram_api/internal/config"
	"vangram_api/internal/http"
	"vangram_api/internal/http/handlers"
	"vangram_api/internal/postgres"
	"vangram_api/internal/service/user"
	"vangram_api/internal/storage"
	"vangram_api/pkg/logger"
)

func main() {
	cfg := config.MustLoad()
	log := logger.SetupLogger(logger.EnvLocal)
	log.Info("Starting vangram_api 💅🏻", slog.String("env", cfg.Env), slog.String("version", "123"))

	log.Debug("debug messages are enabled")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	db, err := postgres.NewPostgresDB(ctx, cfg)
	defer db.Close()

	if err != nil {
		log.Error("Failed to init userStorage", err.Error())
		os.Exit(1)
	}

	userStorage := storage.NewUserStorage(db)
	userService := user.New(userStorage)

	route := handlers.New(userService)

	done := make(chan os.Signal, 1)

	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	srv := http.New(cfg, route.InitRoutes())

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Error("failed to start server")
		}
		log.Info("server started")
	}()

	<-done
	log.Info("stopping server")

	if err := srv.Shutdown(ctx); err != nil {
		log.Error("failed to stop server", err.Error())
		return
	}

	log.Info("server stopped")
}
