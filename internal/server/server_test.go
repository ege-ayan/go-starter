package server_test

import (
	"encoding/json"
	"io"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ege-ayan/go-starter/internal/config"
	"github.com/ege-ayan/go-starter/internal/server"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestServerRoutes(t *testing.T) {
	cfg := &config.Config{
		AppName:    "go-starter",
		AppVersion: "test",
		Env:        "test",
		Port:       "0",
	}
	logger := slog.New(slog.NewTextHandler(io.Discard, nil))

	srv := server.New(cfg, logger, nil)
	ts := httptest.NewServer(srv.Handler())
	t.Cleanup(ts.Close)

	tests := []struct {
		name       string
		path       string
		wantStatus int
		checkBody  func(t *testing.T, body map[string]string)
	}{
		{
			name:       "health",
			path:       "/health",
			wantStatus: http.StatusOK,
			checkBody: func(t *testing.T, body map[string]string) {
				assert.Equal(t, "ok", body["status"])
				assert.Equal(t, "skipped", body["database"])
			},
		},
		{
			name:       "hello default",
			path:       "/api/v1/hello",
			wantStatus: http.StatusOK,
			checkBody: func(t *testing.T, body map[string]string) {
				assert.Equal(t, "Hello, World!", body["message"])
			},
		},
		{
			name:       "hello with name",
			path:       "/api/v1/hello?name=Docker",
			wantStatus: http.StatusOK,
			checkBody: func(t *testing.T, body map[string]string) {
				assert.Equal(t, "Hello, Docker!", body["message"])
			},
		},
		{
			name:       "status",
			path:       "/api/v1/status",
			wantStatus: http.StatusOK,
			checkBody: func(t *testing.T, body map[string]string) {
				assert.Equal(t, "go-starter", body["app"])
				assert.Equal(t, "test", body["version"])
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := http.Get(ts.URL + tt.path)
			require.NoError(t, err)
			defer resp.Body.Close()

			assert.Equal(t, tt.wantStatus, resp.StatusCode)

			var body map[string]string
			require.NoError(t, json.NewDecoder(resp.Body).Decode(&body))
			tt.checkBody(t, body)
		})
	}
}
