package main

import (
	"Booksiary/authorization-service/config"
	"Booksiary/authorization-service/internal/http-server"
	"Booksiary/authorization-service/pkg/handler"
	"Booksiary/authorization-service/pkg/service"
	"log"
	"log/slog"
	"time"
)

func main() {
	testConfig := config.ServerConfig{
		Port:         "8080",
		WriteTimeout: 5 * time.Second,
		ReadTimeout:  5 * time.Second,
	}

	services := service.NewService()
	handlers := handler.NewHandler(services, slog.Default())

	server := new(http_server.Server)
	err := server.Run(testConfig, handlers.Init(), slog.Logger{})
	if err != nil {
		log.Fatal("Ошибка в запуске сервера - ", err.Error())
	}

}
