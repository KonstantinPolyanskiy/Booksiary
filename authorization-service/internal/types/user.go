package types

import "time"

type Personality struct {
	Name    string `json:"name"`
	Surname string `json:"surname"`
}

type User struct {
	Uuid        string      `json:"id"`
	Personality Personality `json:"personality"`
	Email       string      `json:"email"`
}

type UserCreate struct {
	Id int `json:"id"`
	User
}

type SavingUser struct {
	User User
	Code MailAccessCodeData
}
type MailAccessCodeData struct {
	AccessCode int
	ExpiredAt  time.Time
}
type Code struct {
	Code int `json:"code"`
}
