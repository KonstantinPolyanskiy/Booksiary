package service

import "Booksiary/authorization-service/internal/types"

type Creator interface {
	// UserCode создает код подтверждения почты
	UserCode(user types.User) error
}
type Service struct {
	Creator
}

func NewService() *Service {
	return &Service{Creator: NewCreatorService()}
}
