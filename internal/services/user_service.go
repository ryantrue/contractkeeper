package services

import (
	"errors"

	"github.com/ryantrue/contractkeeper/internal/models"
	"github.com/ryantrue/contractkeeper/internal/repositories"
)

type UserService interface {
	CreateUser(user *models.User) error
	Authenticate(username, password string) (*models.User, error)
}

type userService struct {
	repo repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) CreateUser(user *models.User) error {
	return s.repo.Save(user)
}

func (s *userService) Authenticate(username, password string) (*models.User, error) {
	user, err := s.repo.FindByUsername(username)
	if err != nil {
		return nil, err
	}

	if user.Password != password {
		return nil, errors.New("invalid credentials")
	}

	return user, nil
}
