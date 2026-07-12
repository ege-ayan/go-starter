package database_test

import (
	"context"
	"testing"
	"time"

	"github.com/ege-ayan/go-starter/internal/database"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestConnectInvalidURL(t *testing.T) {
	_, err := database.Connect(context.Background(), "not-a-valid-url")
	require.Error(t, err)
}

func TestConnectUnreachable(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	_, err := database.Connect(ctx, "postgres://postgres:postgres@127.0.0.1:59999/nope?sslmode=disable")
	require.Error(t, err)
	assert.Contains(t, err.Error(), "ping database")
}
