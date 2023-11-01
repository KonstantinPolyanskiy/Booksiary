package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) deleteUser(ctx *gin.Context) {
	h.service.Delete()
	ctx.Status(http.StatusOK)
}
