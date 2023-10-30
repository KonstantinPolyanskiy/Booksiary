package handler

import (
	custom_response "Booksiary/authorization-service/internal/custom-response"
	"Booksiary/authorization-service/internal/types"
	"errors"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"io"
	"net/http"
)

func (h *Handler) createUser(chiCtx *chi.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var input types.SaveUser
		err := render.DecodeJSON(r.Body, &input)
		if errors.Is(err, io.EOF) {
			h.L.Error("тело пустое")
			custom_response.NewErrorResponse(w, r, http.StatusBadRequest, "empty body")
			return
		}
		if err != nil {
			h.L.Error("некорректное тело")
			custom_response.NewErrorResponse(w, r, http.StatusBadRequest, "неправильное тело")
			return
		}

		err = h.Service.Creator.UserCode(input)
		if err != nil {
			h.L.Error("idk")
			custom_response.NewErrorResponse(w, r, http.StatusInternalServerError, err.Error())
			return
		}

		render.JSON(w, r, "Код отправлен на почту")

	}
}
func (h *Handler) createUserCallback(chi *chi.Context) http.HandlerFunc {
	//TODO: коллбек на создание пользователя
	panic("сделай")
}

func responseUserCreateOK(wr http.ResponseWriter, r *http.Request, create types.User, id int) {
	render.JSON(wr, r, types.UserCreate{
		Id:   id,
		User: create,
	})
}
