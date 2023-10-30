package types

import "time"

type User struct {
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Email   string `json:"email"`
}

type UserCreate struct {
	Id int `json:"id"`
	User
}

type SaveUser struct {
	CodeData
	User
}
type CodeData struct {
	Code      int
	ExpiredAt time.Time
}
type Code int
