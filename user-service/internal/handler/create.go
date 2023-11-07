package handler

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func (h *Handler) createUser(ctx *gin.Context) {
	h.service.Creator.Create()
	ctx.Status(http.StatusOK)
}

func (h *Handler) createUserWithExternal(ctx *gin.Context) {
	var input string
	err := ctx.BindJSON(&input)
	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		return
	}
	log.Print(input)
}
