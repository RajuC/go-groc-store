package config

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestConfig(t *testing.T) {
	cfg, _ := NewConfigService()
	assert.Equal(t, cfg.App.Name, "go-groc-store")
	assert.Equal(t, cfg.App.Version, "1.0")
}
