package main

import (
	"log"
	"os"

	"github.com/nixpig/bakery/internal/server"
)

func main() {
	appConfig := server.AppConfig{Port: "8080"}

	if err := server.Start(appConfig); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
