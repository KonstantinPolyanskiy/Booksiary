package proxy

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

type Handle struct {
	Host string
	Port string
}

type Service struct {
	Auth *Handle
	User *Handle
}

func (s *Service) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var target *url.URL
	if strings.Contains(r.RequestURI, "/api/user") {
		target, _ = url.Parse("http://" + s.User.Host + ":" + s.User.Port)
	} else if strings.Contains(r.RequestURI, "/api/autoCommentList") {
		target, _ = url.Parse("http://" + s.Auth.Host + ":" + s.Auth.Port)
	} else {
		fmt.Fprintf(w, "404 Not Found")
		return
	}

	tQuery := target.RawQuery

	director := func(r *http.Request) {
		r.URL.Scheme = target.Scheme
		r.URL.Host = target.Host
		r.URL.Path = target.Path

		if tQuery == "" || r.URL.RawQuery == "" {
			r.URL.RawQuery = tQuery + r.URL.RawQuery
		}
	}

	proxy := &httputil.ReverseProxy{Director: director}

	proxy.ModifyResponse = func(response *http.Response) error {
		content, err := io.ReadAll(response.Body)
		if err != nil {
			log.Print("Ошибка в изменении ответа - ", err)
			return err
		}
		log.Print(string(content))

		response.Body = io.NopCloser(bytes.NewReader(content))
		return nil
	}

	proxy.ServeHTTP(w, r)

}

func StartServer(service Service) error {
	s := &Service{
		Auth: service.Auth,
		User: service.User,
	}

	err := http.ListenAndServe(":8888", s)
	if err != nil {
		return err
	}

	return nil
}
