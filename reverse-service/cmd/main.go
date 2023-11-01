package main

import (
	"Booksiary/reverse-service/internal/proxy"
	"log"
)

func main() {
	s := proxy.Service{
		Auth: &proxy.Handle{
			Host: "localhost",
			Port: "8080",
		},
		User: &proxy.Handle{
			Host: "localhost",
			Port: "8081",
		},
	}
	err := proxy.StartServer(s)
	if err != nil {
		log.Fatal(err)
	}
}
