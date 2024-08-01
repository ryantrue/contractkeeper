package repositories

import (
	"github.com/jinzhu/gorm"
	"github.com/ryantrue/contractkeeper/internal/models"
)

type ContractorRepository interface {
	Save(contractor *models.Contractor) error
	FindAll() ([]models.Contractor, error)
	FindByID(id uint) (*models.Contractor, error)
}

type contractorRepository struct {
	DB *gorm.DB
}

func NewContractorRepository(db *gorm.DB) ContractorRepository {
	return &contractorRepository{DB: db}
}

func (r *contractorRepository) Save(contractor *models.Contractor) error {
	return r.DB.Create(contractor).Error
}

func (r *contractorRepository) FindAll() ([]models.Contractor, error) {
	var contractors []models.Contractor
	err := r.DB.Find(&contractors).Error
	return contractors, err
}

func (r *contractorRepository) FindByID(id uint) (*models.Contractor, error) {
	var contractor models.Contractor
	err := r.DB.First(&contractor, id).Error
	return &contractor, err
}
