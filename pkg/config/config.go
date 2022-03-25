package config

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	HTTP     `yaml:"http"`
	GRPC     `yaml:"grpc"`
	Postgres `yaml:"postgres"`
}

type HTTP struct {
	Port string `env-required:"true" yaml:"port"`
}

type GRPC struct {
	Port string `env-required:"true" yaml:"port"`
}

type Postgres struct {
	DSN         string `env-required:"true" env:"PG_DSN"`
	MaxPoolSize int    `evn-required:"true" yaml:"maxPoolSize"`
}

func New(confFile string) (*Config, error) {
	if confFile == "" {
		return nil, fmt.Errorf("confFile can't be empty")
	}

	if _, err := os.Stat(".env"); err == nil {
		err = godotenv.Load()
		if err != nil {
			return nil, err
		}
	}

	config := &Config{}
	err := cleanenv.ReadConfig(confFile, config)
	if err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}

	err = cleanenv.ReadEnv(config)
	if err != nil {
		return nil, err
	}

	return config, nil
}
