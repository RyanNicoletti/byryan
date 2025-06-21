package models

import (
	"database/sql"
	"errors"
	"time"
)

type Post struct {
	ID      int
	Title   string
	Slug    string
	Content string
	Created time.Time
	Updated time.Time
}

type PostModel struct {
	DB *sql.DB
}

func (pm *PostModel) GetBySlug(slug string) (Post, error) {
	stmt := `SELECT id, title, slug, content, created, updated FROM posts WHERE slug=$1`
	row := pm.DB.QueryRow(stmt, slug)
	var p Post
	err := row.Scan(&p.ID, &p.Title, &p.Slug, &p.Content, &p.Created, &p.Updated)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return Post{}, ErrNoRecord
		}
		return Post{}, err
	}
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
		err = rows.Scan(&p.ID, &p.Title, &p.Slug, &p.Content, &p.Created, &p.Updated)
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
