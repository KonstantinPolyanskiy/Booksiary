package mail

import (
	mail "github.com/xhit/go-simple-mail/v2"
	"strconv"
	"time"
)

type Sender interface {
	SendCode(code int) error
}

type Client struct {
	server  *mail.SMTPServer
	client  *mail.SMTPClient
	address string
}

func NewEmailClient(port int, host, username, password, emailAddress string) (*Client, error) {
	server := mail.NewSMTPClient()

	server.Host = host
	server.Port = port
	server.Username = username
	server.Password = password
	server.Encryption = mail.EncryptionSSLTLS
	server.ConnectTimeout = 100 * time.Second
	server.SendTimeout = 20 * time.Second

	smtpClient, err := server.Connect()

	if err != nil {
		return nil, err
	}
	return &Client{
		server:  server,
		client:  smtpClient,
		address: emailAddress,
	}, nil
}

func (c *Client) SendCode(code int, to string) error {
	email := mail.NewMSG()
	email.SetFrom("work.polyanskiy@mail.ru").AddTo(to).SetSubject("Код регистрации")
	email.GetFrom()
	email.SetBody(mail.TextPlain, strconv.Itoa(code))

	return email.Send(c.client)
}
