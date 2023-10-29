package custom_response

import (
	"github.com/go-chi/render"
	"net/http"
)

type errorResponse struct {
	Message string `json:"message"`
}

func NewErrorResponse(rw http.ResponseWriter, r *http.Request, code int, message string) {
	render.Status(r, code)
	render.JSON(rw, r, errorResponse{Message: message})
}
