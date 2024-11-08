package main

import (
	"fmt"
	"layered_architecture/internal/api/handler"
	"layered_architecture/internal/db"
	"layered_architecture/internal/repository"
	"layered_architecture/internal/router"
	"layered_architecture/internal/service"
	"layered_architecture/pkg/config"
	"log"
	"net/http"
)

func main() {
	// Initialize the database connection
	db.Init()
	defer db.Close()

	// Singleton config instance
	cfg := config.GetConfig()

	// Initialize the repository layer
	repo := repository.NewTVShowRepository(db.DB)

	// Initialize the service layer
	service := service.NewTVShowService(repo, cfg)

	// Initialize the handler layer
	handler := handler.NewTVShowHandler(service)

	// Set up the router using the native net/http
	r := router.NewRouter(handler)

	// Start the HTTP server
	log.Printf("Starting server on :%d", cfg.AppPort)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", cfg.AppPort), r))
}
