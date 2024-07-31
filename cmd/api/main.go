package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/ryantrue/contractkeeper/internal/config"
	"github.com/ryantrue/contractkeeper/internal/database"
	"github.com/ryantrue/contractkeeper/internal/handlers"
	"github.com/ryantrue/contractkeeper/internal/repositories"
	"github.com/ryantrue/contractkeeper/internal/services"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Initialize database
	database.InitDB(cfg)

	// Set up repository
	requestRepo := repositories.NewRequestRepository(database.DB)

	// Set up service
	requestService := services.NewRequestService(requestRepo)

	// Set up handlers
	requestHandler := handlers.NewRequestHandler(requestService)

	// Set up router
	r := mux.NewRouter()
	r.HandleFunc("/", requestHandler.RenderRequestForm).Methods("GET")
	r.HandleFunc("/submit", requestHandler.SubmitFormHandler).Methods("POST")
	r.HandleFunc("/requests", requestHandler.ViewRequests).Methods("GET")
	r.HandleFunc("/edit/{id}", requestHandler.EditRequestForm).Methods("GET")
	r.HandleFunc("/update/{id}", requestHandler.UpdateRequestHandler).Methods("POST")
	r.HandleFunc("/delete/{id}", requestHandler.DeleteRequestHandler).Methods("POST")
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	log.Println("Server started at http://localhost:8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalf("Could not start server: %v", err)
	}
}
