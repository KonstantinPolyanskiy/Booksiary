package repository

import (
	"Booksiary/user-service/internal/domain"
	"errors"
	"github.com/google/uuid"
	"github.com/jackc/pgx"
	"github.com/jmoiron/sqlx"
	"log"
	"time"
)

type UserRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}
func (r *UserRepository) GetByEmail(email string) (domain.UserPersonality, error) {
	var personality domain.UserPersonality

	getUserByEmailQuery := `
	SELECT name, surname, email
	FROM person
	WHERE email=$1
`
	err := r.db.Get(&personality, getUserByEmailQuery, email)
	if errors.Is(err, pgx.ErrNoRows) {
		return domain.UserPersonality{}, nil
	}
	if err != nil {
		return domain.UserPersonality{}, nil
	}

	return personality, nil
}
func (r *UserRepository) Record(user domain.RegisteredUser) (uuid.UUID, error) {
	var userTechnicalId int
	var createdUserUUID uuid.UUID

	potentialUUID := uuid.New()

	recordUserTechnicalQuery := `
	INSERT INTO technical_data 
	(created_at, updated_at) 
	VALUES 
	($1, $2)
	RETURNING id
	`

	recordUserQuery := `
	INSERT INTO person
	(uuid, technical_data_id, name, surname, email) 
	VALUES
	($1, $2, $3, $4, $5)
	RETURNING uuid
`

	tx, err := r.db.Begin()
	if err != nil {
		return uuid.UUID{}, err
	}

	err = tx.QueryRow(recordUserTechnicalQuery, time.Now(), time.Now()).Scan(&userTechnicalId)
	if err != nil {
		log.Printf("Ошибка в записи технической информации пользователя %v\n", err)
		return uuid.UUID{}, tx.Rollback()
	}

	err = tx.QueryRow(recordUserQuery, potentialUUID, userTechnicalId, user.Name, user.Surname, user.Email).Scan(&createdUserUUID)
	if err != nil {
		log.Printf("Ошибка в записи пользователя в базу данных - %v\n", err)
		return uuid.UUID{}, tx.Rollback()
	}

	log.Print("Пользователь успешно записан в базу данных")
	return createdUserUUID, tx.Commit()
}

func (r *UserRepository) Delete(userUuid uuid.UUID) error {
	return errors.New("еще не реализованно")
}
