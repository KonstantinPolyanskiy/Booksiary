package handler

import (
	"Booksiary/auth-service/internal/handler/lib"
	"Booksiary/auth-service/internal/service"
	"errors"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"io"
	"net/http"
)

func (h *Handler) getAccount() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		login := chi.URLParam(r, "login")
		if login == "" {
			lib.NewErrorResponse(w, r, http.StatusBadRequest, "пустой логин")
			return
		}
		account, err := h.Service.Account.Get(login)
		if err != nil {
			lib.NewErrorResponse(w, r, http.StatusInternalServerError, err.Error())
			return
		}

		render.JSON(w, r, &account)
	}
}

func (h *Handler) saveAccount() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var input service.UserAccountRequest

		err := render.DecodeJSON(r.Body, &input)
		if errors.Is(err, io.EOF) {
			lib.NewErrorResponse(w, r, http.StatusBadRequest, "empty body request")
			return
		} else if err != nil {
			lib.NewErrorResponse(w, r, http.StatusBadRequest, err.Error())
			return
		}

		err = h.Service.Account.Save(input)
		if err != nil {
			lib.NewErrorResponse(w, r, http.StatusInternalServerError, err.Error())
			return
		}

		w.WriteHeader(http.StatusOK)
		render.JSON(w, r, map[string]interface{}{
			"status": "Аккаунт сохранен",
		})
	}
}
