package handler

import (
	"Booksiary/user-service/internal/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) Init() *gin.Engine {
	router := gin.New()

	api := router.Group("/api")
	{
		user := api.Group("/user")
		{
			user.POST("create", h.createUser)
			user.GET(":id", h.getUser)
			user.DELETE(":id", h.deleteUser)
			user.PUT(":id", h.updateUser)
		}
	}

	return router
}
