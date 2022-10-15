package repositories

import (
	"fmt"

	"git.garena.com/sea-labs-id/batch-02/natanael-benediktus/final-project-backend/internal/dtos"
	"git.garena.com/sea-labs-id/batch-02/natanael-benediktus/final-project-backend/internal/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type HouseRepository interface {
	GetHouse(queryParams map[string]interface{}, querySearchCity []string) ([]*dtos.HouseDetailsDTO, error)
	GetHouseByHostID(hostID int) ([]*dtos.HouseDetailsDTO, error)
	GetHouseByHouseID(houseID int) (*dtos.HouseDetailsDTO, error)
	CreateNewHouse(newHouse *models.House) (*models.House, error)
	AddHousePhoto(newHousePhoto *models.HousePhoto) (*models.HousePhoto, error)
	UpdateHouseByHouseID(houseID int, houseNewData *models.House) (*models.House, error)
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

func (r *houseRepository) GetHouse(queryParams map[string]interface{}, querySearchCity []string) ([]*dtos.HouseDetailsDTO, error) {
	var houseDetails []*dtos.HouseDetailsDTO
	if queryParams["sortBy"] == "" || queryParams["sortBy"] == nil {
		queryParams["sortBy"] = "t.house_name"
	} else {
		queryParams["sortBy"] = fmt.Sprintf("t.house_%s", queryParams["sortBy"])
	}

	qFilterCityName := ""
	print(len(querySearchCity))
	if len(querySearchCity) > 0 {
		qFilterCityName += "AND t.house_city_name IN ("
		for i, city := range querySearchCity {
			if i == len(querySearchCity)-1 {
				qFilterCityName += "'" + city + "')"
				break
			}
			qFilterCityName += "'" + city + "', "
		}
	}

	qFilterHouseName := ""
	if queryParams["searchName"] != nil {
		qFilterHouseName = fmt.Sprintf("%v", queryParams["searchName"])
	}

	qSortBy := fmt.Sprintf("ORDER BY %s", queryParams["sortBy"])

	qSort := ""
	if queryParams["sort"] != nil {
		qSort = fmt.Sprintf("%v", queryParams["sort"])

	}
	query := fmt.Sprintf("SELECT * FROM (?) t WHERE t.house_name ILIKE '%s' %s %s %s", fmt.Sprintf("%%%s%%", qFilterHouseName), qFilterCityName, qSortBy, qSort)

	result := r.db.Raw(query, r.db.Model(&models.House{}).
		Select("houses.id AS house_id, houses.name AS house_name, houses.price AS house_price, houses.description AS house_description, cities.name AS house_city_name, houses.max_guest AS house_max_guest, houses.status AS house_status, ARRAY( SELECT photo FROM house_photos WHERE house_id = houses.id) AS house_photo, users.id AS host_id, users.name AS host_name, users.email AS host_email").
		Joins("LEFT JOIN cities ON houses.city_id = cities.id").
		Joins("LEFT JOIN users ON houses.user_id = users.id")).
		Where("houses.status = ?", "active").
		Scan(&houseDetails)

	// result := r.db.Model(&models.House{}).
	// 	Select("houses.id AS house_id, houses.name AS house_name, houses.price AS house_price, houses.description AS house_description, cities.name AS house_city_name, houses.max_guest AS house_max_guest, houses.status AS house_status, users.id AS host_id, users.name AS host_name, users.email AS host_email").
	// 	Joins("LEFT JOIN cities ON houses.city_id = cities.id").
	// 	Joins("LEFT JOIN users ON houses.user_id = users.id").
	// 	Find(&houseDetails)

	if result.Error != nil {
		return nil, result.Error
	}

	return houseDetails, nil
}

func (r *houseRepository) GetHouseByHostID(hostID int) ([]*dtos.HouseDetailsDTO, error) {
	var houseDetails []*dtos.HouseDetailsDTO

	result := r.db.Model(&models.House{}).
		Select("houses.id AS house_id, houses.name AS house_name, houses.price AS house_price, houses.description AS house_description, cities.name AS house_city_name, houses.max_guest AS house_max_guest, houses.status AS house_status, ARRAY( SELECT photo FROM house_photos WHERE house_id = houses.id) AS house_photo, users.id AS host_id, users.name AS host_name, users.email AS host_email").
		Joins("LEFT JOIN cities ON houses.city_id = cities.id").
		Joins("LEFT JOIN users ON houses.user_id = users.id").
		Where("houses.user_id = ?", hostID).
		Find(&houseDetails)

	if result.Error != nil {
		return nil, result.Error
	}

	return houseDetails, nil
}

func (r *houseRepository) GetHouseByHouseID(houseID int) (*dtos.HouseDetailsDTO, error) {
	var houseDetails *dtos.HouseDetailsDTO

	result := r.db.Model(&models.House{}).
		Select("houses.id AS house_id, houses.name AS house_name, houses.price AS house_price, houses.description AS house_description, cities.name AS house_city_name, houses.max_guest AS house_max_guest, houses.status AS house_status, users.id AS host_id, users.name AS host_name, users.email AS host_email").
		Joins("LEFT JOIN cities ON houses.city_id = cities.id").
		Joins("LEFT JOIN users ON houses.user_id = users.id").
		Where("houses.id = ?", houseID).
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

func (r *houseRepository) UpdateHouseByHouseID(houseID int, houseNewData *models.House) (*models.House, error) {
	var house models.House

	result := r.db.
		Model(&house).
		Clauses(clause.Returning{}).
		Where("id = ?", houseID).
		Updates(houseNewData)

	if result.Error != nil {
		return nil, result.Error
	}

	return &house, nil
}
