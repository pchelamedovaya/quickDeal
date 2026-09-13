package config

import (
	"os"
	"time"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	Port            string
	DBDsn           string
	JWTAccessSecret string
	AccessTokenTTL  time.Duration
	RefreshTokenTTL time.Duration
}

func Load() (*AppConfig, error) {
	if err := godotenv.Load(); err != nil && !os.IsNotExist(err) {
		return nil, err
	}

	accessTTL, err := time.ParseDuration(os.Getenv("ACCESS_TOKEN_TTL"))
	if err != nil {
		return nil, err
	}

	refreshTTL, err := time.ParseDuration(os.Getenv("REFRESH_TOKEN_TTL"))
	if err != nil {
		return nil, err
	}

	return &AppConfig{
		Port:            os.Getenv("PORT"),
		DBDsn:           os.Getenv("DB_DSN"),
		JWTAccessSecret: os.Getenv("JWT_ACCESS_SECRET"),
		AccessTokenTTL:  accessTTL,
		RefreshTokenTTL: refreshTTL,
	}, nil
}
