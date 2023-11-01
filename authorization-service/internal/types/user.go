package types

import "time"

type Personality struct {
	Name    string `json:"name"`
	Surname string `json:"surname"`
}
type UserCreateResponse struct {
	Personality
	Email string `json:"email"`
}
type UserCreate struct {
	Id int `json:"id"`
	User
}

type User struct {
	Uuid        string      `json:"-"`
	Personality Personality `json:"personality"`
	Email       string      `json:"email"`
}

type SavingUser struct {
	User User               `json:"user"`
	Code MailAccessCodeData `json:"-"`
}

type MailAccessCodeData struct {
	AccessCode int
	ExpiredAt  time.Time
}

type Code struct {
	Code int `json:"code"`
}
