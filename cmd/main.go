package main

import (
	"github.com/Lekhman-vold/OrganizerApi"
	"github.com/Lekhman-vold/OrganizerApi/pkg/handler"
	"log"
)

func main() {
	handlers := new(handler.Handler)
	srv := new(OrganizerApi.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}
