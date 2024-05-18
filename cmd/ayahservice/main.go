package main

import (
	"log"
	"net/http"

	"github.com/Aadil101/ayah-backend/pkg/handler"
)

func main() {
	server := &http.Server{
		Addr:    ":8080",
		Handler: handler.NewHandler(),
	}
	log.Println("Listening...")
	log.Fatal(server.ListenAndServe())
}
