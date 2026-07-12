package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ege-ayan/go-starter/internal/config"
	"github.com/ege-ayan/go-starter/internal/database"
	"github.com/ege-ayan/go-starter/internal/server"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		slog.Error("failed to load config", "error", err)
		os.Exit(1)
	}

	logger := config.NewLogger(cfg)
	slog.SetDefault(logger)

	var db *database.DB
	if cfg.DatabaseURL != "" {
		db, err = database.Connect(context.Background(), cfg.DatabaseURL)
		if err != nil {
			logger.Error("failed to connect to database", "error", err)
			os.Exit(1)
		}
		defer db.Close()
		logger.Info("connected to database")
	}

	srv := server.New(cfg, logger, db)

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	go func() {
		if err := srv.Start(); err != nil {
			slog.Error("server error", "error", err)
			stop()
		}
	}()

	<-ctx.Done()
	slog.Info("shutting down")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(shutdownCtx); err != nil {
		slog.Error("shutdown error", "error", err)
		os.Exit(1)
	}

	slog.Info("server stopped")
}
