package model

import (
	"net/url"
	"strings"
	"time"
)

type CreatePost struct {
	UserID     int64
	Title      string
	Content    string
	Categories []string
	CreateAtt  time.Time
}

func (c *CreatePost) ParseForm(form url.Values) {
	c.Title = strings.TrimSpace(form.Get("title"))
	c.Content = strings.TrimSpace(form.Get("content"))
	c.Categories = form["categories"]
}

func (c *CreatePost) Validate() error {
	if c.UserID <= 0 {
		return &fillingError{"Invalid UserID"}
	}
	if c.Title == "" {
		return &fillingError{"Title cannot be empty"}
	}
	if len([]rune(c.Title)) < 5 {
		return &fillingError{"Title must be at least 5 characters"}
	}
	if c.Content == "" {
		return &fillingError{"Content cannot be empty"}
	}
	if len([]rune(c.Content)) < 10 {
		return &fillingError{"Content must be at least 10 characters"}
	}
	if len(c.Categories) == 0 {
		return &fillingError{"Categories cannot be empty"}
	}
	return nil
}

type Post struct {
	ID          int64
	UserName    string
	Title       string
	Content     string
	Categories  []string
	CreateAtt   time.Time
	Comments    []Comment
	SumComments int64
	Reaction
}
