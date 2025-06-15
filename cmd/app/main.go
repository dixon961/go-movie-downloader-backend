package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/dixon961/go-movie-downloader-backend/internal/config"
	"github.com/dixon961/go-movie-downloader-backend/internal/handler"
)

func main() {
	cfg := config.Load()

	router := handler.NewRouter()

	addr := fmt.Sprintf(":%s", cfg.Port)
	log.Println("Starting server on %s", addr)

	if err := http.ListenAndServe(addr, router); err != nil {
		log.Fatalf("could not start server: %v", err)
	}
}
