package mail

import (
	"errors"
	"log"
)

type Mail interface {
	Send() error
}

type Email struct {
	Addr string
	Msg  string
}

func NewEmail(addr, msg string) Email {
	return Email{
		Addr: addr,
		Msg:  msg,
	}
}

func (m Email) Send() error {
	if m.Addr == "" && m.Msg == "" {
		return errors.New("empty addr/msg")
	}
	log.Print("Отправлено - ", m.Msg, m.Addr)
	return nil
}
