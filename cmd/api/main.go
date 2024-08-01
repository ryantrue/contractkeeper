package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/ryantrue/contractkeeper/internal/config"
	"github.com/ryantrue/contractkeeper/internal/handlers"
	"github.com/ryantrue/contractkeeper/internal/repositories"
	"github.com/ryantrue/contractkeeper/internal/services"
	"github.com/sirupsen/logrus"
)

func main() {
	cfg := config.LoadConfig()

	db, err := gorm.Open("postgres", cfg.DBURL)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	logger := logrus.New()
	logger.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})

	contractorRepo := repositories.NewContractorRepository(db)
	contractRepo := repositories.NewContractRepository(db)

	contractorService := services.NewContractorService(contractorRepo)
	contractService := services.NewContractService(contractRepo)

	contractorHandler := handlers.NewContractorHandler(contractorService, logger)
	contractHandler := handlers.NewContractHandler(contractService, contractorService, logger)

	r := mux.NewRouter()

	r.HandleFunc("/contractors", contractorHandler.ListHandler).Methods("GET")
	r.HandleFunc("/contractors/new", contractorHandler.CreatePageHandler).Methods("GET")
	r.HandleFunc("/contractors", contractorHandler.CreateHandler).Methods("POST")

	r.HandleFunc("/contracts", contractHandler.ListHandler).Methods("GET")
	r.HandleFunc("/contracts/new", contractHandler.CreatePageHandler).Methods("GET")
	r.HandleFunc("/contracts", contractHandler.CreateHandler).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", r))
}
