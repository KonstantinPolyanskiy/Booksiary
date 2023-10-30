package service

import "Booksiary/authorization-service/internal/types"

type Saver interface {
	// Save сохранение пользователя на уровне бизнес-логики
	Save(user types.SavingUser) (int, error)
}
type Creator interface {
	// UserCode создает код подтверждения почты
	UserCode(user types.SavingUser) error
	// CheckCode проверяет, валиден ли код, и если да - возвоащает пользователя для сохранения
	CheckCode(code types.Code) (types.SavingUser, error)
}
type Service struct {
	Creator
	Saver
}

func NewService() *Service {
	return &Service{
		Creator: NewCreatorService(),
		Saver:   NewSaver(),
	}
}
