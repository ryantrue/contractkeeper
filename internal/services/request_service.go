package services

import (
	"github.com/ryantrue/contractkeeper/internal/models"
	"github.com/ryantrue/contractkeeper/internal/repositories"
)

type RequestService interface {
	CreateRequest(request *models.Request) error
	GetAllRequests() ([]models.Request, error)
	GetRequestByID(id uint) (*models.Request, error)
	UpdateRequest(request *models.Request) error
	DeleteRequest(id uint) error
}

type requestService struct {
	repo repositories.RequestRepository
}

func NewRequestService(repo repositories.RequestRepository) RequestService {
	return &requestService{repo: repo}
}

func (s *requestService) CreateRequest(request *models.Request) error {
	return s.repo.Save(request)
}

func (s *requestService) GetAllRequests() ([]models.Request, error) {
	return s.repo.FindAll()
}

func (s *requestService) GetRequestByID(id uint) (*models.Request, error) {
	return s.repo.FindByID(id)
}

func (s *requestService) UpdateRequest(request *models.Request) error {
	return s.repo.Update(request)
}

func (s *requestService) DeleteRequest(id uint) error {
	return s.repo.Delete(id)
}
