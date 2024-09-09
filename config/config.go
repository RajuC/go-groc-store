package config

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type (
	Config struct {
		App `yaml:"app"`
		Log `yaml:"log"`
		Db  `yaml:"db"`
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
		DbName string `env-required:"false" yaml:"db_name" env:"DB_NAME"`
		DbPath string `env-required:"false" yaml:"db_path" env:"DB_PATH"`
		DbUrl  string `env-required:"false" yaml:"db_url" env:"DB_URL"`
	}
)

func NewConfigService() (*Config, error) {
	viperCfg := viper.New()
	viperCfg.SetConfigName("config")
	viperCfg.SetConfigType("yaml")
	viperCfg.AddConfigPath(".")
	cfg := &Config{}

	err := viperCfg.ReadInConfig()
	if err != nil {
		panic(err)
	}
	// fmt.Println(viperCfg)
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
	fmt.Println(cfg)

	return cfg, nil
}
