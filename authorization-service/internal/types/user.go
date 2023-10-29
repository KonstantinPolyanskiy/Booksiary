package types

type User struct {
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Email   string `json:"email"`
}

type UserCreate struct {
	Id int `json:"id"`
	User
}
