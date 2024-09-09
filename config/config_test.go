package config

import (
	"go-groc-store/pkg/log"
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestConfig(t *testing.T) {
	logger := log.NewLoggerService()
	cfg, _ := NewConfigService(logger, ".")
	assert.Equal(t, cfg.App.Name, "go-groc-store")
	assert.Equal(t, cfg.App.Version, "1.0")
}
