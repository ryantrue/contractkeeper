package services

import (
	"github.com/ryantrue/contractkeeper/internal/models"
	"github.com/ryantrue/contractkeeper/internal/repositories"
)

type ContractorService interface {
	CreateContractor(contractor *models.Contractor) error
	GetAllContractors() ([]models.Contractor, error)
	GetContractorByID(id uint) (*models.Contractor, error)
}

type contractorService struct {
	repo repositories.ContractorRepository
}

func NewContractorService(repo repositories.ContractorRepository) ContractorService {
	return &contractorService{repo: repo}
}

func (s *contractorService) CreateContractor(contractor *models.Contractor) error {
	return s.repo.Save(contractor)
}

func (s *contractorService) GetAllContractors() ([]models.Contractor, error) {
	return s.repo.FindAll()
}

func (s *contractorService) GetContractorByID(id uint) (*models.Contractor, error) {
	return s.repo.FindByID(id)
}
