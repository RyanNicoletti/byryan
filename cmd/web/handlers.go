package main

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"byryan.net/config"
	"byryan.net/internal/models"
)

type commentFormData struct {
	Name        string
	Website     string
	Comment     string
	FieldErrors map[string]string
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
		data.Form = commentFormData{}
		render(w, r, app, http.StatusOK, "post.tmpl", data)
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
			Name:        r.PostForm.Get("name"),
			Comment:     r.PostForm.Get("comment"),
			FieldErrors: map[string]string{},
		}

		postId := r.PostForm.Get("post_id")

		website, err := url.Parse(r.PostForm.Get("website"))
		if err != nil {
			formData.FieldErrors["website"] = "Invalid URL"
		}

		formData.Website = website.String()

		if !strings.HasPrefix(formData.Website, "http://") && !strings.HasPrefix(formData.Website, "https://") {
			formData.Website = "https://" + formData.Website
		}

		if strings.TrimSpace(formData.Name) == "" {
			formData.FieldErrors["name"] = "This field cannot be blank"
		}

		if strings.TrimSpace(formData.Comment) == "" {
			formData.FieldErrors["comment"] = "This field cannot be blank"
		}

		if len(formData.FieldErrors) > 0 {
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

			data := newTemplateData()
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
		id, err := app.Comments.Insert(formData.Name, formData.Website, formData.Comment, postId)
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
