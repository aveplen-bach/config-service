package config

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
)

type (
	Config struct {
		ServerConfig ServerConfig `yaml:"server"`
	}

	ServerConfig struct {
		Addr  string  `yaml:"addr" env-defaul:":8082"`
		Debug *string `env:"GIN_MODE"`
	}
)

func ReadConfig(filename string, cfg *Config) error {
	if err := cleanenv.ReadConfig(filename, cfg); err != nil {
		return fmt.Errorf("could not read config: %w", err)
	}
	return nil
}
