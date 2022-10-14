package repositories

import (
	"git.garena.com/sea-labs-id/batch-02/natanael-benediktus/final-project-backend/internal/models"
	"gorm.io/gorm"
)

type HouseRepository interface {
	GetHouse() ([]*models.House, error)
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
