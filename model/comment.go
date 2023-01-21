package model

type Comment struct {
	Id       int
	UserId   int
	PostId   int
	ParentId int
	Content  string
}
