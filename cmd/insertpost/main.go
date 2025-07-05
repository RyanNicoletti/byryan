package main

import (
	"context"
	"database/sql"
	"errors"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/lib/pq"
)

type tagsSlice []string

func (t *tagsSlice) Set(s string) error {
	*t = append(*t, s)
	return nil
}

// String is an implementation of the flag.Value interface
func (i *tagsSlice) String() string {
	return "wooo"
}

func isEmpty(flag string) bool {
	return strings.TrimSpace(flag) == ""
}

func validateFlags(title, slug, path, dsn string, tags tagsSlice) error {
	invalid := false

	if isEmpty(title) {
		invalid = true
		fmt.Fprintf(os.Stderr, "%s cannot be empty", title)
	}
	if isEmpty(slug) {
		invalid = true
		fmt.Fprintf(os.Stderr, "%s cannot be empty", slug)
	}
	if isEmpty(path) {
		invalid = true
		fmt.Fprintf(os.Stderr, "%s cannot be empty", path)
	}
	if isEmpty(dsn) {
		invalid = true
		fmt.Fprintf(os.Stderr, "%s cannot be empty", dsn)
	}
	for _, t := range tags {
		if isEmpty(t) {
			invalid = true
			fmt.Fprintf(os.Stderr, "%s cannot be empty", t)
		}
	}

	_, err := os.Stat(path)
	if err != nil {
		invalid = true
		if errors.Is(err, os.ErrNotExist) {
			fmt.Fprintf(os.Stderr, "file not found: %s", path)
		} else {
			fmt.Fprintf(os.Stderr, "something went wrong getting: %s", path)
			fmt.Fprintf(os.Stderr, "%v", err)
		}
	}

	if invalid {
		return errors.New("invalid flag('s) found, exiting")
	}
	return nil
}

func main() {
	var (
		title        string
		slug         string
		tags         tagsSlice
		path         string
		contentBytes []byte
		dsn          string
	)
	flag.StringVar(&title, "title", "", "Blog post title")
	flag.StringVar(&slug, "slug", "", "Blog post slug")
	flag.Var(&tags, "tags", "Tags for the blog post")
	flag.StringVar(&path, "path", "", "Path to html file containing the post")
	flag.StringVar(&dsn, "dsn", "", "Data source name for byryan database")
	flag.Parse()

	err := validateFlags(title, slug, path, dsn, tags)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(1)
	}

	contentBytes, err = os.ReadFile(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(1)
	}
	content := string(contentBytes)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(1)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err = db.PingContext(ctx)
	if err != nil {
		db.Close()
		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(1)
	}

	insertStmt := `INSERT into posts (title, slug, content, tags) 
				   VALUES($1, $2, $3, $4)
				   ON CONFLICT (slug) DO UPDATE SET
				   title = EXCLUDED.title,
				   content = EXCLUDED.content,
				   tags = EXCLUDED.tags`
	_, err = db.Exec(insertStmt, title, slug, content, pq.Array(tags))
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(1)
	}

	db.Close()
	fmt.Fprintf(os.Stdout, "%s", "Successfully inserted post into db")
}
