package main

import (
	"fmt"
	"os"

	"github.com/romanpovol/music_library/internal/config"

	"log"
	"log/slog"
)

func main() {
	cfg := config.LoadBackendConfig()

	fmt.Println(*cfg)

	log := setupLogger(cfg.LogLevel)
	log.Info(
		"starting server",
		slog.String("server", cfg.HTTPServer.Address),
	)

	// TODO: init storage

	// TODO: init router

	// TODO: run server
}

const (
	LevelDebug = "Debug"
	LevelInfo  = "Info"
	LevelWarn  = "Warn"
	LevelError = "Error"
)

func setupLogger(logLevel string) *slog.Logger {
	var logger *slog.Logger

	switch logLevel {
	case LevelDebug:
		logger = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case LevelInfo:
		logger = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	case LevelWarn:
		logger = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelWarn}),
		)
	case LevelError:
		logger = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelError}),
		)
	default:
		log.Fatalf("unexpected log level %s", logLevel)
	}

	return logger
}
