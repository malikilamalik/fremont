package main

import (
	"log"

	"github.com/malikilamalik/freemont/server"
)

func main() {
	server := server.NewServer()

	log.Fatalf(server.Run().Error())
}
