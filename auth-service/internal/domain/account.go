package domain

import "github.com/google/uuid"

type UserRegistrationData struct {
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Login    string `json:"login"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type UserAccount struct {
	UUID         uuid.UUID
	Login        string `json:"login"`
	PasswordHash string `json:"passwordHash"`
	Role         int    `json:"role"`
}

type UserAccountResponse struct {
	UUID     uuid.UUID `json:"UUID"`
	Login    string    `json:"login"`
	Password string    `json:"passwordHash"`
}

type UserAccountDB struct {
	UUID         uuid.UUID
	Login        string
	PasswordHash string
}
type UserTokenData struct {
	UUID   uuid.UUID `json:"UUID"`
	RoleId int       `json:"roleId"`
}
