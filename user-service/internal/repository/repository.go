package repository

import (
	"github.com/dgraph-io/badger/v4"
	"github.com/jmoiron/sqlx"
)

type Repository struct {
	ConfirmationRepository
}

func NewRepository(db *sqlx.DB, memoryDb *badger.DB) *Repository {
	return &Repository{
		ConfirmationRepository: NewConfirmationRepository(memoryDb),
	}
}
