package main

import (
	"errors"
	"os"
)

type Config struct {
	DB string
}

var ErrNoDB = errors.New("No database defined")

func configure() (Config, error) {
	cfg := Config{
		DB: os.Getenv("DATABASE"),
	}
	if cfg.DB == "" {
		return cfg, ErrNoDB
	}

	return cfg, nil
}
