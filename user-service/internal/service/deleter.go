package service

import "log"

type Deleter interface {
	Delete()
}

type DeleteService struct {
}

func NewDeleteService() *DeleteService {
	return &DeleteService{}
}

func (s *DeleteService) Delete() {
	log.Print("Пользователь удален")
}
