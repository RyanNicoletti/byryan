package main

import (
	"flag"
	"log/slog"
	"net/http"
	"os"

	"byryan.net/config"
)

func main() {
	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	app := config.NewApplication(logger)

	logger.Info("starting server", slog.String("addr", *addr))
	err := http.ListenAndServe(*addr, routes(app))
	logger.Error(err.Error())
	os.Exit(1)
}
