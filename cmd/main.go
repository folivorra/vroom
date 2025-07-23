package main

import (
	"context"
	"github.com/folivorra/vroom/internal/repository/postgres"
	"log/slog"
	"os"

	"github.com/folivorra/vroom/application"
	"github.com/folivorra/vroom/internal/config"
)

const (
	envLocal  = "local"
	envServer = "server"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cfg := config.MustLoad()

	log := setupLogger(cfg.Env)

	app := application.NewApp(ctx, log)
	defer app.Shutdown()

	_, err := postgres.NewPostgres(ctx, app, cfg.PostgresDsn)
	if err != nil {
		log.Error(err.Error())
		return
	}
	log.Info("postgres connected")

	app.Run()
}

func setupLogger(env string) *slog.Logger {
	var logger *slog.Logger

	switch env {
	case envLocal:
		logger = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case envServer:
		logger = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	}

	return logger
}
