package handler

import (
	"fmt"
	"net/http"
	"strings"
)

func (h *Handler) proxy(w http.ResponseWriter, r *http.Request) {
	switch rootPath(r.URL.Path) {
	case "/user":
		h.service.Redirect(w, r)
	default:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "")
	}
}

func path(r *http.Request) string {
	index := strings.LastIndex(r.RequestURI, "/") + 1
	servicePath := r.RequestURI[index:]
	return servicePath
}

func rootPath(path string) string {
	var slashCount int
	var root string

	for _, ch := range path {
		if slashCount == 2 {
			root += string(ch)
		}

		if ch == '/' {
			if slashCount == 1 {
				slashCount++
				root += string(ch)
			} else {
				slashCount++
			}
		}
	}

	root = root[:len(root)-1]

	return root
}
