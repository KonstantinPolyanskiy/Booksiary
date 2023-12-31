package main

import (
	"Booksiary/auth-service/internal/handler"
	http_server "Booksiary/auth-service/internal/http-server"
	"Booksiary/auth-service/internal/repository"
	"Booksiary/auth-service/internal/service"
	"github.com/spf13/viper"
	"log"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("Ошибка в инициализации конфига - %v", err)
	}

	logger := slog.Logger{}

	db, err := repository.NewPostgresDB(repository.PostgresConfig{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: viper.GetString("db.password"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})
	if err != nil {
		log.Fatalf("Ошибка в созданни базы данных - %v", err)
	}
	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services, &logger)

	server := new(http_server.Server)

	go func() {
		log.Fatalf("Ошибка в запуске сервера Авторизации - %v", server.Run(handlers.Init()))
	}()
	log.Print("Сервер запустился")
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
}

func initConfig() error {
	viper.AddConfigPath("auth-service/configs")
	viper.SetConfigName("dev")

	return viper.ReadInConfig()
}
