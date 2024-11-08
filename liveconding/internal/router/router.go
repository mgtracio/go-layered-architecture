package router

import (
	"layered_architecture/internal/api/handler"
	"net/http"
)

func NewRouter(handler *handler.TVShowHandler) http.Handler {
	server := http.NewServeMux()
	server.HandleFunc("/fetch-tv-shows", handler.FetchTvShows)
	// Return the multiplexer which serves as the router
	return server
}
