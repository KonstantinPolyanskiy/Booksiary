package internal

import "github.com/gin-gonic/gin"

type ER struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func ErrorResponse(ctx *gin.Context, code int, message string) {
	ctx.JSON(code, message)
}
