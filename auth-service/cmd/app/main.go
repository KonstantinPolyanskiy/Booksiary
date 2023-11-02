package main

import (
	"Booksiary/auth-service/internal/handler"
	http_server "Booksiary/auth-service/internal/http-server"
	"Booksiary/auth-service/internal/service"
	"log"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	logger := slog.Logger{}
	//repository := repository.NewRepository()

	services := service.NewService()
	handlers := handler.NewHandler(services, &logger)

	server := new(http_server.Server)

	go func() {
		log.Fatalf("Ошибка в запуске сервера Авторизации - %v", server.Run(handlers.Init()))
	}()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
}
