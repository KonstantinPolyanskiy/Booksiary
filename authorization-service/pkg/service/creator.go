package service

import (
	"Booksiary/authorization-service/internal/types"
	"errors"
	"math/rand"
	"time"
)

type CreatorService struct {
}

func NewCreatorService() *CreatorService {
	return &CreatorService{}
}

func (cs CreatorService) User(user types.User) (int, error) {
	rand.Seed(time.Now().UnixNano())
	if user.Email == "" {
		return 0, errors.New("пустая почта")
	}
	return rand.Intn(100), nil
}
