package service

import "net/http"

// Root представляет собой каталог микросервиса: /user /auth
type Root string

type Handle struct {
	Host string
	Port string
}

type Service struct {
	UserService
	AuthService
	ProxyAddrMap
}

func NewService(addrMap ProxyAddrMap) *Service {
	return &Service{
		UserService:  *NewUserService(addrMap),
		AuthService:  *NewAuthService(addrMap),
		ProxyAddrMap: addrMap,
	}
}
