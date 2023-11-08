package user

import (
	"errors"
	"github.com/jmoiron/sqlx"
)

type ExistRepository struct {
	db *sqlx.DB
}

func NewExistRepository(db *sqlx.DB) *ExistRepository {
	return &ExistRepository{db: db}
}

func (r *ExistRepository) Exist(login string) error {
	var loginExist bool

	loginExistQuery := `
	SELECT EXISTS (SELECT login FROM credentials WHERE login=$1)
`
	err := r.db.QueryRow(loginExistQuery, login).Scan(&loginExist)
	if err != nil {
		return err
	}
	if loginExist {
		return errors.New("логин занят")
	}
	return nil
}
