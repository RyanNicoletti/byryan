package main

import (
	"errors"
	"net/http"

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
		data := newTemplateData()
		data.Posts = posts
		render(w, r, app, http.StatusOK, "home.tmpl", data)
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
		c, err := app.Comments.GetByPostId(p.ID)
		if err != nil {
			if errors.Is(err, models.ErrNoRecord) {
				http.NotFound(w, r)
			} else {
				serverError(app, w, r, err)
			}
			return
		}
		data := newTemplateData()
		data.Post = p
		data.Comments = c
		render(w, r, app, http.StatusOK, "post.tmpl", data)
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
