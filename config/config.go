package config

import (
	"database/sql"
	"flag"
	"fmt"
	"html/template"
	"log/slog"

	"byryan.net/internal/models"
	_ "github.com/lib/pq"
)

type Config struct {
	Addr string
	DSN  string
}

type Application struct {
	Logger        *slog.Logger
	DB            *sql.DB
	TemplateCache map[string]*template.Template
	Posts         *models.PostModel
	Comments      *models.CommentModel
}

func Load() (*Config, error) {
	cfg := &Config{}

	flag.StringVar(&cfg.Addr, "addr", ":4000", "HTTP server address")
	flag.StringVar(&cfg.DSN, "dsn", "", "PostgreSQL data source name")

	flag.Parse()

	if err := cfg.validate(); err != nil {
		return nil, err
	}

	return cfg, nil
}

func (c *Config) validate() error {
	if c.DSN == "" {
		return fmt.Errorf("database DSN is required (use -dsn flag)")
	}
	return nil
}

func NewApplication(logger *slog.Logger, db *sql.DB, templateCache map[string]*template.Template) *Application {
	return &Application{
		Logger:        logger,
		DB:            db,
		TemplateCache: templateCache,
		Posts:         &models.PostModel{DB: db},
		Comments:      &models.CommentModel{DB: db},
	}
}
