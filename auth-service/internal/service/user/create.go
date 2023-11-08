package user

import (
	"Booksiary/auth-service/internal/domain"
	"Booksiary/auth-service/internal/repository"
	"Booksiary/auth-service/internal/repository/user"
	"errors"
	"github.com/google/uuid"
	"log"
)

type CreateService struct {
	CreatorRepository repository.Recorder
}

func NewCreateService(repository user.RecordRepository) *CreateService {
	return &CreateService{CreatorRepository: &repository}
}
func (s *CreateService) Create(user domain.RegisteredUser) (uuid.UUID, error) {
	userUUID, err := s.CreatorRepository.Record(user)
	if err != nil {
		log.Print("ошибка при записи пользователя")
		return uuid.UUID{}, errors.New("ошибка при записи пользователя")
	}

	return userUUID, nil
}
