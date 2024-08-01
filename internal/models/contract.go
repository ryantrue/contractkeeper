package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Contract struct {
	gorm.Model
	ContractorID     uint       `json:"contractor_id"`
	Contractor       Contractor `gorm:"foreignkey:ContractorID"`
	Type             string     `json:"type"`
	Number           string     `json:"number"`
	Date             time.Time  `json:"date"`
	Initiator        string     `json:"initiator"`
	Amount           float64    `json:"amount"`
	Subject          string     `json:"subject"`
	Status           string     `json:"status"`
	StartDate        time.Time  `json:"start_date"`
	EndDate          time.Time  `json:"end_date"`
	PaymentProcedure string     `json:"payment_procedure"`
	IsRegular        bool       `json:"is_regular"`
	Article          string     `json:"article"`
	PaymentAccounts  []string   `json:"payment_accounts" gorm:"type:text[]"`
}
