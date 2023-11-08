package repository

import "github.com/dgraph-io/badger/v4"

type ConfirmationRepository struct {
	db *badger.DB
}

func NewConfirmationRepository(db *badger.DB) ConfirmationRepository {
	return ConfirmationRepository{
		db: db,
	}
}
