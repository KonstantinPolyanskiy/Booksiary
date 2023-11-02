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

// NewDirector возвращает
func NewDirector(root Root, M map[Root]Handle) func(request *http.Request) {
	director := func(request *http.Request) {
		h := M[root] // Получение хоста/порта сервиса

		request.URL.Scheme = "http"
		request.URL.Host = h.Host + ":" + h.Port
	}
	return director
}
