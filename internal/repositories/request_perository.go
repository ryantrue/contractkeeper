package repositories

import (
	"github.com/ryantrue/contractkeeper/internal/models"

	"github.com/jinzhu/gorm"
)

type RequestRepository interface {
	Save(request *models.Request) error
	FindAll() ([]models.Request, error)
	FindByID(id uint) (*models.Request, error)
	Update(request *models.Request) error
	Delete(id uint) error
}

type requestRepository struct {
	DB *gorm.DB
}

func NewRequestRepository(db *gorm.DB) RequestRepository {
	return &requestRepository{DB: db}
}

func (r *requestRepository) Save(request *models.Request) error {
	return r.DB.Create(request).Error
}

func (r *requestRepository) FindAll() ([]models.Request, error) {
	var requests []models.Request
	err := r.DB.Find(&requests).Error
	return requests, err
}

func (r *requestRepository) FindByID(id uint) (*models.Request, error) {
	var request models.Request
	err := r.DB.First(&request, id).Error
	return &request, err
}

func (r *requestRepository) Update(request *models.Request) error {
	return r.DB.Save(request).Error
}

func (r *requestRepository) Delete(id uint) error {
	return r.DB.Delete(&models.Request{}, id).Error
}
