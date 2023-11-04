package service

import (
	"Booksiary/auth-service/internal/domain"
	"Booksiary/auth-service/internal/repository"
	"Booksiary/auth-service/pkg/password"
	"github.com/google/uuid"
)

const passwordSalt = "testsalt-dsgfkhergtfdrfwm"

// Creator создает пользователя в системе, сохраняя авторизационные данные
// и передавая их сервису Users
type Creator interface {
	Create(user domain.UserRegistrationData) (int, error)
}

type AuthService struct {
	repos *repository.Repository
}

func NewAuthService(repos *repository.Repository) *AuthService {
	return &AuthService{repos: repos}
}

func (s *AuthService) Create(userData domain.UserRegistrationData) (uuid.UUID, error) {
	registeredUser := domain.RegisteredUser{
		Name:         userData.Name,
		Surname:      userData.Surname,
		Login:        userData.Login,
		PasswordHash: password.Hash(userData.Password, passwordSalt),
		Email:        userData.Email,
	}

	userUUID, err := s.repos.AuthRepo.Record(registeredUser)
	if err != nil {
		return uuid.UUID{}, err
	}
	return userUUID, err
}
