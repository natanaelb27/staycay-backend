package repositories

import (
	"git.garena.com/sea-labs-id/batch-02/natanael-benediktus/final-project-backend/internal/dtos"
	"git.garena.com/sea-labs-id/batch-02/natanael-benediktus/final-project-backend/internal/models"
	"gorm.io/gorm"
)

type HouseRepository interface {
	GetHouse() ([]*dtos.HouseDetailsDTO, error)
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

func (r *houseRepository) GetHouse() ([]*dtos.HouseDetailsDTO, error) {
	var houseDetails []*dtos.HouseDetailsDTO

	result := r.db.Model(&models.House{}).
		Select("houses.id AS house_id, houses.name AS house_name, houses.price, houses.description, cities.name AS city_name, houses.max_guest, users.name AS host_name, users.email AS host_email").
		Joins("LEFT JOIN cities ON houses.city_id = cities.id").
		Joins("LEFT JOIN users ON houses.user_id = users.id").
		Find(&houseDetails)

	if result.Error != nil {
		return nil, result.Error
	}

	return houseDetails, nil
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
