package registration

import (
	"Booksiary/user-service/internal/domain"
	"Booksiary/user-service/internal/repository"
	"github.com/google/uuid"
)

type RecordService struct {
	repo repository.User
}

func NewRecordService(repo repository.User) *RecordService {
	return &RecordService{
		repo: repo,
	}
}

// Record создает пользователя в базе данных и возвращает UUID созданного пользователя
func (s *RecordService) Record(user domain.RegisteredUser) (uuid.UUID, error) {
	userUUID, err := s.repo.Record(user)
	if err != nil {
		return uuid.UUID{}, err
	}

	return userUUID, nil
}
