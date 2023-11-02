package lib

import (
	"github.com/go-chi/render"
	"net/http"
)

type ErrorResponse struct {
	Message string `json:"message"`
}

// NewErrorResponse устанавливает переданный заголовок ответа
// и передает сообщение в формате JSON
func NewErrorResponse(w http.ResponseWriter, r *http.Request, code int, message string) {
	w.WriteHeader(code)
	render.JSON(w, r, ErrorResponse{Message: message})
}
