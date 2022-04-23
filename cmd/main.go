package main

import (
	"OrganizerApi"
	"log"
)

func main() {
	srv := new(OrganizerApi.Server)
	if err := srv.Run("8000"); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}
