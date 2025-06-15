package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/dixon961/go-movie-downloader-backend/internal/app"
	"github.com/dixon961/go-movie-downloader-backend/internal/config"
	"github.com/dixon961/go-movie-downloader-backend/internal/handler"
	"github.com/dixon961/go-movie-downloader-backend/internal/repository"
	"github.com/go-chi/chi/v5"
)

func main() {
	cfg := config.Load()
	if cfg.JackettAPIKey == "" || cfg.JackettBaseURL == "" {
		log.Fatal("Jackett configuration (JACKETT_BASE_URL, JACKETT_API_KEY) is missing")
	}

	httpClient := &http.Client{Timeout: 10 * time.Second}
	jackettClient := repository.NewJackettClient(httpClient, cfg.JackettBaseURL, cfg.JackettAPIKey)

	searchService := app.NewSearchService(jackettClient)

	httpHandler := handler.NewHTTPHandler(searchService)

	router := chi.NewRouter()
	httpHandler.SetupRoutes(router)

	addr := fmt.Sprintf(":%s", cfg.Port)
	log.Printf("Starting server on %s", addr)

	if err := http.ListenAndServe(addr, router); err != nil {
		log.Fatalf("could not start server: %v", err)
	}
}
