package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) getUser(ctx *gin.Context) {
	h.service.Creator.Create()
	ctx.Status(http.StatusOK)
}
