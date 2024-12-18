package main

import (
	"log"

	"github.com/sSmok/ya-shortener/internal/api/link"
)

func main() {
	api := link.NewAPI()
	err := api.Run()
	if err != nil {
		log.Fatalf(`fail to run server: %v`, err)
	}
}
