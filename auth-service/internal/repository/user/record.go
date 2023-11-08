package user

import (
	"Booksiary/auth-service/internal/domain"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"log"
	"time"
)

type RecordRepository struct {
	db *sqlx.DB
}

func NewRecordRepository(db *sqlx.DB) *RecordRepository {
	return &RecordRepository{db: db}
}

func (r *RecordRepository) Record(user domain.RegisteredUser) (uuid.UUID, error) {
	var creditionalsId int
	var uuidCreatedUser uuid.UUID

	userUUID, err := uuid.NewUUID()
	if err != nil {
		log.Printf("Ошибка в генерации UUID пользователя - %v", err)
		return uuid.UUID{}, err
	}

	addCreditionalsQuery := `
	INSERT INTO credentials
	(login, password_hash, created_at)
	VALUES 
	($1, $2, $3)
	RETURNING id
`
	addUserEntity := `
	INSERT INTO user_auth_service
	(uuid, credentials_id) 
	VALUES
	($1, $2)
	RETURNING uuid
`

	tx, err := r.db.Begin()
	if err != nil {
		log.Printf("Ошибка в старте транзакции - %v", err)
		return uuid.UUID{}, err
	}

	err = tx.QueryRow(addCreditionalsQuery, user.Login, user.PasswordHash, time.Now()).Scan(&creditionalsId)
	if err != nil {
		return uuid.UUID{}, tx.Rollback()
	}

	err = tx.QueryRow(addUserEntity, userUUID, creditionalsId).Scan(&uuidCreatedUser)
	if err != nil {
		return uuid.UUID{}, tx.Rollback()
	}

	log.Print("Успешно записано")
	return uuidCreatedUser, tx.Commit()
}
