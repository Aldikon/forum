package model

import (
	"net/url"
	"strings"
	"time"
)

type CreateComment struct {
	UserID    int64
	PostID    int64
	ParentID  int64
	Content   string
	CreateAtt time.Time
}

func (c *CreateComment) ParseForm(form url.Values) {
	c.PostID = atoi64(form.Get("post_id"))
	c.ParentID = atoi64(form.Get("parent_id"))
	c.Content = strings.TrimSpace(form.Get("content"))
}

func (c *CreateComment) Validate() error {
	if c.UserID <= 0 {
		return &fillingError{"Invalid UserID"}
	}
	if c.PostID <= 0 {
		return &fillingError{"Invalid PostID"}
	}
	if c.Content == "" {
		return &fillingError{"Content cannot be empty"}
	}
	return nil
}

type Comment struct {
	ID        int64
	UserName  string
	ParentID  int64
	Content   string
	CreateAtt time.Time
	Reaction
}
