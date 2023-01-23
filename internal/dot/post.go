package dot

import (
	"net/url"
)

type CreatePost struct {
	Title   string
	Content string
	UserId  string
}

func FillingCreatePost(data url.Values) *CreatePost {
	return &CreatePost{
		Title:   data.Get("title"),
		Content: data.Get("content"),
		UserId:  data.Get("user_id"),
	}
}
