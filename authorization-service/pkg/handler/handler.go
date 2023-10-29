package handler

import (
	mwLogger "Booksiary/authorization-service/internal/http-server/middleware"
	"Booksiary/authorization-service/pkg/service"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log/slog"
)

type Handler struct {
	L       slog.Logger
	Service *service.Service
}

func NewHandler(service *service.Service, logger *slog.Logger) *Handler {
	return &Handler{
		L:       *logger,
		Service: service,
	}
}

func (h *Handler) Init() *chi.Mux {
	r := chi.NewRouter()

	r.Use(mwLogger.New(&h.L))
	r.Use(middleware.AllowContentType("application/json"))
	r.Route("/api", func(r chi.Router) {
		r.Post("/create-user", h.createUser(chi.NewRouteContext()))
		//r.Post("/create-user-callback", h.createUserCallback(chi.NewRouteContext()))

	})

	return r
}
