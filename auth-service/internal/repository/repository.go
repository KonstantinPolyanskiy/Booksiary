package repository

import (
	"Booksiary/auth-service/internal/repository/user"
	"github.com/jmoiron/sqlx"
)

type Repository struct {
	RecordRepository user.RecordRepository
	ExistRepository  user.ExistRepository
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		RecordRepository: *user.NewRecordRepository(db),
		ExistRepository:  *user.NewExistRepository(db),
	}
}
