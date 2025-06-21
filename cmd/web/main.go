package main

import (
	"database/sql"
	"flag"
	"log/slog"
	"net/http"
	"os"

	"byryan.net/config"
)

func main() {
	addr := flag.String("addr", ":4000", "HTTP network address")
	dsn := flag.String("dsn", "postgres://ryannicoletti@localhost/byryan?sslmode=disable", "PostgreSQL data source name")
	flag.Parse()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	db, err := openDB(*dsn)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	app := config.NewApplication(logger, db)

	logger.Info("starting server", slog.String("addr", *addr))
	err = http.ListenAndServe(*addr, routes(app))
	logger.Error(err.Error())
	os.Exit(1)
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("sqlite", dsn)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, err

	}
	return db, nil
}
