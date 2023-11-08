package registration

import (
	http_clients "Booksiary/user-service/http-clients"
	"Booksiary/user-service/internal/domain"
	"errors"
)

type AuthClient interface {
	GetAccount(login string) (domain.UserAccount, error)
}

var existError = errors.New("login or email is already exist")

type UserProviderService struct {
	AuthClient
}

func NewUserProviderService() *UserProviderService {
	return &UserProviderService{
		AuthClient: http_clients.NewAuthClient(),
	}
}

func (s *UserProviderService) LoginOrEmailExist(login, email string) error {
	//TODO: реализовать проверку существования логина и почты
	receivedAccount, err := s.GetAccount(login)
	if err != nil {
		return err
	}

	return existError
}
