package handlers

import (
	"html/template"
	"net/http"
	"path/filepath"
	"strconv"
	"time"

	"github.com/ryantrue/contractkeeper/internal/models"
	"github.com/ryantrue/contractkeeper/internal/services"
	"github.com/sirupsen/logrus"
)

type ContractHandler struct {
	service           services.ContractService
	contractorService services.ContractorService
	logger            *logrus.Logger
}

func NewContractHandler(service services.ContractService, contractorService services.ContractorService, logger *logrus.Logger) *ContractHandler {
	return &ContractHandler{service: service, contractorService: contractorService, logger: logger}
}

func (h *ContractHandler) CreatePageHandler(w http.ResponseWriter, r *http.Request) {
	contractors, err := h.contractorService.GetAllContractors()
	if err != nil {
		h.logger.Errorf("Error getting contractors: %v", err)
		http.Error(w, "Could not get contractors", http.StatusInternalServerError)
		return
	}

	data := struct {
		Contractors     []models.Contractor
		PaymentAccounts []string
	}{
		Contractors: contractors,
		PaymentAccounts: []string{
			"40602810009800000077",
			"40602810438000000216",
			"40602810238000000212",
			"40602810638000000210",
			"40602810838000000214",
			"40602810752800000010",
			"40602810309800000078",
			"40602810052800000011",
		},
	}
	tmplPath := filepath.Join("templates", "create_contract.html")
	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		h.logger.Errorf("Could not load template: %v", err)
		http.Error(w, "Could not load template", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, data)
}

func (h *ContractHandler) CreateHandler(w http.ResponseWriter, r *http.Request) {
	contractorID, _ := strconv.Atoi(r.FormValue("contractor_id"))
	contractType := r.FormValue("type")
	number := r.FormValue("number")
	date, _ := time.Parse("2006-01-02", r.FormValue("date"))
	initiator := r.FormValue("initiator")
	amount, _ := strconv.ParseFloat(r.FormValue("amount"), 64)
	subject := r.FormValue("subject")
	status := r.FormValue("status")
	startDate, _ := time.Parse("2006-01-02", r.FormValue("start_date"))
	endDate, _ := time.Parse("2006-01-02", r.FormValue("end_date"))
	paymentProcedure := r.FormValue("payment_procedure")
	isRegular := r.FormValue("is_regular") == "on"
	article := r.FormValue("article")
	paymentAccounts := r.Form["payment_account"]

	contract := models.Contract{
		ContractorID:     uint(contractorID),
		Type:             contractType,
		Number:           number,
		Date:             date,
		Initiator:        initiator,
		Amount:           amount,
		Subject:          subject,
		Status:           status,
		StartDate:        startDate,
		EndDate:          endDate,
		PaymentProcedure: paymentProcedure,
		IsRegular:        isRegular,
		Article:          article,
		PaymentAccounts:  paymentAccounts,
	}

	err := h.service.CreateContract(&contract)
	if err != nil {
		h.logger.Errorf("Error creating contract: %v", err)
		http.Error(w, "Could not create contract", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/contracts", http.StatusSeeOther)
}

func (h *ContractHandler) ListHandler(w http.ResponseWriter, r *http.Request) {
	contracts, err := h.service.GetAllContracts()
	if err != nil {
		h.logger.Errorf("Error getting contracts: %v", err)
		http.Error(w, "Could not get contracts", http.StatusInternalServerError)
		return
	}

	data := struct {
		Contracts []models.Contract
	}{
		Contracts: contracts,
	}
	tmplPath := filepath.Join("templates", "contracts.html")
	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		h.logger.Errorf("Could not load template: %v", err)
		http.Error(w, "Could not load template", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, data)
}
