package service

import "Booksiary/authorization-service/internal/types"

type Saver interface {
	// Save сохранение пользователя на уровне бизнес-логики
	Save(user types.SaveUser) (int, error)
}
type Creator interface {
	// UserCode создает код подтверждения почты
	UserCode(user types.SaveUser) error
	// CheckCode проверяет, валиден ли код, и если да - возвоащает пользователя для сохранения
	CheckCode(code types.Code) (types.SaveUser, error)
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
