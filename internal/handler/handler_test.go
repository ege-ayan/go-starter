package handler_test

import (
	"encoding/json"
	"io"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ege-ayan/go-starter/internal/config"
	"github.com/ege-ayan/go-starter/internal/handler"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func newTestHandler(t *testing.T) *handler.Handler {
	t.Helper()

	cfg := &config.Config{
		AppName:    "go-starter",
		AppVersion: "test",
		Env:        "test",
	}
	logger := slog.New(slog.NewTextHandler(io.Discard, nil))

	return handler.New(cfg, logger, nil)
}

func TestHealth(t *testing.T) {
	h := newTestHandler(t)

	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	rec := httptest.NewRecorder()

	h.Health(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, "application/json", rec.Header().Get("Content-Type"))

	var body map[string]string
	require.NoError(t, json.NewDecoder(rec.Body).Decode(&body))
	assert.Equal(t, "ok", body["status"])
	assert.Equal(t, "skipped", body["database"])
}

func TestGreetDefault(t *testing.T) {
	h := newTestHandler(t)

	req := httptest.NewRequest(http.MethodGet, "/api/v1/hello", nil)
	rec := httptest.NewRecorder()

	h.Greet(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)

	var body map[string]string
	require.NoError(t, json.NewDecoder(rec.Body).Decode(&body))
	assert.Equal(t, "Hello, World!", body["message"])
}

func TestGreetWithName(t *testing.T) {
	h := newTestHandler(t)

	req := httptest.NewRequest(http.MethodGet, "/api/v1/hello?name=Go", nil)
	rec := httptest.NewRecorder()

	h.Greet(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)

	var body map[string]string
	require.NoError(t, json.NewDecoder(rec.Body).Decode(&body))
	assert.Equal(t, "Hello, Go!", body["message"])
}

func TestStatus(t *testing.T) {
	h := newTestHandler(t)

	req := httptest.NewRequest(http.MethodGet, "/api/v1/status", nil)
	rec := httptest.NewRecorder()

	h.Status(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)

	var body map[string]string
	require.NoError(t, json.NewDecoder(rec.Body).Decode(&body))
	assert.Equal(t, "go-starter", body["app"])
	assert.Equal(t, "test", body["version"])
	assert.Equal(t, "test", body["env"])
}
