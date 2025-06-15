package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/dixon961/go-movie-downloader-backend/internal/app"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type HTTPHandler struct {
	searchService *app.SearchService
}

func NewHTTPHandler(searchService *app.SearchService) *HTTPHandler {
	return &HTTPHandler{
		searchService: searchService,
	}
}

func (h *HTTPHandler) SetupRoutes(router *chi.Mux) {
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello from Go!"))
	})

	router.Route("/api", func(r chi.Router) {
		r.Get("/search", h.searchTorrents)
	})
}

func (h *HTTPHandler) searchTorrents(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("q")
	if query == "" {
		http.Error(w, "Query parameter 'q' is required", http.StatusBadRequest)
		return
	}

	results, err := h.searchService.SearchTorrents(r.Context(), query)
	if err != nil {
		log.Printf("ERROR: failed to search torrents for query '%s': %v", query, err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(results)
}
