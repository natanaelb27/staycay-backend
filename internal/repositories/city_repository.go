package repositories

import (
	"github.com/natanaelb27/staycay-backend/internal/dtos"
	"github.com/natanaelb27/staycay-backend/internal/models"
	"gorm.io/gorm"
)

type CityRepository interface {
	GetCity() ([]*dtos.CitiesDTO, error)
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

func (r *cityRepository) GetCity() ([]*dtos.CitiesDTO, error) {
	var city []*dtos.CitiesDTO

	result := r.db.
		Model(&models.City{}).
		Select("id AS city_id, name AS city_name").
		Find(&city)

	if result.Error != nil {
		return nil, result.Error
	}

	return city, nil
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
