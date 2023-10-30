package service

import (
	"Booksiary/authorization-service/internal/types"
	"log"
)

type SaverService struct {
}

func (s *SaverService) Save(user types.SavingUser) (int, error) {
	log.Printf("Пользователь %v сохранен", user)
	return 1, nil
}

func NewSaver() *SaverService {
	return &SaverService{}
}
