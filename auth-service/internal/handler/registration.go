package handler

import (
	"Booksiary/auth-service/internal/domain"
	"Booksiary/auth-service/internal/handler/lib"
	"errors"
	"github.com/go-chi/render"
	"io"
	"net/http"
)

func (h *Handler) signUp() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var input domain.User

		err := render.DecodeJSON(r.Body, &input)
		if errors.Is(err, io.EOF) {
			h.Logger.Warn("Пришло пустое тело от ", r.RemoteAddr)
			lib.NewErrorResponse(w, r, http.StatusBadRequest, "empty request body")
			return
		}
		if err != nil {
			h.Logger.Warn("Пришло некорректное тело от ", r.RemoteAddr)
			lib.NewErrorResponse(w, r, http.StatusInternalServerError, "incorrect body")
			return
		}

		id, err := h.Service.AuthService.Create(input)

		render.JSON(w, r, map[string]interface{}{
			"ID": id,
		})
	}
}
