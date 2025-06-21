package config

import (
	"database/sql"
	"log/slog"

	"byryan.net/internal/models"
)

type Application struct {
	Logger *slog.Logger
	DB     *sql.DB
	Posts  *models.PostModel
}

func NewApplication(logger *slog.Logger, db *sql.DB) *Application {
	return &Application{
		Logger: logger,
		DB:     db,
		Posts:  &models.PostModel{DB: db},
	}
}
