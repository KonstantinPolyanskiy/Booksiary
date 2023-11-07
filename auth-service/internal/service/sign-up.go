package service

import (
	"Booksiary/auth-service/internal/domain"
	"Booksiary/auth-service/internal/repository"
	"Booksiary/auth-service/pkg/mail"
	"Booksiary/auth-service/pkg/password"
	"github.com/google/uuid"
	"log"
)

var htmlBody = `<html>
	<head>
		<meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
		<title>Hello Gophers!</title>
	</head>
	<body>
		<p>Тест</p>
	</body>
</html>`

const passwordSalt = "testsalt-dsgfkhergtfdrfwm"

// Creator создает пользователя в системе, сохраняя авторизационные данные
// и передавая их сервису Users
type Creator interface {
	Create(user domain.UserRegistrationData) (int, error)
}

type AuthService struct {
	repos      *repository.Repository
	mailClient *mail.Client
}

func NewAuthService(repos *repository.Repository, mailClient *mail.Client) *AuthService {
	return &AuthService{
		repos:      repos,
		mailClient: mailClient,
	}
}

func (s *AuthService) Create(userData domain.UserRegistrationData) (uuid.UUID, error) {
	//var wg sync.WaitGroup
	var userUUID uuid.UUID
	var err error

	log.Printf("Сообщение отправлено")

	registeredUser := domain.RegisteredUser{
		Name:         userData.Name,
		Surname:      userData.Surname,
		Login:        userData.Login,
		PasswordHash: password.Hash(userData.Password, passwordSalt),
		Email:        userData.Email,
	}

	err = s.mailClient.SendCode(1234, registeredUser.Email)
	if err != nil {
		log.Print(err)
		return uuid.UUID{}, err
	}
	log.Printf("Код отправлен на почту - %s", registeredUser.Email)

	/*wg.Add(1)
	go func() {
		defer wg.Done()
		userUUID, err = s.repos.AuthRepo.Record(registeredUser)
	}()
	wg.Wait()

	//client := http_client.Default()

	r, err := http.DefaultClient.Post("http://localhost:8080/api/user/user-data", "application/json", bytes.NewReader([]byte(registeredUser.Name)))
	if err != nil {
		log.Print(err)
	}
	log.Printf("Ответ - %s", r.Request.Host)*/
	//client.Post("localhost:8080/api/user-data", "application/json", bytes.NewReader([]byte(registeredUser.Name)))

	if err != nil {
		return uuid.UUID{}, err
	}
	return userUUID, err
}
