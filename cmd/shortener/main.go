package main

import (
	"log"

	"github.com/sSmok/ya-shortener/internal/app"
)

func main() {
	newApp, err := app.NewApp()
	if err != nil {
		log.Fatalf("failed to create app: %v", err)
	}
	err = newApp.Run()
	if err != nil {
		log.Fatalf("failed to run app: %v", err)
	}
}
