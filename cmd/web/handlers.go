package main

import (
	"errors"
	"net/http"
	"text/template"

	"byryan.net/config"
	"byryan.net/internal/models"
)

func home(app *config.Application) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		posts, err := app.Posts.GetAll()
		if err != nil {
			if errors.Is(err, models.ErrNoRecord) {
				http.NotFound(w, r)
			} else {
				serverError(app, w, r, err)
			}
			return
		}
		data := struct{ Posts []models.Post }{Posts: posts}

		files := []string{"./ui/html/base.tmpl", "./ui/html/pages/home.tmpl", "./ui/html/partials/nav.tmpl"}
		ts, err := template.ParseFiles(files...)
		if err != nil {
			serverError(app, w, r, err)
			return
		}
		err = ts.ExecuteTemplate(w, "base", data)
		if err != nil {
			serverError(app, w, r, err)
		}
	})
}

func postView(app *config.Application) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		s := r.PathValue("slug")
		p, err := app.Posts.GetBySlug(s)
		if err != nil {
			if errors.Is(err, models.ErrNoRecord) {
				http.NotFound(w, r)
			} else {
				serverError(app, w, r, err)
			}
			return
		}
		data := struct{ Post models.Post }{Post: p}

		files := []string{"./ui/html/base.tmpl", "./ui/html/pages/post.tmpl", "./ui/html/partials/nav.tmpl"}
		ts, err := template.ParseFiles(files...)
		if err != nil {
			serverError(app, w, r, err)
			return
		}
		err = ts.ExecuteTemplate(w, "base", data)
		if err != nil {
			serverError(app, w, r, err)
		}
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
