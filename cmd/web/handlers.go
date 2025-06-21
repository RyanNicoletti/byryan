package main

import (
	"fmt"
	"net/http"
	"text/template"

	"byryan.net/config"
)

func home(app *config.Application) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		files := []string{"./ui/html/base.tmpl", "./ui/html/pages/home.tmpl", "./ui/html/partials/nav.tmpl"}
		ts, err := template.ParseFiles(files...)
		if err != nil {
			serverError(app, w, r, err)
			return
		}
		err = ts.ExecuteTemplate(w, "base", nil)
		if err != nil {
			serverError(app, w, r, err)
		}
	})
}

func postView(app *config.Application) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		p := r.PathValue("postID")
		fmt.Fprintf(w, "gotem %s", p)
	})
}

func createComment(app *config.Application) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("saved new comment"))
	})
}

func commentsView(app *config.Application) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	})
}

func about(app *config.Application) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("about"))
	})
}

func arcade(app *config.Application) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("arcade"))
	})
}
