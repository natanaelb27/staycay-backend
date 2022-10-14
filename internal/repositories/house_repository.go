package repositories

import (
	"git.garena.com/sea-labs-id/batch-02/natanael-benediktus/final-project-backend/internal/models"
	"gorm.io/gorm"
)

type HouseRepository interface {
	GetHouse() ([]*models.House, error)
	CreateNewHouse(newHouse *models.House) (*models.House, error)
	AddHousePhoto(newHousePhoto *models.HousePhoto) (*models.HousePhoto, error)
}

type houseRepository struct {
	db *gorm.DB
}

type HRConfig struct {
	DB *gorm.DB
}

func NewHouseRepository(c *HRConfig) HouseRepository {
	return &houseRepository{
		db: c.DB,
	}
}

func (r *houseRepository) GetHouse() ([]*models.House, error) {
	var house []*models.House

	result := r.db.Find(&house)

	if result.Error != nil {
		return nil, result.Error
	}

	return house, nil
}

func (r *houseRepository) CreateNewHouse(newHouse *models.House) (*models.House, error) {
	result := r.db.Create(&newHouse)

	if result.Error != nil {
		return nil, result.Error
	}

	return newHouse, nil
}

func (r *houseRepository) AddHousePhoto(newHousePhoto *models.HousePhoto) (*models.HousePhoto, error) {
	result := r.db.Create(&newHousePhoto)

	if result.Error != nil {
		return nil, result.Error
	}

	return newHousePhoto, nil
}
