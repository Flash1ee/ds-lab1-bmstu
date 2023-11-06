package config

import (
	"flag"
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Host  string `toml:"host"`
	Port  string `toml:"http_port"`
	Level string `toml:"log_level"`
}

type PostgresConfig struct {
	Host     string `env:"HOST"`
	Port     string `env:"PORT"`
	Database string `env:"DB"`
	User     string `env:"USER"`
	Password string `env:"PASSWORD"`
}

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "./configs/config.toml", "path to config file")
}

func NewConfig() (*Config, *PostgresConfig, error) {
	flag.Parse()
	cfg, dbCfg := &Config{}, &PostgresConfig{}
	err := cleanenv.ReadConfig(configPath, cfg)
	if err != nil {
		return nil, nil, fmt.Errorf("env config error: %w", err)
	}
	err = cleanenv.ReadEnv(dbCfg)
	if err != nil {
		return nil, nil, fmt.Errorf("config error: %w", err)
	}

	return cfg, dbCfg, nil
}
