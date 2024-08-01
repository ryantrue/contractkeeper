package handlers

import (
	"html/template"
	"net/http"
	"path/filepath"

	"github.com/ryantrue/contractkeeper/internal/models"
	"github.com/ryantrue/contractkeeper/internal/services"
	"github.com/sirupsen/logrus"
)

type ContractorHandler struct {
	service services.ContractorService
	logger  *logrus.Logger
}

func NewContractorHandler(service services.ContractorService, logger *logrus.Logger) *ContractorHandler {
	return &ContractorHandler{service: service, logger: logger}
}

func (h *ContractorHandler) CreatePageHandler(w http.ResponseWriter, r *http.Request) {
	tmplPath := filepath.Join("templates", "create_contractor.html")
	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		h.logger.Errorf("Could not load template: %v", err)
		http.Error(w, "Could not load template", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}

func (h *ContractorHandler) CreateHandler(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	inn := r.FormValue("inn")
	ogrn := r.FormValue("ogrn")
	requisites := r.FormValue("requisites")

	contractor := models.Contractor{
		Name:       name,
		INN:        inn,
		OGRN:       ogrn,
		Requisites: requisites,
	}

	err := h.service.CreateContractor(&contractor)
	if err != nil {
		h.logger.Errorf("Error creating contractor: %v", err)
		http.Error(w, "Could not create contractor", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/contractors", http.StatusSeeOther)
}

func (h *ContractorHandler) ListHandler(w http.ResponseWriter, r *http.Request) {
	contractors, err := h.service.GetAllContractors()
	if err != nil {
		h.logger.Errorf("Error getting contractors: %v", err)
		http.Error(w, "Could not get contractors", http.StatusInternalServerError)
		return
	}

	data := struct {
		Contractors []models.Contractor
	}{
		Contractors: contractors,
	}
	tmplPath := filepath.Join("templates", "contractors.html")
	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		h.logger.Errorf("Could not load template: %v", err)
		http.Error(w, "Could not load template", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, data)
}
