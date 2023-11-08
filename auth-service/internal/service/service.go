package service

import (
	"Booksiary/auth-service/internal/repository"
	"Booksiary/auth-service/pkg/mail"
)

type Service struct {
	*AuthService
	*mail.Client
}

func NewService(repos *repository.Repository, emailClient *mail.Client) *Service {
	return &Service{
		AuthService: NewAuthService(repos, emailClient),
	}
}
