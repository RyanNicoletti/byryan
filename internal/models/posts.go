package models

import (
	"database/sql"
	"errors"
	"html/template"
	"os"
	"path/filepath"
	"time"

	"github.com/lib/pq"
)

type Post struct {
	ID      string
	Title   string
	Slug    string
	Content template.HTML
	Tags    []string
	Created time.Time
	Updated time.Time
}

type PostModel struct {
	DB *sql.DB
}

func (pm *PostModel) GetById(id string) (Post, error) {
	stmt := `SELECT id, title, slug, tags, created, updated FROM posts WHERE id=$1`
	row := pm.DB.QueryRow(stmt, id)
	var p Post
	err := row.Scan(&p.ID, &p.Title, &p.Slug, pq.Array(&p.Tags), &p.Created, &p.Updated)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return Post{}, ErrNoRecord
		}
		return Post{}, err
	}
	filename := filepath.Join("ui", "html", "posts", p.Slug+".html")
	content, err := os.ReadFile(filename)
	if err != nil {
		return Post{}, ErrNoRecord
	}
	p.Content = template.HTML(string(content))
	return p, nil
}

func (pm *PostModel) GetBySlug(slug string) (Post, error) {
	stmt := `SELECT id, title, slug, tags, created, updated FROM posts WHERE slug=$1`
	row := pm.DB.QueryRow(stmt, slug)
	var p Post
	err := row.Scan(&p.ID, &p.Title, &p.Slug, pq.Array(&p.Tags), &p.Created, &p.Updated)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return Post{}, ErrNoRecord
		}
		return Post{}, err
	}
	filename := filepath.Join("ui", "html", "posts", slug+".html")
	content, err := os.ReadFile(filename)
	if err != nil {
		return Post{}, ErrNoRecord
	}
	p.Content = template.HTML(string(content))
	return p, nil
}

func (pm *PostModel) GetAll() ([]Post, error) {
	stmt := `SELECT * FROM posts`
	rows, err := pm.DB.Query(stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var posts []Post
	for rows.Next() {
		var p Post
		err = rows.Scan(&p.ID, &p.Title, &p.Slug, pq.Array(&p.Tags), &p.Created, &p.Updated)
		if err != nil {
			return nil, err
		}
		posts = append(posts, p)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return posts, nil
}
