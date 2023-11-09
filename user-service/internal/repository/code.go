package repository

import (
	"Booksiary/user-service/internal/domain"
	"encoding/json"
	"github.com/dgraph-io/badger/v4"
	"time"
)

type ConfirmationRepository struct {
	db *badger.DB
}

func NewConfirmationRepository(db *badger.DB) ConfirmationRepository {
	return ConfirmationRepository{
		db: db,
	}
}

func (r *ConfirmationRepository) Add(code int, user domain.RegisteredUser) error {
	codeKey, err := json.Marshal(code)
	if err != nil {
		return err
	}

	userValue, err := json.Marshal(user)
	if err != nil {
		return err
	}

	err = r.db.Update(func(txn *badger.Txn) error {
		e := badger.NewEntry(codeKey, userValue).WithTTL(2 * time.Minute)
		err := txn.SetEntry(e)
		return err
	})
	if err != nil {
		return err
	}

	return nil
}

func (r *ConfirmationRepository) Get(code int) (domain.RegisteredUser, error) {
	var userCopy []byte
	var user domain.RegisteredUser
	codeKey, err := json.Marshal(code)
	if err != nil {
		return domain.RegisteredUser{}, err
	}

	err = r.db.View(func(txn *badger.Txn) error {
		item, err := txn.Get(codeKey)
		if err != nil {
			return err
		}
		userCopy, err = item.ValueCopy(userCopy)
		if err != nil {
			return err
		}
		return nil
	})

	err = json.Unmarshal(userCopy, &user)
	if err != nil {
		return domain.RegisteredUser{}, err
	}

	return user, nil
}
