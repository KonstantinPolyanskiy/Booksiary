package domain

import "github.com/google/uuid"

type UserRegistrationData struct {
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Login    string `json:"login"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type RegisteredUser struct {
	Name         string
	Surname      string
	Login        string
	PasswordHash string
	Email        string
}

type UserAccount struct {
	UUID     uuid.UUID `json:"UUID"`
	Login    string    `json:"login"`
	Password string    `json:"password"`
}

type UserPersonality struct {
	Name    string
	Surname string
	Email   string
}
