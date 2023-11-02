package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) updateUser(ctx *gin.Context) {
	h.service.Updater.Update()
	ctx.Status(http.StatusOK)
}
