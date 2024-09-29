package config

import (
	"fmt"
	"log/slog"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type (
	Config struct {
		App  `yaml:"app"`
		Log  `yaml:"log"`
		Db   `yaml:"db"`
		Http `yaml:"http"`
	}

	App struct {
		Name        string `env-required:"true" yaml:"name" env:"APP_NAME"`
		Description string `env-required:"true" yaml:"description" env:"APP_DESC"`
		Version     string `env-required:"true" yaml:"version" env:"APP_VERSION"`
		Environment string `env-required:"true" yaml:"environment" env:"APP_ENVIRONMENT"`
	}

	Log struct {
		Level string `env-required:"true" yaml:"level" env:"LOG_LEVEL"`
	}

	Db struct {
		DbName string `env-required:"true" yaml:"dbname" env:"DB_NAME"`
		DbPath string `env-required:"true" yaml:"dbpath" env:"DB_PATH"`
		DbUrl  string `env-required:"true" yaml:"dburl" env:"DB_URL"`
	}
	Http struct {
		Port string `env-required:"true" yaml:"port" env:"HTTP_PORT"`
	}
)

func NewConfigService(logger *slog.Logger, path string) (*Config, error) {
	cfg := &Config{}
	viperCfg := viper.New()
	viperCfg.SetConfigName("config")
	viperCfg.SetConfigType("yaml")
	viperCfg.AddConfigPath(path)

	err := viperCfg.ReadInConfig()
	if err != nil {
		panic(err)
	}
	err = viperCfg.Unmarshal(cfg)
	if err != nil {
		return nil, err
	}
	viperCfg.WatchConfig()
	viperCfg.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
		if err = viperCfg.Unmarshal(cfg); err != nil {
			fmt.Println(err)
		}
	})

	return cfg, nil
}
