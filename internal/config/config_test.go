package config_test

import (
	"testing"

	"github.com/ege-ayan/go-starter/internal/config"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLoadDefaults(t *testing.T) {
	t.Setenv("APP_NAME", "")
	t.Setenv("APP_VERSION", "")
	t.Setenv("ENV", "")
	t.Setenv("PORT", "")
	t.Setenv("LOG_LEVEL", "")
	t.Setenv("DATABASE_URL", "")

	cfg, err := config.Load()
	require.NoError(t, err)

	assert.Equal(t, "go-starter", cfg.AppName)
	assert.Equal(t, "dev", cfg.AppVersion)
	assert.Equal(t, "development", cfg.Env)
	assert.Equal(t, "8080", cfg.Port)
	assert.Equal(t, "info", cfg.LogLevel)
}

func TestLoadFromEnv(t *testing.T) {
	t.Setenv("APP_NAME", "test-app")
	t.Setenv("APP_VERSION", "1.0.0")
	t.Setenv("ENV", "production")
	t.Setenv("PORT", "3000")
	t.Setenv("LOG_LEVEL", "debug")

	cfg, err := config.Load()
	require.NoError(t, err)

	assert.Equal(t, "test-app", cfg.AppName)
	assert.Equal(t, "1.0.0", cfg.AppVersion)
	assert.Equal(t, "production", cfg.Env)
	assert.Equal(t, "3000", cfg.Port)
	assert.Equal(t, "debug", cfg.LogLevel)
}

func TestNewLogger(t *testing.T) {
	cfg := &config.Config{
		AppName:    "go-starter",
		AppVersion: "test",
		Env:        "development",
		LogLevel:   "info",
	}

	logger := config.NewLogger(cfg)
	assert.NotNil(t, logger)
}
