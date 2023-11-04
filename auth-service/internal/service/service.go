package service

import (
	"Booksiary/auth-service/internal/repository"
)

type Service struct {
	*AuthService
}

func NewService(repos *repository.Repository) *Service {
	return &Service{AuthService: NewAuthService(repos)}
}
