package service

// Root представляет собой каталог микросервиса: /user /auth
type Root string

type Handle struct {
	Host string
	Port string
}

type Service struct {
	UserService
	ProxyAddrMap
}

func NewService(addrMap ProxyAddrMap) *Service {
	return &Service{
		UserService:  *NewUserService(addrMap),
		ProxyAddrMap: addrMap,
	}
}
