package registration

import (
	"Booksiary/user-service/internal/domain"
	"Booksiary/user-service/internal/http-clients"
	"Booksiary/user-service/internal/repository"
	"errors"
	"sync"
)

type AuthClient interface {
	GetAccount(login string) (domain.UserAccount, error)
	SaveAccount(account domain.UserAccount) error
}

var existError = errors.New("login or email is already exist")

type UserProviderService struct {
	repo repository.User
	AuthClient
}

func NewUserProviderService(repo repository.User) *UserProviderService {
	return &UserProviderService{
		repo:       repo,
		AuthClient: http_clients.NewAuthClient(),
	}
}

func (s *UserProviderService) LoginOrEmailExist(login, email string) error {
	var wg sync.WaitGroup
	var errExist error

	wg.Add(2)

	go func() {
		defer wg.Done()

		account, err := s.AuthClient.GetAccount(login)
		if err != nil {
			errExist = err
		}
		if account.Login == login {
			errExist = existError
		}
	}()

	go func() {
		defer wg.Done()

		personality, err := s.repo.GetByEmail(email)
		if err != nil {
			errExist = err
		}
		if personality.Email == email {
			errExist = existError
		}
	}()

	wg.Wait()

	return errExist
}

func (s *UserProviderService) SaveAccount(account domain.UserAccount) error {
	err := s.AuthClient.SaveAccount(account)
	if err != nil {
		return err
	}
	return nil
}
