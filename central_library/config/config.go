package config

import (
	flags "github.com/jessevdk/go-flags"
)

type Config struct {
	Host   string `long:"host" env:"CENTRAL_HOST"`
	Port   string `long:"port" env:"CENTRAL_PORT"`
	DbName string `long:"db-name" env:"CENTRAL_DB_NAME"`
	DbHost string `long:"db-host" env:"CENTRAL_DB_HOST"`
	DbPort string `long:"db-port" env:"CENTRAL_DB_PORT"`
	DbUser string `long:"db-user" env:"CENTRAL_DB_USER"`
	DbPass string `long:"db-pass" env:"CENTRAL_DB_PASS"`
}

func LoadConfig() (Config, error) {
	var cfg Config
	parser := flags.NewParser(&cfg, flags.Default)
	_, err := parser.Parse()
	return cfg, err
}
