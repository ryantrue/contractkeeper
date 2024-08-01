package repositories

import (
	"github.com/jinzhu/gorm"
	"github.com/ryantrue/contractkeeper/internal/models"
	"github.com/sirupsen/logrus"
)

type UserRepository interface {
	Save(user *models.User) error
	FindByUsername(username string) (*models.User, error)
}

type userRepository struct {
	DB     *gorm.DB
	logger *logrus.Logger
}

func NewUserRepository(db *gorm.DB, logger *logrus.Logger) UserRepository {
	return &userRepository{DB: db, logger: logger}
}

func (r *userRepository) Save(user *models.User) error {
	return r.DB.Create(user).Error
}

func (r *userRepository) FindByUsername(username string) (*models.User, error) {
	var user models.User
	err := r.DB.Where("username = ?", username).First(&user).Error
	if err != nil {
		r.logger.Errorf("Error finding user by username: %v", err)
	}
	return &user, err
}
