package model

import "time"

type User struct {
	Id        int
	Name      string
	Password  string
	BirthDay  time.Time
	CreateAt  time.Time
	DeletedAt time.Time
}
