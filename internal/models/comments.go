package models

import (
	"database/sql"
	"time"
)

type Comment struct {
	ID      string
	Name    string
	Website *string
	Content string
	PostID  string
	Created time.Time
}

type CommentModel struct {
	DB *sql.DB
}

func (cm *CommentModel) GetByPostId(postID string) ([]Comment, error) {
	stmt := `SELECT id, name, website, content, post_id, created FROM comments WHERE post_id=$1`
	rows, err := cm.DB.Query(stmt, postID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var comments []Comment
	for rows.Next() {
		var c Comment
		err = rows.Scan(&c.ID, &c.Name, &c.Website, &c.Content, &c.PostID, &c.Created)
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

func (cm *CommentModel) Insert(name string, website *string, content string, postID string) (string, error) {
	stmt := `INSERT INTO comments (name, website, content, post_id, created) VALUES ($1, $2, $3, $4, NOW()) RETURNING id`
	var id string
	err := cm.DB.QueryRow(stmt, name, website, content, postID).Scan(&id)
	if err != nil {
		return "", err
	}
	return id, nil
}
