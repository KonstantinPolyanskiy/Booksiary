package user

import (
	"Booksiary/auth-service/internal/repository/user"
)

type ExistService struct {
	ExistRepository user.ExistRepository
}

func (s *ExistService) Exist(login string) error {
	if err := s.ExistRepository.Exist(login); err != nil {
		return err
	}
	return nil
}

func NewExistService(repository user.ExistRepository) *ExistService {
	return &ExistService{ExistRepository: repository}
}
