package service

import (
	"Booksiary/user-service/internal/domain"
	"github.com/google/uuid"
)

type Sender interface {
	SendCode(code int) error
}
type Creator interface {
	Create(user domain.RegisteredUser) (uuid.UUID, error)
}
type UserProvider interface {
	LoginOrEmailExist(login, password string) error
}
type CodeProvider interface {
	Add(code string, user domain.RegisteredUser) error
	Get(code string) error
}
type SignUpService struct {
	Creator
	UserProvider
}

func NewSignUpService() *SignUpService {
	return &SignUpService{}
}

func (s *SignUpService) SignUp(data domain.UserRegistrationData) (uuid.UUID, error) {
	//TODO: получить пользователя
	//TODO: проверить занят ли логин и пароль
	//TODO: сгенерировать код, отправить на почту, ждать пока он придет
	//TODO: если код верен, сохранить пользователя в базу данных
	//TODO: отправить логин/пароль вместе c UUID на сторонние сервисы
	err := s.UserProvider.LoginOrEmailExist(data.Login, data.Email)
	if err != nil {
		return uuid.UUID{}, err
	}

	user := domain.RegisteredUser{
		Name:         data.Name,
		Surname:      data.Surname,
		Login:        data.Login,
		PasswordHash: data.Password,
		Email:        data.Email,
	}
	userUUID, err := s.Creator.Create(user)
	return userUUID, err
}
