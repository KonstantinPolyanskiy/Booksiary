package service

import "Booksiary/auth-service/internal/handler/domain"

// Creator создает пользователя в системе, сохраняя авторизационные данные
// и передавая их сервису Users
type Creator interface {
	Create(user domain.User) (int, error)
}

type AuthService struct {
	Creator
}

func NewAuthService() *AuthService {
	return &AuthService{}
}

func (s *AuthService) Create(user domain.User) (int, error) {
	return 1, nil
}
