package repository

import "github.com/jmoiron/sqlx"

type Repository struct {
	AuthRepo
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		AuthRepo: *NewAuthRepo(db),
	}
}
