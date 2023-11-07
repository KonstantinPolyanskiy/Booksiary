package service

import (
	"Booksiary/auth-service/internal/domain"
	"Booksiary/auth-service/internal/repository"
	"Booksiary/auth-service/pkg/password"
	"bytes"
	"github.com/google/uuid"
	mail "github.com/xhit/go-simple-mail/v2"
	"log"
	"net/http"
	"sync"
	"time"
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
	repos *repository.Repository
}

func NewAuthService(repos *repository.Repository) *AuthService {
	return &AuthService{repos: repos}
}

func (s *AuthService) Create(userData domain.UserRegistrationData) (uuid.UUID, error) {
	var wg sync.WaitGroup
	var userUUID uuid.UUID
	var err error

	client := mail.NewSMTPClient()

	client.Host = "smtp.mail.ru"
	client.Port = 465
	client.Username = "email"
	client.Password = "passoword"
	client.Encryption = mail.EncryptionSSLTLS
	client.ConnectTimeout = 100 * time.Second
	client.SendTimeout = 10 * time.Second

	smtpClient, err := client.Connect()
	if err != nil {
		log.Fatalf("Ошибка в создание SmtpClient - %v", err)
	}

	err = sendEmail(htmlBody, "to_email", smtpClient)
	if err != nil {
		log.Print(err)
		return uuid.UUID{}, err
	}
	log.Printf("Сообщение отправлено")

	registeredUser := domain.RegisteredUser{
		Name:         userData.Name,
		Surname:      userData.Surname,
		Login:        userData.Login,
		PasswordHash: password.Hash(userData.Password, passwordSalt),
		Email:        userData.Email,
	}

	wg.Add(1)
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
	log.Printf("Ответ - %s", r.Request.Host)
	//client.Post("localhost:8080/api/user-data", "application/json", bytes.NewReader([]byte(registeredUser.Name)))

	if err != nil {
		return uuid.UUID{}, err
	}
	return userUUID, err
}

func sendEmail(body, to string, client *mail.SMTPClient) error {
	email := mail.NewMSG()

	email.SetFrom("work.polyanskiy@mail.ru").AddTo(to).SetSubject("Тестовое сообщение")

	email.GetFrom()
	email.SetBody(mail.TextHTML, body)

	if email.Error != nil {
		return email.Error
	}

	return email.Send(client)
}
