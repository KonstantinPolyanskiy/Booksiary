package service

import "log"

type Updater interface {
	Update() // заменить struct на обновленного пользователя
}

type UpdateService struct {
}

func NewUpdateService() *UpdateService {
	return &UpdateService{}
}

func (s *UpdateService) Update() {
	log.Print("Пользователь обновлен")
	return
}
