package config

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/caarlos0/env/v11"
)

type Config struct {
	AppName     string `env:"APP_NAME" envDefault:"go-starter"`
	AppVersion  string `env:"APP_VERSION" envDefault:"dev"`
	Env         string `env:"ENV" envDefault:"development"`
	Port        string `env:"PORT" envDefault:"8080"`
	LogLevel    string `env:"LOG_LEVEL" envDefault:"info"`
	DatabaseURL string `env:"DATABASE_URL"`
}

func Load() (*Config, error) {
	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		return nil, fmt.Errorf("parse env: %w", err)
	}
	return cfg, nil
}

func NewLogger(cfg *Config) *slog.Logger {
	level := slog.LevelInfo
	switch cfg.LogLevel {
	case "debug":
		level = slog.LevelDebug
	case "warn":
		level = slog.LevelWarn
	case "error":
		level = slog.LevelError
	}

	opts := &slog.HandlerOptions{Level: level}
	var handler slog.Handler
	if cfg.Env == "development" {
		handler = slog.NewTextHandler(os.Stdout, opts)
	} else {
		handler = slog.NewJSONHandler(os.Stdout, opts)
	}

	return slog.New(handler).With(
		"app", cfg.AppName,
		"version", cfg.AppVersion,
		"env", cfg.Env,
	)
}
