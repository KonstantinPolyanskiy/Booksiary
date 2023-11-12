package service

import (
	"Booksiary/auth-service/internal/domain"
	"Booksiary/auth-service/internal/repository"
)

type Account interface {
	Get(login string) (domain.UserAccountResponse, error)
	Save(account UserAccountRequest) error
}

type Service struct {
	Account
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Account: NewAccountService(repos.Account),
	}
}
