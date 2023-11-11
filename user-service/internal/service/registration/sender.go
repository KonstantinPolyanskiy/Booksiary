package registration

import "Booksiary/user-service/internal/mail"

type SenderService struct {
	client mail.Mail
}

func NewSenderService(client mail.Mail) *SenderService {
	return &SenderService{
		client: client,
	}
}

func (s *SenderService) SendCode(code int, to string) error {
	err := s.client.SendCode(code, to)
	if err != nil {
		return err
	}

	return nil
}
