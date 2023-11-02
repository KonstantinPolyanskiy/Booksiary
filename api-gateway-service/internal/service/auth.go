package service

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"net/http/httputil"
)

type AuthService struct {
	ProxyAddrMap
}

func NewAuthService(addrMap ProxyAddrMap) *AuthService {
	return &AuthService{
		ProxyAddrMap: addrMap,
	}
}

func (s *AuthService) Redirect(w http.ResponseWriter, r *http.Request, root Root) {
	director := NewDirector(root, s.M)

	proxy := &httputil.ReverseProxy{Director: director}

	proxy.ModifyResponse = func(response *http.Response) error {
		cont, err := io.ReadAll(response.Body)
		if err != nil {
			log.Print("Ошибка чтения ответа - ", err)
		}
		log.Print("Тело ответа - ", cont)
		response.Body = io.NopCloser(bytes.NewReader(cont))

		return nil
	}

	proxy.ServeHTTP(w, r)
}
