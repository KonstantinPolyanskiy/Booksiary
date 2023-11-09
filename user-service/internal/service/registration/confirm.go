package registration

import (
	"Booksiary/user-service/internal/domain"
	"Booksiary/user-service/internal/repository"
)

type ConfirmService struct {
	repo repository.ConfirmationRepository
}

func NewConfirmService(repo repository.ConfirmationRepository) *ConfirmService {
	return &ConfirmService{
		repo: repo,
	}
}

func (s *ConfirmService) Add(code int, user domain.RegisteredUser) error {
	err := s.repo.Add(code, user)
	if err != nil {
		return err
	}
	return nil
}

func (s *ConfirmService) Get(code int) (domain.RegisteredUser, error) {
	user, err := s.repo.Get(code)
	if err != nil {
		return domain.RegisteredUser{}, err
	}

	return user, nil
}
