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
	"vangram_api/internal/service/post"
	"vangram_api/internal/service/user"
	"vangram_api/internal/storage"
	"vangram_api/internal/storage/postgres"
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
		slog.Error("Failed to init userStorage", err.Error())
		os.Exit(1)
	}

	userStorage := storage.NewUserStorage(db)
	userService := user.New(userStorage)

	postStorage := storage.NewPostStorage(db)
	postService := post.NewService(postStorage)

	route := handlers.New(userService, postService)

	done := make(chan os.Signal, 1)

	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	srv := http.New(cfg, route.InitRoutes())

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Error("failed to start server")
		}
		slog.Info("server started")
	}()

	<-done
	slog.Info("stopping server")

	if err := srv.Shutdown(ctx); err != nil {
		slog.Error("failed to stop server", err.Error())
		return
	}

	slog.Info("server stopped")
}
