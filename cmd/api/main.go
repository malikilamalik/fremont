package main

import (
	"log"

	"github.com/malikilamalik/fremont/server"
)

func main() {
	app := server.NewServer()
	log.Fatalf(app.Run().Error())
}
