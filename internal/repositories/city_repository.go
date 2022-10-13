package repositories

import (
	"git.garena.com/sea-labs-id/batch-02/natanael-benediktus/final-project-backend/internal/models"
	"gorm.io/gorm"
)

type CityRepository interface {
	GetCityIDByName(cityName string) (*models.City, error)
}

type cityRepository struct {
	db *gorm.DB
}

type CRConfig struct {
	DB *gorm.DB
}

func NewCityRepository(c *CRConfig) CityRepository {
	return &cityRepository{
		db: c.DB,
	}
}

func (r *cityRepository) GetCityIDByName(cityName string) (*models.City, error) {
	var city *models.City

	result := r.db.
		Where("name = ?", cityName).
		First(&city)

	if result.Error != nil {
		return nil, result.Error
	}

	return city, nil
}
