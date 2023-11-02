package handler

import (
	"Booksiary/user-service/internal"
	"github.com/gin-gonic/gin"
	"net/http"
)

type User struct {
	Name    string `json:"name"`
	Surname string `json:"surname"`
}

func (h *Handler) createUser(ctx *gin.Context) {
	var user User
	err := ctx.BindJSON(&user)
	if err != nil {
		internal.ErrorResponse(ctx, http.StatusBadRequest, "неправильное тело запроса")
		return
	}

	id, err := h.service.Creator.Create()
	if err != nil {
		internal.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, id)
}
