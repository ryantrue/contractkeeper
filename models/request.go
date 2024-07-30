package models

import "gorm.io/gorm"

type Request struct {
	gorm.Model
	Contractor     string  `json:"contractor"`
	Contract       string  `json:"contract"`
	ContractDate   string  `json:"contract_date"`
	Subject        string  `json:"subject"`
	Amount         float64 `json:"amount"`
	ContractAmount float64 `json:"contract_amount"`
	Article        string  `json:"article"`
	StartDate      string  `json:"start_date"`
	Deadline       string  `json:"deadline"`
	PaymentAccount string  `json:"payment_account"`
	DeadlineDate   string  `json:"deadline_date"`
}
