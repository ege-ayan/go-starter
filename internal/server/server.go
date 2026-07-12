package server

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/ege-ayan/go-starter/internal/config"
	"github.com/ege-ayan/go-starter/internal/database"
	"github.com/ege-ayan/go-starter/internal/handler"
)

type Server struct {
	cfg    *config.Config
	logger *slog.Logger
	http   *http.Server
}

func New(cfg *config.Config, logger *slog.Logger, db *database.DB) *Server {
	h := handler.New(cfg, logger, db)

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(30 * time.Second))

	r.Get("/health", h.Health)

	r.Route("/api/v1", func(r chi.Router) {
		r.Get("/hello", h.Greet)
		r.Get("/status", h.Status)
		r.Get("/db", h.DBInfo)
	})

	return &Server{
		cfg:    cfg,
		logger: logger,
		http: &http.Server{
			Addr:         ":" + cfg.Port,
			Handler:      r,
			ReadTimeout:  5 * time.Second,
			WriteTimeout: 10 * time.Second,
			IdleTimeout:  60 * time.Second,
		},
	}
}

func (s *Server) Handler() http.Handler {
	return s.http.Handler
}

func (s *Server) Start() error {
	s.logger.Info("starting server", "addr", s.http.Addr)
	if err := s.http.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return fmt.Errorf("listen and serve: %w", err)
	}
	return nil
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.http.Shutdown(ctx)
}
