package service

import (
	"Booksiary/auth-service/internal/domain"
	"Booksiary/auth-service/internal/repository"
	"Booksiary/auth-service/internal/service/user"
	"Booksiary/auth-service/pkg/code"
	"Booksiary/auth-service/pkg/mail"
	"Booksiary/auth-service/pkg/password"
	"github.com/google/uuid"
	"log"
	"sync"
)

const passwordSalt = "testsalt-dsgfkhergtfdrfwm"

type CodeStore struct {
	mu      *sync.Mutex
	CodeMap map[int]domain.RegisteredUser
}

type UserProvider interface {
	Exist(login string) error
}

type Creator interface {
	Create(user domain.RegisteredUser) (uuid.UUID, error)
}
type Sender interface {
	SendCode(code int, to string) error
}
type AuthService struct {
	Sender
	Creator
	UserProvider
	repos     *repository.Repository
	codeStore CodeStore
}

func NewAuthService(repos *repository.Repository, mailClient *mail.Client) *AuthService {
	return &AuthService{
		repos:        repos,
		Sender:       mailClient,
		Creator:      user.NewCreateService(repos.RecordRepository),
		UserProvider: user.NewExistService(repos.ExistRepository),
	}
}

func (s *AuthService) Create(userData domain.UserRegistrationData) (uuid.UUID, error) {
	var wg sync.WaitGroup
	var userUUID uuid.UUID
	var err error

	registeredUser := domain.RegisteredUser{
		Name:         userData.Name,
		Surname:      userData.Surname,
		Login:        userData.Login,
		PasswordHash: password.Hash(userData.Password, passwordSalt),
		Email:        userData.Email,
	}
	err = s.UserProvider.Exist(registeredUser.Login)
	if err != nil {
		log.Print(err)
		return uuid.UUID{}, err
	}

	err = s.Sender.SendCode(code.Code(), registeredUser.Email)
	if err != nil {
		log.Print(err)
		return uuid.UUID{}, err
	}
	log.Printf("Код отправлен на почту - %s", registeredUser.Email)

	wg.Add(1)
	go func() {
		defer wg.Done()
		userUUID, err = s.Creator.Create(registeredUser)
	}()
	wg.Wait()

	//client := http_client.Default()

	//r, err := http.DefaultClient.Post("http://localhost:8080/api/user/user-data", "application/json", bytes.NewReader([]byte(registeredUser.Name)))
	if err != nil {
		log.Print(err)
	}
	//log.Printf("Ответ - %s", r.Request.Host)*/
	//client.Post("localhost:8080/api/user-data", "application/json", bytes.NewReader([]byte(registeredUser.Name)))

	if err != nil {
		return uuid.UUID{}, err
	}
	return userUUID, err
}
