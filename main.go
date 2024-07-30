package main

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/pressly/goose/v3"
	"github.com/ryantrue/contractkeeper/database"
	"github.com/ryantrue/contractkeeper/models"
)

// RequestFormData represents the data structure for the request form.
type RequestFormData struct {
	Contractor      string
	Contract        string
	ContractDate    string
	Subject         string
	Amount          float64
	ContractAmount  float64
	Article         string
	StartDate       string
	DeadlineOptions []string
	PaymentAccounts []string
}

func main() {
	database.InitDB()
	runMigrations()

	r := mux.NewRouter()
	r.HandleFunc("/", renderRequestForm).Methods("GET")
	r.HandleFunc("/submit", submitFormHandler).Methods("POST")
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	log.Println("Server started at http://localhost:8080")
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatalf("Could not start server: %s\n", err)
	}
}

func runMigrations() {
	db, err := database.DB.DB()
	if err != nil {
		log.Fatalf("Failed to get database instance: %v", err)
	}

	if err := goose.SetDialect("sqlite3"); err != nil {
		log.Fatalf("Failed to set goose dialect: %v", err)
	}

	if err := goose.Up(db, "migrations"); err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}
}

func renderRequestForm(w http.ResponseWriter, r *http.Request) {
	// Define the data for the form
	data := RequestFormData{
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
	tmplPath := filepath.Join("templates", "request.html")
	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		http.Error(w, "Could not load template", http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, "Could not render template", http.StatusInternalServerError)
	}
}

func submitFormHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	contractor := r.FormValue("contractor")
	contract := r.FormValue("contract")
	contractDate := r.FormValue("contractDate")
	subject := r.FormValue("subject")
	amount := r.FormValue("amount")
	contractAmount := r.FormValue("contractAmount")
	article := r.FormValue("article")
	startDate := r.FormValue("startDate")
	deadline := r.FormValue("deadline")
	paymentAccount := r.FormValue("paymentAccount")

	var deadlineDate time.Time
	startDateParsed, _ := time.Parse("2006-01-02", startDate)
	switch deadline {
	case "день в день":
		deadlineDate = startDateParsed
	case "7 рабочих дней":
		deadlineDate = addBusinessDays(startDateParsed, 7)
	case "фиксированная дата":
		deadlineDateInput := r.FormValue("fixedDate")
		deadlineDate, _ = time.Parse("2006-01-02", deadlineDateInput)
	}

	request := models.Request{
		Contractor:     contractor,
		Contract:       contract,
		ContractDate:   contractDate,
		Subject:        subject,
		Amount:         parseAmount(amount),
		ContractAmount: parseAmount(contractAmount),
		Article:        article,
		StartDate:      startDate,
		Deadline:       deadline,
		PaymentAccount: paymentAccount,
		DeadlineDate:   deadlineDate.Format("2006-01-02"),
	}

	database.DB.Create(&request)

	log.Printf("Received form data: %+v", request)

	http.Redirect(w, r, "/", http.StatusSeeOther)
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

func parseAmount(amount string) float64 {
	parsedAmount, _ := strconv.ParseFloat(amount, 64)
	return parsedAmount
}
