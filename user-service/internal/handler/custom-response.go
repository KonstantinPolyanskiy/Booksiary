package handler

import "github.com/gin-gonic/gin"

type errResponse struct {
	Message string `json:"message"`
}

func NewErrResponse(ctx *gin.Context, statusCode int, message string) {
	ctx.AbortWithStatusJSON(statusCode, errResponse{Message: message})
}
