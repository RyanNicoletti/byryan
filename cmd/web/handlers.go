package main

import (
	"errors"
	"fmt"
	"net/http"

	"byryan.net/config"
	"byryan.net/internal/models"
	"byryan.net/internal/validator"
)

type commentFormData struct {
	Name    string
	Website string
	Comment string
	validator.Validator
}

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
		data := newTemplateData(r)
		data.Posts = posts
		render(w, r, app, http.StatusOK, "home.tmpl", data)
	})
}

func postView(app *config.Application) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		slug := r.PathValue("slug")
		if slug == "" {
			http.NotFound(w, r)
			return
		}
		post, err := app.Posts.GetBySlug(slug)
		if err != nil {
			if errors.Is(err, models.ErrNoRecord) {
				http.NotFound(w, r)
			} else {
				serverError(app, w, r, err)
			}
			return
		}

		data := newTemplateData(r)
		data.Post = post
		render(w, r, app, http.StatusOK, "post.tmple", data)
	})
}

func createComment(app *config.Application) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()
		if err != nil {
			clientError(w, http.StatusBadRequest)
			return
		}
		formData := commentFormData{
			Name:    r.PostForm.Get("name"),
			Comment: r.PostForm.Get("comment"),
			Website: r.PostForm.Get("website"),
		}

		formData.ValidateNotBlank(formData.Name, "name", "This field cannot be blank")
		formData.ValidateNotBlank(formData.Comment, "comment", "This field cannot be blank")
		formData.ValidateLength(formData.Name, 100, "name", "This field cannot be more than 100 characters long")
		url := formData.ValidateUrl(formData.Website, "website")

		postId := r.PostForm.Get("post_id")

		if !formData.IsValid() {
			p, err := app.Posts.GetById(postId)
			if err != nil {
				serverError(app, w, r, err)
				return
			}

			c, err := app.Comments.GetByPostId(p.ID)
			if err != nil {
				serverError(app, w, r, err)
				return
			}

			data := newTemplateData(r)
			data.Post = p
			data.Comments = c
			data.Form = formData
			render(w, r, app, http.StatusUnprocessableEntity, "post.tmpl", data)
			return
		}

		p, err := app.Posts.GetById(postId)
		if err != nil {
			serverError(app, w, r, err)
			return
		}
		id, err := app.Comments.Insert(formData.Name, &url, formData.Comment, postId)
		if err != nil {
			serverError(app, w, r, err)
			return
		}
		http.Redirect(w, r, fmt.Sprintf("/post/%s#comment-%s", p.Slug, id), http.StatusSeeOther)
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
