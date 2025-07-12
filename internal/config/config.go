package config

import (
	"gopkg.in/yaml.v3"
	"os"
	"time"
)

type Config struct {
	PostgresDsn string        `yaml:"postgres_dsn"`
	TokenTTL    time.Duration `yaml:"token_ttl"`
	Env         string        `yaml:"env"`

	Grpc struct {
		Port    int           `yaml:"port"`
		Timeout time.Duration `yaml:"timeout"`
	}
}

func MustLoad() *Config {
	path := os.Getenv("CONFIG_PATH")

	if path == "" {
		panic("CONFIG_PATH environment variable not set")
	}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		panic("CONFIG_PATH does not exist")
	}

	data, err := os.ReadFile(path)
	if err != nil {
		panic("failed to read config file")
	}

	var cfg Config

	if err := yaml.Unmarshal(data, &cfg); err != nil {
		panic("failed to parse config file")
	}

	return &cfg
}
