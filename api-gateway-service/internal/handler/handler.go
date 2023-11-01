package handler

import (
	"Booksiary/api-gateway-service/internal/service"
	"net/http"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) Init() *http.ServeMux {
	router := http.NewServeMux()

	router.HandleFunc("/", h.proxy)

	return router
}
