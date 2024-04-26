package main

import (
	"log/slog"
	"os"

	config "github.com/KaffeeMaschina/http-rest-api/internal"
	mwLogger "github.com/KaffeeMaschina/http-rest-api/internal/http-server/middlewarelogger"
	"github.com/KaffeeMaschina/http-rest-api/internal/lib/logger/sl"
	"github.com/KaffeeMaschina/http-rest-api/internal/nats"
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

	log := setupLogger(cfg.Env)

	log.Info("i", slog.String("env", cfg.Env))
	log.Debug("b")

	csh := postgres.NewCashe()
	storage, err := postgres.New(cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.Database, csh)
	if err != nil {
		log.Error("failed to init storage", sl.Err(err))
		os.Exit(1)
	}

	nats.NewStreamingHandler(storage)

	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(mwLogger.New())
	router.Use(middleware.Recoverer)

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
