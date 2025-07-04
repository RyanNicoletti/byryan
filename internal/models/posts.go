package models

import (
	"database/sql"
	"errors"
	"html/template"
	"time"

	"github.com/lib/pq"
)

type Post struct {
	ID          string
	Title       string
	Slug        string
	Content     template.HTML
	Tags        []string
	PublishedAt *time.Time
	Created     time.Time
	Updated     time.Time
}

type PostModel struct {
	DB *sql.DB
}

func (p *PostModel) GetById(id string) (Post, error) {
	stmt := `SELECT id, title, slug, content, tags, published_at, created, updated FROM posts WHERE id=$1`
	row := p.DB.QueryRow(stmt, id)
	var post Post
	err := row.Scan(&post.ID, &post.Title, &post.Slug, &post.Content, pq.Array(&post.Tags), &post.PublishedAt, &post.Created, &post.Updated)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return Post{}, ErrNoRecord
		}
		return Post{}, err
	}
	return post, nil
}

func (p *PostModel) GetBySlug(slug string) (Post, error) {
	stmt := `SELECT id, title, slug, content, tags, published_at, created, updated 
             FROM posts WHERE slug = $1`
	row := p.DB.QueryRow(stmt, slug)
	var post Post
	err := row.Scan(&post.ID, &post.Title, &post.Slug, &post.Content, pq.Array(&post.Tags),
		&post.PublishedAt, &post.Created, &post.Updated)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return Post{}, ErrNoRecord
		}
		return Post{}, err
	}
	return post, nil
}

func (p *PostModel) GetAll() ([]Post, error) {
	stmt := `SELECT id, title, slug, content, tags, published_at, created, updated 
             FROM posts 
             WHERE published_at IS NOT NULL 
             ORDER BY created DESC`
	rows, err := p.DB.Query(stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []Post
	for rows.Next() {
		var post Post
		err = rows.Scan(&post.ID, &post.Title, &post.Slug, &post.Content, pq.Array(&post.Tags),
			&post.PublishedAt, &post.Created, &post.Updated)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return posts, nil
}

func (p *PostModel) Insert(title, slug, content string, tags []string) (string, error) {
	stmt := `INSERT INTO posts (title, slug, content, tags, published_at) VALUES ($1, $2, $3, $4, NOW())`
	result, err := p.DB.Exec(stmt, title, slug, content, pq.Array(tags))
	if err != nil {
		return "", err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return "", err
	}
	return string(id), nil
}
