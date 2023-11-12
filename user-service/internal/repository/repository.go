package repository

import (
	"Booksiary/user-service/internal/domain"
	"github.com/dgraph-io/badger/v4"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type ConfirmationCode interface {
	Add(code int, user domain.RegisteredUser) error
	Get(code int) (domain.RegisteredUser, error)
}

type User interface {
	Record(user domain.RegisteredUser) (uuid.UUID, error)
	Delete(userUuid uuid.UUID) error
	GetByEmail(email string) (domain.UserPersonality, error)
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
