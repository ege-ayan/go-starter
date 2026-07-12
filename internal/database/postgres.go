package database

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type DB struct {
	pool *pgxpool.Pool
}

func Connect(ctx context.Context, url string) (*DB, error) {
	cfg, err := pgxpool.ParseConfig(url)
	if err != nil {
		return nil, fmt.Errorf("parse database url: %w", err)
	}

	cfg.MaxConns = 10
	cfg.MinConns = 1
	cfg.MaxConnLifetime = time.Hour
	cfg.MaxConnIdleTime = 30 * time.Minute

	pool, err := pgxpool.NewWithConfig(ctx, cfg)
	if err != nil {
		return nil, fmt.Errorf("create pool: %w", err)
	}

	if err := pool.Ping(ctx); err != nil {
		pool.Close()
		return nil, fmt.Errorf("ping database: %w", err)
	}

	return &DB{pool: pool}, nil
}

func (db *DB) Ping(ctx context.Context) error {
	return db.pool.Ping(ctx)
}

func (db *DB) Close() {
	db.pool.Close()
}

func (db *DB) Version(ctx context.Context) (string, error) {
	var version string
	err := db.pool.QueryRow(ctx, "SELECT version()").Scan(&version)
	if err != nil {
		return "", fmt.Errorf("query version: %w", err)
	}
	return version, nil
}

func (db *DB) ListGreetings(ctx context.Context) ([]string, error) {
	rows, err := db.pool.Query(ctx, "SELECT name FROM greetings ORDER BY id")
	if err != nil {
		return nil, fmt.Errorf("query greetings: %w", err)
	}
	defer rows.Close()

	var names []string
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			return nil, fmt.Errorf("scan greeting: %w", err)
		}
		names = append(names, name)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iterate greetings: %w", err)
	}

	return names, nil
}
