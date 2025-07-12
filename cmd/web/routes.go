package main

import (
	"net/http"

	"byryan.net/config"
	"byryan.net/ui"
)

type MiddlewareChain struct {
	middlewares []func(http.Handler) http.Handler
}

func (mc MiddlewareChain) Then(h http.Handler) http.Handler {
	return middlewareChain(h, mc.middlewares...)
}

func NewChain(middlewares ...func(http.Handler) http.Handler) MiddlewareChain {
	return MiddlewareChain{middlewares: middlewares}
}

func middlewareChain(h http.Handler, middlewares ...func(http.Handler) http.Handler) http.Handler {
	for i := len(middlewares) - 1; i >= 0; i-- {
		h = middlewares[i](h)
	}
	return h
}

func routes(app *config.Application) http.Handler {
	mux := http.NewServeMux()

	mux.Handle("GET /static/", http.FileServerFS(ui.Files))

	mux.Handle("GET /{$}", home(app))
	mux.Handle("GET /post/{slug}", postView(app))
	mux.Handle("POST /comment/create", createComment(app))
	mux.Handle("GET /about", about(app))
	mux.Handle("GET /links", links(app))

	standard := NewChain(func(h http.Handler) http.Handler { return recoverPanic(h, app) },
		func(h http.Handler) http.Handler { return logRequest(h, app) },
		func(h http.Handler) http.Handler { return commonHeaders(h) },
		func(h http.Handler) http.Handler { return noSurf(h) })

	return standard.Then(mux)
}
