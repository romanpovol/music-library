package main

import (
	"fmt"
	"os"

	"github.com/romanpovol/music-library/internal/config"

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

func ParseLevel(s string) (slog.Level, error) {
	var level slog.Level
	var err = level.UnmarshalText([]byte(s))
	return level, err
}

func setupLogger(logLevel string) *slog.Logger {
	var logger *slog.Logger

	level, err := ParseLevel(logLevel)

	if err == nil {
		logger = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: level}),
		)
	} else {
		log.Fatalf("unexpected log level %s", logLevel)
	}

	return logger
}
