package handlers

import (
	"github.com/ryantrue/contractkeeper/internal/models"
	"html/template"
	"net/http"
	"path/filepath"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/ryantrue/contractkeeper/internal/services"
	"github.com/sirupsen/logrus"
)

type RequestHandler struct {
	service services.RequestService
	logger  *logrus.Logger
}

func NewRequestHandler(service services.RequestService, logger *logrus.Logger) *RequestHandler {
	return &RequestHandler{service: service, logger: logger}
}

func (h *RequestHandler) Dashboard(w http.ResponseWriter, r *http.Request) {
	tasks, err := h.service.GetAllRequests()
	if err != nil {
		h.logger.Errorf("Error getting requests: %v", err)
		http.Error(w, "Could not get requests", http.StatusInternalServerError)
		return
	}

	data := struct {
		TaskCount int
		Tasks     []string
	}{
		TaskCount: len(tasks),
		Tasks:     make([]string, len(tasks)),
	}

	for i, task := range tasks {
		data.Tasks[i] = task.Subject
	}

	tmplPath := filepath.Join("templates", "dashboard.html")
	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		h.logger.Errorf("Could not load template: %v", err)
		http.Error(w, "Could not load template", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, data)
}

func (h *RequestHandler) ViewRequests(w http.ResponseWriter, r *http.Request) {
	requests, err := h.service.GetAllRequests()
	if err != nil {
		h.logger.Errorf("Error getting requests: %v", err)
		http.Error(w, "Could not get requests", http.StatusInternalServerError)
		return
	}

	data := struct {
		Requests []models.Request
	}{
		Requests: requests,
	}
	tmplPath := filepath.Join("templates", "requests.html")
	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		h.logger.Errorf("Could not load template: %v", err)
		http.Error(w, "Could not load template", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, data)
}

func (h *RequestHandler) EditRequestForm(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	request, err := h.service.GetRequestByID(uint(id))
	if err != nil {
		h.logger.Errorf("Error getting request: %v", err)
		http.Error(w, "Could not get request", http.StatusInternalServerError)
		return
	}

	data := struct {
		Request         models.Request
		DeadlineOptions []string
		PaymentAccounts []string
	}{
		Request:         *request,
		DeadlineOptions: []string{"день в день", "7 рабочих дней", "фиксированная дата"},
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

	tmplPath := filepath.Join("templates", "edit_request.html")
	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		h.logger.Errorf("Could not load template: %v", err)
		http.Error(w, "Could not load template", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, data)
}

func (h *RequestHandler) UpdateRequestHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	contractor := r.FormValue("contractor")
	contract := r.FormValue("contract")
	contractDate := r.FormValue("contractDate")
	subject := r.FormValue("subject")
	amount, _ := strconv.ParseFloat(r.FormValue("amount"), 64)
	contractAmount, _ := strconv.ParseFloat(r.FormValue("contractAmount"), 64)
	article := r.FormValue("article")
	startDate := r.FormValue("startDate")
	deadline := r.FormValue("deadline")
	paymentAccount := r.FormValue("paymentAccount")

	var deadlineDate string
	startDateParsed, _ := time.Parse("2006-01-02", startDate)
	switch deadline {
	case "день в день":
		deadlineDate = startDateParsed.Format("2006-01-02")
	case "7 рабочих дней":
		deadlineDate = addBusinessDays(startDateParsed, 7).Format("2006-01-02")
	case "фиксированная дата":
		deadlineDate = r.FormValue("fixedDate")
	}

	request, err := h.service.GetRequestByID(uint(id))
	if err != nil {
		h.logger.Errorf("Error getting request: %v", err)
		http.Error(w, "Could not get request", http.StatusInternalServerError)
		return
	}

	request.Contractor = contractor
	request.Contract = contract
	request.ContractDate = contractDate
	request.Subject = subject
	request.Amount = amount
	request.ContractAmount = contractAmount
	request.Article = article
	request.StartDate = startDate
	request.Deadline = deadline
	request.PaymentAccount = paymentAccount
	request.DeadlineDate = deadlineDate

	err = h.service.UpdateRequest(request)
	if err != nil {
		h.logger.Errorf("Error updating request: %v", err)
		http.Error(w, "Could not update request", http.StatusInternalServerError)
		return
	}

	h.logger.Infof("Request %d updated", id)
	http.Redirect(w, r, "/requests", http.StatusSeeOther)
}

func (h *RequestHandler) DeleteRequestHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	err := h.service.DeleteRequest(uint(id))
	if err != nil {
		h.logger.Errorf("Error deleting request: %v", err)
		http.Error(w, "Could not delete request", http.StatusInternalServerError)
		return
	}

	h.logger.Infof("Request %d deleted", id)
	http.Redirect(w, r, "/requests", http.StatusSeeOther)
}

func addBusinessDays(date time.Time, days int) time.Time {
	for days > 0 {
		date = date.AddDate(0, 0, 1)
		if date.Weekday() != time.Saturday && date.Weekday() != time.Sunday {
			days--
		}
	}
	return date
}
