package main

import (
	"fmt"
	"net/http"
	"text/template"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	files := []string{"./ui/html/base.tmpl", "./ui/html/pages/home.tmpl", "./ui/html/partials/nav.tmpl"}
	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, r, err)
		return
	}
	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		app.serverError(w, r, err)
	}
}

func (app *application) postView(w http.ResponseWriter, r *http.Request) {
	p := r.PathValue("postID")

	fmt.Fprintf(w, "gotem %s", p)
}

func (app *application) createComment(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("saved n ew comment"))
}

func (app *application) commentsView(w http.ResponseWriter, r *http.Request) {

}

func (app *application) about(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("about"))
}

func (app *application) arcade(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("arcade"))
}
