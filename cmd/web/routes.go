package main

import (
	"net/http"

	"byryan.net/config"
)

func routes(app *config.Application) http.Handler {
	mux := http.NewServeMux()
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	mux.Handle("GET /{$}", home(app))
	mux.Handle("GET /post/{slug}", postView(app))
	mux.Handle("GET /post/{postID}/comments", commentsView(app))
	mux.Handle("POST /comment/create", createComment(app))
	mux.Handle("GET /about", about(app))
	mux.Handle("GET /arcade", arcade(app))
	return recoverPanic(logRequest(commonHeaders(mux), app), app)
}
