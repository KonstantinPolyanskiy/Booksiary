package main

import (
	"Booksiary/api-gateway-service/internal/handler"
	"Booksiary/api-gateway-service/internal/service"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	m := service.NewProxyAddrMap()
	m.M["/user"] = service.Handle{
		Host: "localhost",
		Port: "8080",
	}

	service := service.NewService(m)
	handler := handler.NewHandler(service)

	go func() {
		log.Fatal(http.ListenAndServe(":8888", handler.Init()))
	}()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
}
