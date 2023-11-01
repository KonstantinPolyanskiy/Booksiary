package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) createUser(ctx *gin.Context) {
	h.service.Creator.Create()
	ctx.Status(http.StatusOK)
}
