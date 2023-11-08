package repository

import (
	"Booksiary/auth-service/internal/domain"
	"github.com/google/uuid"
)

type UserProvider interface {
	Exist(login string) error
}
type Recorder interface {
	Record(user domain.RegisteredUser) (uuid.UUID, error)
}

type AuthRepository struct {
	UserProvider
	Recorder
}

func (r *Repository) Save(user domain.RegisteredUser) (uuid.UUID, error) {
	var err error
	var uuidUser uuid.UUID

	err = r.ExistRepository.Exist(user.Login)
	if err != nil {
		return uuid.UUID{}, err
	}

	uuidUser, err = r.RecordRepository.Record(user)
	if err != nil {
		return uuid.UUID{}, err
	}

	return uuidUser, nil
}
