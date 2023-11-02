package http_server

import (
	"Booksiary/authorization-service/config"
	"context"
	"log/slog"
	"net/http"
)

type Server struct {
	httpServer *http.Server
	l          slog.Logger
}

func (s *Server) Run(cfg config.ServerConfig, handlers http.Handler, logger slog.Logger) error {
	s.l = logger

	s.httpServer = &http.Server{
		Addr:         ":" + cfg.Port,
		Handler:      handlers,
		ReadTimeout:  cfg.ReadTimeout,
		WriteTimeout: cfg.WriteTimeout,
	}

	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
