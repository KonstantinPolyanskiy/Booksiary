package service

import (
	"Booksiary/user-service/internal/domain"
	"Booksiary/user-service/internal/mail"
	"Booksiary/user-service/internal/repository"
	"Booksiary/user-service/internal/service/registration"
	. "Booksiary/user-service/pkg/code"
	"github.com/google/uuid"
	"log"
)

// Sender отвечает за отправку кода на почту
type Sender interface {
	SendCode(code int, to string) error
}

// Creator отвечает за запись пользователя в хранилище
type Creator interface {
	Record(user domain.RegisteredUser) (uuid.UUID, error)
}

// UserProvider отвечает за существование пользователя в системе
type UserProvider interface {
	LoginOrEmailExist(login, email string) error
	SaveAccount(account domain.UserAccount) error
}

// CodeProvider отвечает за работу с пользователями, ожидающими подтверждения кода по почте
type CodeProvider interface {
	Add(code int, user domain.RegisteredUser) error
	Get(code int) (domain.RegisteredUser, error)
}
type SignUpService struct {
	Creator
	Sender
	UserProvider
	CodeProvider
}

func NewSignUpService(repo *repository.Repository, client mail.Mail) *SignUpService {
	return &SignUpService{
		Creator:      registration.NewRecordService(repo.User),
		Sender:       registration.NewSenderService(client),
		UserProvider: registration.NewUserProviderService(repo.User),
		CodeProvider: registration.NewConfirmService(repo.ConfirmationCode),
	}
}

func (s *SignUpService) SignUp(data domain.UserRegistrationData) error {
	var code int

	err := s.UserProvider.LoginOrEmailExist(data.Login, data.Email)
	if err != nil {
		return err
	}

	user := domain.RegisteredUser{
		Name:         data.Name,
		Surname:      data.Surname,
		Login:        data.Login,
		PasswordHash: data.Password,
		Email:        data.Email,
	}
	code = Code()

	err = s.Sender.SendCode(code, user.Email)
	if err != nil {
		log.Printf("Ошибка в отправлении кода - %v\n", err)
		return err
	}
	log.Printf("Код отправлен на почту")

	err = s.CodeProvider.Add(code, user)
	if err != nil {
		log.Printf("Ошибка записи кода в badger - %v\n", err)
		return err
	}
	log.Printf("Код записан в badger")

	return nil
}
func (s *SignUpService) SignUpCallback(code int) (uuid.UUID, error) {
	user, err := s.CodeProvider.Get(code)
	if err != nil {
		log.Printf("Ошибка в получении пользователя по коду - %v\n", err)
		return uuid.UUID{}, err
	}

	userUUID, err := s.Creator.Record(user)
	if err != nil {
		log.Printf("Ошибка в записи пользователя в базу данных - %v\n", err)
		return uuid.UUID{}, err
	}

	account := domain.UserAccount{
		UUID:     userUUID,
		Login:    user.Login,
		Password: user.PasswordHash,
	}
	err = s.UserProvider.SaveAccount(account)

	return userUUID, nil
}
