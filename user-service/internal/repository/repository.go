package repository

import (
	"github.com/dgraph-io/badger/v4"
	"github.com/jmoiron/sqlx"
)

type ConfirmationCode interface {
	Add(code int, user domain.RegisteredUser) error
	Get(code int) (domain.RegisteredUser, error)
}

type User interface {
	Record(user domain.RegisteredUser) (uuid.UUID, error)
	Delete(userUuid uuid.UUID) error
}

type Repository struct {
	ConfirmationCode
	User
}

func NewRepository(db *sqlx.DB, memoryDb *badger.DB) *Repository {
	return &Repository{
		ConfirmationCode: NewConfirmationRepository(memoryDb),
		User:             NewUserRepository(db),
	}
}
