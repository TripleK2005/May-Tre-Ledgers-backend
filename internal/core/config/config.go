package config

import (
	"errors"
	"os"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	Port        string
	DatabaseURL string

	JWTSecret string

	AccessTokenExpire  time.Duration
	RefreshTokenExpire time.Duration
}

func Load() (*Config, error) {
	_ = godotenv.Load()

	accessExpire, err := time.ParseDuration(
		os.Getenv("ACCESS_TOKEN_EXPIRE"),
	)
	if err != nil {
		return nil, errors.New("invalid ACCESS_TOKEN_EXPIRE")
	}

	refreshExpire, err := time.ParseDuration(
		os.Getenv("REFRESH_TOKEN_EXPIRE"),
	)
	if err != nil {
		return nil, errors.New("invalid REFRESH_TOKEN_EXPIRE")
	}

	cfg := Config{
		Port:               os.Getenv("PORT"),
		DatabaseURL:        os.Getenv("DATABASE_URL"),
		JWTSecret:          os.Getenv("JWT_SECRET"),
		AccessTokenExpire:  accessExpire,
		RefreshTokenExpire: refreshExpire,
	}

	if err := validateConfig(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

func validateConfig(cfg *Config) error {
	if cfg.Port == "" {
		return errors.New("missing PORT")
	}

	if cfg.DatabaseURL == "" {
		return errors.New("missing DATABASE_URL")
	}

	if cfg.JWTSecret == "" {
		return errors.New("missing JWT_SECRET")
	}

	if cfg.AccessTokenExpire <= 0 {
		return errors.New("invalid ACCESS_TOKEN_EXPIRE")
	}

	if cfg.RefreshTokenExpire <= 0 {
		return errors.New("invalid REFRESH_TOKEN_EXPIRE")
	}

	return nil
}
