package model

type IndexPage struct {
	User       User
	Categories []string
	Posts      []Post
}

type PostPage struct {
	User User
	Post Post
}
