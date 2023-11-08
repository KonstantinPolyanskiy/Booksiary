package main

import (
	"Booksiary/reverse-service/internal/proxy"
	"log"
)

func main() {

	s := proxy.Service{HandleMap: map[string]proxy.Handle{"user": {
		Host: "localhost",
		Port: "8080",
	}}}
	err := proxy.StartServer(s)
	if err != nil {
		log.Fatal(err)
	}
}
