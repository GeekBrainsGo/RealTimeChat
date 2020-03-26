package main

import (
	"chat/server"
	"log"
)

func main() {
	serv := server.New()
	log.Println("starting server")
	if err := serv.Start(); err != nil {
		log.Fatal(err)
	}
}
