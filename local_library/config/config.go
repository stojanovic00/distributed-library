package config

import (
	flags "github.com/jessevdk/go-flags"
)

type Config struct {
	Host   string `long:"host" env:"LOCAL_HOST_NS"`
	Port   string `long:"port" env:"LOCAL_PORT_NS"`
	DbName string `long:"db-name" env:"LOCAL_DB_NAME_NS"`
	DbHost string `long:"db-host" env:"LOCAL_DB_HOST_NS"`
	DbPort string `long:"db-port" env:"LOCAL_DB_PORT_NS"`
	DbUser string `long:"db-user" env:"LOCAL_DB_USER_NS"`
	DbPass string `long:"db-pass" env:"LOCAL_DB_PASS_NS"`

	CentralLibHost string `long:"central-host" env:"CENTRAL_HOST"`
	CentralLibPort string `long:"central-port" env:"CENTRAL_PORT"`
}

func LoadConfig() (Config, error) {
	var cfg Config
	parser := flags.NewParser(&cfg, flags.Default)
	_, err := parser.Parse()
	return cfg, err
}
