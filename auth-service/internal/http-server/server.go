package http_server

import "net/http"

type Server struct {
	httpServer *http.Server
}

func (s *Server) Run(handlers http.Handler) error {
	s.httpServer = &http.Server{
		Addr:    ":8081",
		Handler: handlers,
	}

	return s.httpServer.ListenAndServe()
}
