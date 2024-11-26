package config

import (
	"log"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	LogLevel    string `env:"LOG_LEVEL" env-default:"Debug"`
	StoragePath string `env:"STORAGE_PATH" env-required:"true"`
	HTTPServer
}

type HTTPServer struct {
	Address string        `env:"HTTP_ADDRESS" env-default:"localhost:8080"`
	Timeout time.Duration `env:"HTTP_REQUEST_TIMEOUT" env-default:"4s"`
}

func LoadBackendConfig() *Config {
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		log.Fatal("CONFIG_PATH not found in environment variables")
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file does not exist: %s", configPath)
	}

	var cfg Config
	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalf("cannot read config file %s, error: %s", configPath, err)
	}

	return &cfg
}
