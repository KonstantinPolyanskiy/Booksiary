package service

import (
	"Booksiary/authorization-service/internal/types"
	"sync"
)

type UserWithCode struct {
	types.SavingUser
}
type Checker interface {
	Check()
}
type Saver interface {
	// Save сохранение пользователя на уровне бизнес-логики
	Save(user types.SavingUser) (int, error)
}
type Creator interface {
	// UserCode создает код подтверждения почты
	UserCode(user types.UserCreateResponse) error
	// FindUserByCode проверяет, валиден ли код, и если да - возвоащает пользователя для сохранения
	FindUserByCode(code types.Code) (types.SavingUser, error)
}

type Service struct {
	Creator
	Saver
}

func NewService() *Service {
	var mu sync.Mutex
	userMap := make(map[types.Code]types.SavingUser)

	return &Service{
		Creator: NewCreatorService(SavingUserMap{
			mu:      &mu,
			UserMap: userMap,
		}),
		Saver: NewSaver(),
	}
}
