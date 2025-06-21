package config

import (
	"log/slog"
)

type Application struct {
	Logger *slog.Logger
}

func NewApplication(logger *slog.Logger) *Application {
	return &Application{
		Logger: logger,
	}
}
