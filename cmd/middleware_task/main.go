package main

import (
	"log"

	"github.com/asliddinberdiev/middleware_task/internal/app/pkg/app"
)

func main() {
	a, err := app.New()
	if err != nil {
		log.Fatal(err)
	}

	err = a.Run()
	if err != nil {
		log.Fatal(err)
	}
}
