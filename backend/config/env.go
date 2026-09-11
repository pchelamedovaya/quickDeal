package config

import (
	"os"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	Port  string
	DBDsn string
}

func Load() (*AppConfig, error) {
	if err := godotenv.Load(); err != nil && !os.IsNotExist(err) {
		return nil, err
	}

	return &AppConfig{
		Port:  os.Getenv("PORT"),
		DBDsn: os.Getenv("DB_DSN"),
	}, nil
}
