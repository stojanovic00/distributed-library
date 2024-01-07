package config

import (
	flags "github.com/jessevdk/go-flags"
)

type Config struct {
	Host   string `long:"host" env:"LOCAL_HOST"`
	Port   string `long:"port" env:"LOCAL_PORT"`
	DbName string `long:"db-name" env:"LOCAL_DB_NAME"`
	DbHost string `long:"db-host" env:"LOCAL_DB_HOST"`
	DbPort string `long:"db-port" env:"LOCAL_DB_PORT"`
	DbUser string `long:"db-user" env:"LOCAL_DB_USER"`
	DbPass string `long:"db-pass" env:"LOCAL_DB_PASS"`

	CityPath string `long:"city-path" env:"CITY_PATH"`

	CentralLibHost string `long:"central-host" env:"CENTRAL_HOST"`
	CentralLibPort string `long:"central-port" env:"CENTRAL_PORT"`
}

func LoadConfig() (Config, error) {
	var cfg Config
	parser := flags.NewParser(&cfg, flags.Default)
	_, err := parser.Parse()
	return cfg, err
}
