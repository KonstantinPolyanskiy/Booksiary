package service

import (
	"Booksiary/auth-service/internal/domain"
)

// Creator создает пользователя в системе, сохраняя авторизационные данные
// и передавая их сервису Users
type Creator interface {
	Create(user domain.User) (int, error)
}

type AuthService struct {
}

func NewAuthService() *AuthService {
	return &AuthService{}
}

func (s *AuthService) Create(user domain.User) (int, error) {
	return 1, nil
}
