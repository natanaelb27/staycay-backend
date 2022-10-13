package repositories

import (
	"git.garena.com/sea-labs-id/batch-02/natanael-benediktus/final-project-backend/internal/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	CreateNewUser(user *models.User) (*models.User, error)
	GetUserByEmail(email string) (*models.User, error)
}

type userRepository struct {
	db *gorm.DB
}

type URConfig struct {
	DB *gorm.DB
}

func NewUserRepository(c *URConfig) UserRepository {
	return &userRepository{
		db: c.DB,
	}
}

func (r *userRepository) CreateNewUser(user *models.User) (*models.User, error) {
	result := r.db.Create(&user)
	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}

func (r *userRepository) GetUserByEmail(email string) (*models.User, error) {
	var user *models.User

	result := r.db.
		Where("email = ?", email).
		First(&user)

	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}
