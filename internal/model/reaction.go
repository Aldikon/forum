package model

import "net/url"

type Reaction struct {
	Like    int64
	DisLike int64
}

type CreateReactionPost struct {
	UserID int64
	PostID int64
	Type   int64
}

func (c *CreateReactionPost) ParseForm(form url.Values) {
	c.PostID = atoi64(form.Get("post_id"))
	c.Type = atoi64(form.Get("reac"))
}

type CreateReactionComment struct {
	UserID    int64
	CommentID int64
	Type      int64
}

func (c *CreateReactionComment) ParseForm(form url.Values) {
	c.CommentID = atoi64(form.Get("comment_id"))
	c.Type = atoi64(form.Get("reac"))
}
