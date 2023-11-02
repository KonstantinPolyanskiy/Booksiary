package main

import (
	"Booksiary/user-service/internal/handler"
	http_server "Booksiary/user-service/internal/http-server"
	"Booksiary/user-service/internal/service"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	services := service.NewService()
	handlers := handler.NewHandler(services)

	server := new(http_server.Server)

	go func() {
		if err := server.Run("8080", handlers.Init()); err != nil {
			log.Print("Ошибка в запуске сервера - ", err)
		}
	}()

	log.Println("Сервер слушает 8080")
	log.Print("Сервис Пользователей запущен")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
}
