package registration

import (
	"Booksiary/user-service/internal/domain"
	"github.com/google/uuid"
)

type CreateService struct {
}

func NewCreateService() *CreateService {
	return &CreateService{}
}

// Create создает пользователя в базе данных и возвращает UUID созданного пользователя
func (s *CreateService) Create(user domain.RegisteredUser) (uuid.UUID, error) {
	return uuid.UUID{}, nil
}
