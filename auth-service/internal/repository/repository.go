package repository

import (
	"Booksiary/auth-service/internal/domain"
	"github.com/jmoiron/sqlx"
)

type Account interface {
	Get(login string) (domain.UserAccountResponse, error)
	Save(account domain.UserAccountDB) error
}

type Repository struct {
	Account
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Account: NewAccountRepository(db),
	}
}
