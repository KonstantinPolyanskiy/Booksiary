package repository

import (
	"Booksiary/auth-service/internal/domain"
	"errors"
	"github.com/jackc/pgx"
	"github.com/jmoiron/sqlx"
	"log"
	"time"
)

type AccountRepository struct {
	db *sqlx.DB
}

func NewAccountRepository(db *sqlx.DB) *AccountRepository {
	return &AccountRepository{
		db: db,
	}
}

func (r *AccountRepository) Get(login string) (domain.UserAccountResponse, error) {
	var account domain.UserAccountResponse

	getUserByLoginQuery := `
	SELECT (uuid, login, passwordhash)
	FROM account
	WHERE login=$1
`

	err := r.db.Get(&account, getUserByLoginQuery, login)
	if errors.Is(err, pgx.ErrNoRows) {
		return domain.UserAccountResponse{}, nil
	}
	if err != nil {
		return domain.UserAccountResponse{}, err
	}

	return account, nil
}

func (r *AccountRepository) Save(account domain.UserAccountDB) error {
	var technicalId int
	addTechnicalQuery := `
	INSERT INTO technical_data
	(created_at, update_at)
	VALUES 
	($1, $2)
	RETURNING id
`
	addAccountQuery := `
	INSERT INTO account
	(uuid, technical_data_id, login, passwordhash)
	VALUES 
	($1, $2, $3, $4)
`
	tx, err := r.db.Begin()
	if err != nil {
		log.Printf("Ошибка в начале транзакции - %v\n", err)
		return tx.Rollback()
	}

	err = tx.QueryRow(addTechnicalQuery, time.Now(), time.Now()).Scan(&technicalId)
	if err != nil {
		log.Printf("Ошибка в записи технических данных в бд - %v\n", err)
		return tx.Rollback()
	}

	_, err = tx.Exec(addAccountQuery, account.UUID, technicalId, account.Login, account.PasswordHash)
	if err != nil {
		log.Printf("Ошибка в записи аккаунта в бд - %v\n", err)
		return tx.Rollback()
	}

	return tx.Commit()
}
