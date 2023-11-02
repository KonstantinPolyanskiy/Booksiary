package main

import (
	"Booksiary/reverse-service/internal/proxy"
	"log"
)

func main() {
	s := proxy.Service{
		Auth: &proxy.Handle{ //не реализованно
			Host: "127.0.0.1",
			Port: "9999",
		},
		User: &proxy.Handle{
			Host: "127.0.0.1",
			Port: "8080",
		},
	}
	err := proxy.StartServer(s)
	if err != nil {
		log.Fatal(err)
	}
}
