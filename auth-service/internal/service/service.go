package service

import (
	"Booksiary/auth-service/internal/domain"
	"Booksiary/auth-service/internal/repository"
)

const passwordSalt = "gnerlrtnb98m7nzzwd'9"
const signingKey = "dfkslfksdfsfsdfs35bccc"

type Account interface {
	GetByLogin(login string) (domain.UserAccountResponse, error)
	Save(account UserAccountRequest) error
}

type Service struct {
	Account
	Token
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Account: NewAccountService(repos.Account),
		Token:   NewTokenService(repos.Account),
	}
}
