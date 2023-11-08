package mail

import (
	"fmt"
	mail "github.com/xhit/go-simple-mail/v2"
	"strconv"
	"time"
)

type Config struct {
	Host           string
	Username       string
	Password       string
	EmailAddress   string
	Port           int
	Encryption     mail.Encryption
	ConnectTimeout time.Duration
	SendTimeout    time.Duration
}

type Mail struct {
	server       *mail.SMTPServer
	emailAddress string
}

func NewMailClient(cfg Config) (*Mail, error) {
	server := mail.NewSMTPClient()

	server.Host = cfg.Host
	server.Username = cfg.Username
	server.Password = cfg.Password
	server.Port = cfg.Port
	server.Encryption = cfg.Encryption
	server.ConnectTimeout = cfg.ConnectTimeout
	server.SendTimeout = cfg.SendTimeout

	test, err := server.Connect()
	if err != nil {
		return nil, err
	}
	if err = test.Close(); err != nil {
		return nil, err
	}

	return &Mail{
		server:       server,
		emailAddress: cfg.EmailAddress,
	}, nil
}
func (m *Mail) SendCode(code int, to string) error {
	client, err := m.server.Connect()
	if err != nil {
		return err
	}
	defer client.Close()

	email := mail.NewMSG()
	if err != nil {
		return err
	}

	email.
		SetFrom(m.emailAddress).
		AddTo(to).
		SetSubject(m.emailAddress).
		SetBody(mail.TextPlain, fmt.Sprintf("Код регистрации - %s", strconv.Itoa(code))).
		GetFrom()

	return email.Send(client)

}
