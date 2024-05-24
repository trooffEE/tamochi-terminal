package models

type User struct {
	Login string // same as username
}

func NewUser() *User {
	return &User{}
}
