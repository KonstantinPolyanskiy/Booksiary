package service

import (
	"Booksiary/user-service/internal/domain"
	"Booksiary/user-service/internal/mail"
	"Booksiary/user-service/internal/repository"
	"github.com/google/uuid"
)

type Registration interface {
	SignUp(data domain.UserRegistrationData) error
	SignUpCallback(code int) (uuid.UUID, error)
}
type Service struct {
	Registration
}

func NewService(repo *repository.Repository, client mail.Mail) *Service {
	return &Service{
		Registration: NewSignUpService(repo, client),
	}
}
