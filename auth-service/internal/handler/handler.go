package handler

import (
	"Booksiary/auth-service/internal/service"
	"github.com/go-chi/chi/v5"
	"log/slog"
)

type Handler struct {
	Logger  *slog.Logger
	Service *service.Service
}

func NewHandler(service *service.Service, logger *slog.Logger) *Handler {
	return &Handler{
		Logger:  logger,
		Service: service,
	}
}

func (h *Handler) Init() *chi.Mux {
	r := chi.NewRouter()

	r.Route("/api", func(r chi.Router) {
		r.Route("/auth", func(r chi.Router) {
			r.Post("/sign-up", h.signUp())
		})
	})

	r.Route("/interaction", func(r chi.Router) {
		r.Route("/user", func(r chi.Router) {
			r.Get("/{login}", h.returnAccount())     // возвращает аккаунт пользователя
			r.Post("/save-account", h.saveAccount()) //получает аккаунт пользователя и записывает в бд
		})
	})

	return r
}
