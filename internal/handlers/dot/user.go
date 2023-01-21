package dot

import (
	"net/url"
)

type User struct {
	Name     string
	Password string
}

func (u *User) Filling(data url.Values) {
	u.Name = data.Get("name")
	u.Password = data.Get("password")
}
