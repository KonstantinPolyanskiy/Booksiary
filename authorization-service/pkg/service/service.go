package service

import "Booksiary/authorization-service/internal/types"

type Creator interface {
	User(user types.User) (int, error)
}
type Service struct {
	Creator
}

func NewService() *Service {
	return &Service{Creator: NewCreatorService()}
}
