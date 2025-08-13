package models

import (
	"database/sql"
	"time"
)

type Comment struct {
	ID       string
	Name     string
	Website  *string
	Content  string
	PostSlug string
	Created  time.Time
}

type CommentModel struct {
	DB *sql.DB
}

func (c *CommentModel) GetByPostSlug(postSlug string) ([]Comment, error) {
	stmt := `SELECT id, name, website, content, post_slug, created FROM comments WHERE post_slug=$1 ORDER BY created ASC`
	rows, err := c.DB.Query(stmt, postSlug)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var comments []Comment
	for rows.Next() {
		var c Comment
		err = rows.Scan(&c.ID, &c.Name, &c.Website, &c.Content, &c.PostSlug, &c.Created)
		if err != nil {
			return nil, err
		}
		comments = append(comments, c)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return comments, nil
}

func (c *CommentModel) Insert(name string, website *string, content string, postSlug string) (string, error) {
	stmt := `INSERT INTO comments (name, website, content, post_slug, created) VALUES ($1, $2, $3, $4, NOW()) RETURNING id`
	var id string
	err := c.DB.QueryRow(stmt, name, website, content, postSlug).Scan(&id)
	if err != nil {
		return "", err
	}
	return id, nil
}
