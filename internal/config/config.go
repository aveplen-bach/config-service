package config

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/sirupsen/logrus"
)

type (
	Config struct {
		ServerConfig   ServerConfig   `yaml:"server"`
		DatabaseConfig ConfigDatabase `yaml:"database"`
	}

	ConfigDatabase struct {
		Port     string `yaml:"port" env:"DBPORT" env-default:"5432"`
		Host     string `yaml:"host" env:"DBHOST" env-default:"localhost"`
		Name     string `yaml:"name" env:"DBNAME" env-default:"postgres"`
		User     string `yaml:"user" env:"DBUSER" env-default:"user"`
		Password string `yaml:"password" env:"DBPASSWORD"`
	}

	ServerConfig struct {
		Addr string `yaml:"addr" env:"GRPC_LISTEN_ADDR" env-defaul:":8082"`
	}
)

func ReadConfig(filename string, cfg *Config) error {
	logrus.Infof("reading config from %s", filename)
	if err := cleanenv.ReadConfig(filename, cfg); err != nil {
		return fmt.Errorf("could not read config: %w", err)
	}

	logrus.Info("reading env")
	if err := cleanenv.ReadEnv(cfg); err != nil {
		return fmt.Errorf("could not read config: %w", err)
	}
	return nil
}
