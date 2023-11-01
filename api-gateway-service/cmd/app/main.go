package main

import (
	"Booksiary/api-gateway-service/internal/handler"
	"Booksiary/api-gateway-service/internal/service"
	"net/http"
)

func main() {
	m := service.NewProxyAddrMap()
	m.M["/user"] = service.Handle{
		Host: "localhost",
		Port: "8080",
	}

	service := service.NewService(m)
	handler := handler.NewHandler(service)

	http.ListenAndServe(":8888", handler.Init())
}
