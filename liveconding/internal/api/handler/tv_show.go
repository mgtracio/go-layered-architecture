package handler

import (
	"encoding/json"
	"layered_architecture/internal/service"
	"log"
	"net/http"
)

type TVShowHandler struct {
	service service.TVShowService
}

func NewTVShowHandler(service service.TVShowService) *TVShowHandler {
	return &TVShowHandler{service: service}
}

// FetchTvShows fetches tv show feeds and inserting into DB
func (h *TVShowHandler) FetchTvShows(w http.ResponseWriter, r *http.Request) {
	err := h.service.FetchFeeds()
	if err != nil {
		log.Printf("Error on FetchTvShows: %v\n", err)
	}
	// Return the user as a JSON response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Feeds fetched successfully"})
}
