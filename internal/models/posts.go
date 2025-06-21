package models

import (
	"database/sql"
	"time"
)

type Post struct {
	ID      int
	Title   string
	Slug    string
	Created time.Time
	Conent  string
	Updated time.Time
}

type PostModel struct {
	DB *sql.DB
}

func (p *PostModel) Insert(title, slug, content string) (int, error) {
	return 0, nil
}

func (p *PostModel) Get(id int) (*Post, error) {
	return nil, nil
}

func (p *PostModel) GetBySlug(slug string) (*Post, error) {
	return nil, nil
}

func (p *PostModel) Latest(limit int) ([]*Post, error) {
	return nil, nil
}
