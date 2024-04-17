package main

import (
	"fmt"
	"log/slog"
	"os"

	config "github.com/KaffeeMaschina/http-rest-api/internal"
	"github.com/KaffeeMaschina/http-rest-api/internal/lib/logger/sl"
	"github.com/KaffeeMaschina/http-rest-api/internal/nats-streaming"
	postgres "github.com/KaffeeMaschina/http-rest-api/internal/storage/postgres"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	cfg := config.MustLoad()
	fmt.Println(cfg)

	log := setupLogger(cfg.Env)

	log.Info("i", slog.String("env", cfg.Env))
	log.Debug("b")

	storage, err := postgres.New(cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.Database)
	if err != nil {
		log.Error("failed to init storage", sl.Err(err))
		os.Exit(1)

	}
	_ = storage
	nats.Subscriber()
	nats.Publisher()

	router := chi.NewRouter()
	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)

}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger
	switch env {
	case envLocal:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envDev:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	}
	return log
}
