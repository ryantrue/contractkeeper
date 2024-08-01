package models

import "github.com/jinzhu/gorm"

type Contractor struct {
	gorm.Model
	Name       string `json:"name"`
	INN        string `json:"inn"`
	OGRN       string `json:"ogrn"`
	Requisites string `json:"requisites"`
}
