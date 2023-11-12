package http_clients

import (
	"Booksiary/user-service/internal/domain"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type AuthClient struct {
	URL        string
	httpClient *http.Client
}

func NewAuthClient() *AuthClient {
	return &AuthClient{
		URL:        "http://localhost:8081/interaction/user/",
		httpClient: http.DefaultClient,
	}
}

func (c *AuthClient) GetAccount(login string) (domain.UserAccount, error) {
	var account domain.UserAccount
	resp, err := c.httpClient.Get(fmt.Sprintf("%s%s", c.URL, login))
	if err != nil {
		return domain.UserAccount{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return domain.UserAccount{}, err
	}

	if err := json.Unmarshal(body, &account); err != nil {
		return domain.UserAccount{}, err
	}

	return account, nil
}

func (c *AuthClient) SaveAccount(account domain.UserAccount) error {
	path := fmt.Sprintf("%s%s", c.URL, "save-account")
	bodyRequest, err := json.Marshal(account)
	if err != nil {
		return err
	}

	resp, err := c.httpClient.Post(path, "application/json", bytes.NewReader(bodyRequest))
	if err != nil {
		log.Printf("Ошибка в отправке аккаунта в auth - %v\n", err)
		return err
	}
	defer resp.Body.Close()

	return nil
}
