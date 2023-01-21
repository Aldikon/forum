package model

import "time"

type Post struct {
	Id        int
	Title     string
	Content   string
	UserId    int
	Category  []string
	CreateAt  time.Time
	UpdatedAt time.Time
}
