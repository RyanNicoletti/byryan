package main

import (
	"bytes"
	"fmt"
	"html/template"
	"io/fs"
	"net/http"
	"path/filepath"
	"time"

	"byryan.net/config"
	"byryan.net/internal/models"
	"byryan.net/ui"
	"github.com/justinas/nosurf"
)

type templateData struct {
	CurrentYear int
	Post        models.Post
	Posts       []models.Post
	Comments    []models.Comment
	Form        any
	CSRFToken   string
}

func newTemplateData(r *http.Request) templateData {
	return templateData{
		CurrentYear: time.Now().Year(),
		CSRFToken:   nosurf.Token(r),
	}
}

func newTemplateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	pages, err := fs.Glob(ui.Files, "html/pages/*.tmpl")
	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		patterns := []string{"html/base.tmpl", "html/partials/*.tmpl", page}
		ts, err := template.New(name).ParseFS(ui.Files, patterns...)
		if err != nil {
			return nil, err
		}
		cache[name] = ts
	}
	return cache, nil
}

func render(w http.ResponseWriter, r *http.Request, app *config.Application, status int, page string, data templateData) {
	ts, ok := app.TemplateCache[page]
	if !ok {
		err := fmt.Errorf("the template %s does not exist", page)
		serverError(app, w, r, err)
		return
	}

	buf := new(bytes.Buffer)
	err := ts.ExecuteTemplate(buf, "base", data)
	if err != nil {
		serverError(app, w, r, err)
		return
	}

	w.WriteHeader(status)
	buf.WriteTo(w)
}
