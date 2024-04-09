package main

import (
	"fmt"
	"log/slog"
	"os"

	config "github.com/KaffeeMaschina/http-rest-api/internal"
	postgres "github.com/KaffeeMaschina/http-rest-api/internal/storage/postgres"
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

	postgres.New(cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.Database)
	//config.Config()
	//postgres.Connectiondb("postgres", "qwerty", "localhost", "5432", "postgres")
	//nats.ConnectionNS("test-cluster", "Nikita")

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
