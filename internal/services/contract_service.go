package services

import (
	"github.com/ryantrue/contractkeeper/internal/models"
	"github.com/ryantrue/contractkeeper/internal/repositories"
)

type ContractService interface {
	CreateContract(contract *models.Contract) error
	GetAllContracts() ([]models.Contract, error)
	GetContractByID(id uint) (*models.Contract, error)
}

type contractService struct {
	repo repositories.ContractRepository
}

func NewContractService(repo repositories.ContractRepository) ContractService {
	return &contractService{repo: repo}
}

func (s *contractService) CreateContract(contract *models.Contract) error {
	return s.repo.Save(contract)
}

func (s *contractService) GetAllContracts() ([]models.Contract, error) {
	return s.repo.FindAll()
}

func (s *contractService) GetContractByID(id uint) (*models.Contract, error) {
	return s.repo.FindByID(id)
}
