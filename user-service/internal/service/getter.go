package service

import "log"

type Getter interface {
	Get()
}

type GetterService struct {
}

func NewGetterService() *GetterService {
	return &GetterService{}
}

func (s *GetterService) Get() {
	log.Print("пользователь отдан")
}
