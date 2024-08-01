package repositories

import (
	"github.com/jinzhu/gorm"
	"github.com/ryantrue/contractkeeper/internal/models"
	"github.com/sirupsen/logrus"
)

type RequestRepository interface {
	Save(request *models.Request) error
	FindAll() ([]models.Request, error)
	FindByID(id uint) (*models.Request, error)
	Update(request *models.Request) error
	Delete(id uint) error
}

type requestRepository struct {
	DB     *gorm.DB
	logger *logrus.Logger
}

func NewRequestRepository(db *gorm.DB, logger *logrus.Logger) RequestRepository {
	return &requestRepository{DB: db, logger: logger}
}

func (r *requestRepository) Save(request *models.Request) error {
	return r.DB.Create(request).Error
}

func (r *requestRepository) FindAll() ([]models.Request, error) {
	var requests []models.Request
	err := r.DB.Where("deleted_at IS NULL").Find(&requests).Error
	if err != nil {
		r.logger.Errorf("Error finding requests: %v", err)
	}
	return requests, err
}

func (r *requestRepository) FindByID(id uint) (*models.Request, error) {
	var request models.Request
	err := r.DB.Where("deleted_at IS NULL").First(&request, id).Error
	if err != nil {
		r.logger.Errorf("Error finding request by ID: %v", err)
	}
	return &request, err
}

func (r *requestRepository) Update(request *models.Request) error {
	return r.DB.Save(request).Error
}

func (r *requestRepository) Delete(id uint) error {
	return r.DB.Model(&models.Request{}).Where("id = ?", id).Update("deleted_at", gorm.Expr("NOW()")).Error
}
