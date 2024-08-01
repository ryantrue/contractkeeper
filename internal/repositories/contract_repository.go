package repositories

import (
	"github.com/jinzhu/gorm"
	"github.com/ryantrue/contractkeeper/internal/models"
)

type ContractRepository interface {
	Save(contract *models.Contract) error
	FindAll() ([]models.Contract, error)
	FindByID(id uint) (*models.Contract, error)
}

type contractRepository struct {
	DB *gorm.DB
}

func NewContractRepository(db *gorm.DB) ContractRepository {
	return &contractRepository{DB: db}
}

func (r *contractRepository) Save(contract *models.Contract) error {
	return r.DB.Create(contract).Error
}

func (r *contractRepository) FindAll() ([]models.Contract, error) {
	var contracts []models.Contract
	err := r.DB.Preload("Contractor").Find(&contracts).Error
	return contracts, err
}

func (r *contractRepository) FindByID(id uint) (*models.Contract, error) {
	var contract models.Contract
	err := r.DB.Preload("Contractor").First(&contract, id).Error
	return &contract, err
}
