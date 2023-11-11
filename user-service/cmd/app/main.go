package main

import (
	"Booksiary/user-service/internal/handler"
	http_server "Booksiary/user-service/internal/http-server"
	"Booksiary/user-service/internal/mail"
	"Booksiary/user-service/internal/repository"
	"Booksiary/user-service/internal/repository/badger"
	"Booksiary/user-service/internal/repository/postgres"
	"Booksiary/user-service/internal/service"
	mailExternal "github.com/xhit/go-simple-mail/v2"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	mailClient, err := mail.NewMailClient(mail.Config{
		Host:           "smtp.mail.ru",
		Username:       "work.polyanskiy@mail.ru",
		Password:       "fApNn2LEpCzhzA6L9pKe",
		EmailAddress:   "work.polyanskiy@mail.ru",
		Port:           587,
		Encryption:     mailExternal.EncryptionTLS,
		ConnectTimeout: 15 * time.Second,
		SendTimeout:    15 * time.Second,
	})
	if err != nil {
		log.Fatalf("Ошибка в создании email клиента - %s", err)
	}
	memoryDB, err := badger.NewBadgerDB()
	if err != nil {
		log.Fatalf("Ошибка в инициализации badger - %v\n", err)
	}
	db, err := postgres.NewPostgresDB(postgres.Config{
		Host:     "localhost",
		Port:     "5480",
		Username: "konstantin",
		Password: "user12345",
		DBName:   "postgres",
		SSLMode:  "disable",
	})
	if err != nil {
		log.Fatalf("Ошибка в инициализации postgres - %v\n", err)
	}

	repository := repository.NewRepository(db, memoryDB)
	services := service.NewService(repository, *mailClient)
	handlers := handler.NewHandler(services)

	server := new(http_server.Server)

	go func() {
		if err := server.Run("8080", handlers.Init()); err != nil {
			log.Print("Ошибка в запуске сервера - ", err)
		}
	}()

	log.Print("Сервис Пользователей запущен")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
}
