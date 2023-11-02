package handler

import (
	"fmt"
	"log"
	"net/http"
)

func (h *Handler) proxy(w http.ResponseWriter, r *http.Request) {
	switch rootPath(r) {
	case "/user":
		h.service.Redirect(w, r)
	default:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "")
	}
}

/*func path(r *http.Request) string {
	index := strings.LastIndex(r.RequestURI, "/") + 1
	servicePath := r.RequestURI[index:]
	return servicePath
}*/

func rootPath(r *http.Request) string {
	var slashCount int
	var root string

	path := r.URL.Path

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

	var rootRes string

	if len(root) > 0 {
		rootRes = root[:len(root)-1]
	}

	log.Printf("Корень - %s", rootRes)

	return rootRes
}
