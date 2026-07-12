package handler

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/ege-ayan/go-starter/internal/config"
	"github.com/ege-ayan/go-starter/internal/database"
)

type Handler struct {
	cfg    *config.Config
	logger *slog.Logger
	db     *database.DB
}

func New(cfg *config.Config, logger *slog.Logger, db *database.DB) *Handler {
	return &Handler{
		cfg:    cfg,
		logger: logger,
		db:     db,
	}
}

func writeJSON(w http.ResponseWriter, status int, payload any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(payload)
}
