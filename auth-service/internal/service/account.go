package service

import (
	"Booksiary/auth-service/internal/domain"
	"Booksiary/auth-service/internal/repository"
	"Booksiary/auth-service/pkg/password"
	"github.com/google/uuid"
)

type UserAccountRequest struct {
	UUID     uuid.UUID `json:"UUID"`
	Login    string    `json:"Login"`
	Password string    `json:"Password"`
}

type AccountService struct {
	repository.Account
}

func NewAccountService(repo repository.Account) *AccountService {
	return &AccountService{
		Account: repo,
	}
}

func (s *AccountService) Save(account UserAccountRequest) error {
	accountDB := domain.UserAccountDB{
		UUID:         account.UUID,
		Login:        account.Login,
		PasswordHash: password.Hash(account.Password, "gnerlrtnb98m7nzzwd'9"),
	}
	err := s.Account.Save(accountDB)
	if err != nil {
		return err
	}

	return nil
}
func (s *AccountService) Get(login string) (domain.UserAccountResponse, error) {
	account, err := s.Account.Get(login)
	if err != nil {
		return domain.UserAccountResponse{}, err
	}

	return account, nil
}
