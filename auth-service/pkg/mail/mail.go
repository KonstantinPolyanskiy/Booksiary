package mail

import (
	mail "github.com/xhit/go-simple-mail/v2"
	"strconv"
	"time"
)

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

	return &Client{
		server:  server,
		client:  nil,
		address: emailAddress,
	}, nil
}

func (c *Client) SendCode(code int, to string) error {
	var err error

	c.client, err = c.server.Connect()
	if err != nil {
		return err
	}
	email := mail.NewMSG()
	email.SetFrom("work.polyanskiy@mail.ru").AddTo(to).SetSubject("Код регистрации")
	email.GetFrom()
	email.SetBody(mail.TextPlain, strconv.Itoa(code))

	return email.Send(c.client)
}
