package models

import (
	"embed"
	"encoding/json"
	"fmt"
	"html/template"
	"io/fs"
	"path/filepath"
	"sort"
	"time"
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

type PostMeta struct {
	Title   string    `json:"title"`
	Slug    string    `json:"slug"`
	Tags    []string  `json:"tags"`
	Created time.Time `json:"created"`
	Updated time.Time `json:"updated"`
	Draft   bool      `json:"draft,omitempty"`
}

type PostModel struct {
	PostsFS embed.FS
}

func (p *PostModel) GetBySlug(slug string) (Post, error) {
	entries, err := fs.ReadDir(p.PostsFS, ".")
	if err != nil {
		return Post{}, err
	}

	for _, entry := range entries {
		if entry.IsDir() && entry.Name() == slug {
			return p.loadPost(slug)
		}
	}

	return Post{}, ErrNoRecord
}

func (p *PostModel) GetAll() ([]Post, error) {
	entries, err := fs.ReadDir(p.PostsFS, ".")
	if err != nil {
		return nil, err
	}

	var posts []Post
	for _, entry := range entries {
		if entry.IsDir() {
			post, err := p.loadPost(entry.Name())
			if err != nil {
				continue
			}
			posts = append(posts, post)
		}
	}

	sort.Slice(posts, func(i, j int) bool {
		return posts[i].Created.After(posts[j].Created)
	})

	return posts, nil
}

func (p *PostModel) GetById(id string) (Post, error) {
	return p.GetBySlug(id)
}

func (p *PostModel) loadPost(slug string) (Post, error) {
	metaPath := filepath.Join(slug, "meta.json")
	metaBytes, err := fs.ReadFile(p.PostsFS, metaPath)
	if err != nil {
		return Post{}, fmt.Errorf("error reading metadata for %s: %w", slug, err)
	}

	var meta PostMeta
	if err := json.Unmarshal(metaBytes, &meta); err != nil {
		return Post{}, fmt.Errorf("error parsing metadata for %s: %w", slug, err)
	}

	if meta.Draft {
		return Post{}, ErrNoRecord
	}

	contentPath := filepath.Join(slug, "content.html")
	contentBytes, err := fs.ReadFile(p.PostsFS, contentPath)
	if err != nil {
		return Post{}, fmt.Errorf("error reading content for %s: %w", slug, err)
	}

	return Post{
		ID:      meta.Slug,
		Title:   meta.Title,
		Slug:    meta.Slug,
		Content: template.HTML(contentBytes),
		Tags:    meta.Tags,
		Created: meta.Created,
		Updated: meta.Updated,
	}, nil
}
