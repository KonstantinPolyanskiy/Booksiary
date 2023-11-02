package service

import "log"

type Creator interface {
	Create() (int, error)
}
type CreateService struct {
}

func NewCreator() *CreateService {
	return &CreateService{}
}

func (c *CreateService) Create() (int, error) {
	log.Print("Пользователь создан")
	return 123, nil
}
