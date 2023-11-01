package handler

import (
	"net/http"
	"strings"
)

func (h *Handler) proxy(w http.ResponseWriter, r *http.Request) {
	switch path(r) {
	case "user":
		h.service.Redirect(w, r)
	default:
		h.service.Redirect(w, r)
		/*w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "")*/
	}
}

func path(r *http.Request) string {
	index := strings.LastIndex(r.RequestURI, "/") + 1
	servicePath := r.RequestURI[index:]
	return servicePath
}
