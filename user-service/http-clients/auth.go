package http_clients

import (
	"Booksiary/user-service/internal/domain"
	"errors"
	"net/http"
)

type AuthClient struct {
	httpClient *http.Client
}

func NewAuthClient() *AuthClient {
	return &AuthClient{
		httpClient: http.DefaultClient,
	}
}

func (c *AuthClient) GetAccount(login string) (domain.UserAccount, error) {
	return domain.UserAccount{}, errors.New("not implemented")
}
