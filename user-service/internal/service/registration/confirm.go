package registration

import "Booksiary/user-service/internal/repository"

type ConfirmService struct {
	repo repository.ConfirmationRepository
}

func NewConfirmService(repo repository.ConfirmationRepository) *ConfirmService {
	return &ConfirmService{
		repo: repo,
	}
}
