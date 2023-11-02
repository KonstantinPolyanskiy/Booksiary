package service

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"net/http/httputil"
)

type Redirector interface {
	Redirect(w http.ResponseWriter, r *http.Request)
}

type UserService struct {
	ProxyAddrMap
}

func NewUserService(addrMap ProxyAddrMap) *UserService {
	return &UserService{ProxyAddrMap: addrMap}
}

func (s *UserService) Redirect(w http.ResponseWriter, r *http.Request) {
	director := func(request *http.Request) {
		h := s.M["/user"] // Получение хоста/порта сервиса

		request.URL.Scheme = "http"
		request.URL.Host = h.Host + ":" + h.Port
	}

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
