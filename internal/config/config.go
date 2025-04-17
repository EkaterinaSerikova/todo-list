package config

import (
	"cmp"
	"flag"
	"os"
	"strconv"
)

//  загрузка и управление конфигурацией приложения

type Config struct {
	Host        string
	Port        int
	DbDsn       string
	MigratePath string
	Debug       bool
}

const (
	defaultHost            = "localhost"
	defaultPort        int = 8080
	defaultDbDst           = "postgres://user:password@localhost:5432/gt5?sslmode=disable"
	defaultMigratePath     = "migrations"
)

func ReadConfig() (*Config, error) {
	var cfg Config

	flag.StringVar(&cfg.Host, "host", defaultHost, "flag for explicit server host specification")
	flag.IntVar(&cfg.Port, "port", defaultPort, "flag for explicit server port specifications")
	flag.StringVar(&cfg.DbDsn, "db", defaultDbDst, "flag for explicit db connection string")
	flag.StringVar(&cfg.MigratePath, "migrate", defaultMigratePath, "flag for explicit migrate path")
	flag.BoolVar(&cfg.Debug, "debug", false, "flag for explicit debug mode")

	flag.Parse()

	if cfg.Host == "localhost" {
		cfg.Host = cmp.Or(os.Getenv("HOST"), cfg.Host)
	}

	if cfg.Port == 8080 {
		defPort := strconv.Itoa(cfg.Port)
		envPort := cmp.Or(os.Getenv("PORT"), defPort)

		port, err := strconv.Atoi(envPort)
		if err != nil {
			return nil, err
		}
		cfg.Port = port
	}

	return &cfg, nil
}
