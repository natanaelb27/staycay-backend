package repositories

import (
	"fmt"
	"time"

	"git.garena.com/sea-labs-id/batch-02/natanael-benediktus/final-project-backend/internal/dtos"
	e "git.garena.com/sea-labs-id/batch-02/natanael-benediktus/final-project-backend/internal/errors"
	"git.garena.com/sea-labs-id/batch-02/natanael-benediktus/final-project-backend/internal/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type HouseRepository interface {
	GetHouse(queryParams map[string]interface{}) ([]*dtos.HouseDetailsDTO, error)
	GetHouseByHostID(hostID int) ([]*dtos.HouseDetailsDTO, error)
	GetHousePhotoByHouseID(houseID int) ([]*dtos.HousePhotosDTO, error)
	GetUnscopedHousePhotoByHouseID(houseID int) ([]*dtos.HousePhotosDTO, error)
	GetHouseByHouseID(houseID int) (*dtos.HouseDetailsDTO, error)
	GetUnscopedHouseByHouseID(houseID int) (*dtos.HouseDetailsDTO, error)
	GetHouseByHousePhotoID(housePhotoID int) (*dtos.HouseDetailsDTO, error)
	GetHouseStatusByName(status string) (*models.HouseStatus, error)
	GetHouseStatus() ([]*dtos.HouseStatusDTO, error)
	GetHouseStatusByID(houseStatusID int) (*models.HouseStatus, error)
	CreateNewHouse(newHouse *models.House) (*models.House, error)
	AddHousePhoto(newHousePhoto *models.HousePhoto) (*models.HousePhoto, error)
	UpdateHouseByHouseID(houseID int, houseNewData *models.House) (*models.House, error)
	DeleteHousePhotoByHouseID(houseID int) error
	DeleteHouseByID(houseID int) (*models.House, error)
	ValidateTotalHousePhoto(houseID int) (int64, error)
	DeleteHousePhotoByID(housePhotoID int) (*models.HousePhoto, error)
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

func (r *houseRepository) GetHouse(queryParams map[string]interface{}) ([]*dtos.HouseDetailsDTO, error) {
	var houseDetails []*dtos.HouseDetailsDTO
	if queryParams["sortBy"] == "" || queryParams["sortBy"] == nil {
		queryParams["sortBy"] = "t.house_name"
	} else {
		queryParams["sortBy"] = fmt.Sprintf("t.house_%s", queryParams["sortBy"])
	}

	// if len(querySearchCity) > 0 {
	// 	qFilterCityName += "AND t.house_city_name IN ("
	// 	for i, city := range querySearchCity {
	// 		if i == len(querySearchCity)-1 {
	// 			qFilterCityName += "'" + city + "')"
	// 			break
	// 		}
	// 		qFilterCityName += "'" + city + "', "
	// 	}
	// }

	qFilterHouseName := ""
	if queryParams["searchName"] != nil {
		qFilterHouseName = fmt.Sprintf("%v", queryParams["searchName"])
	}

	qFilterCityID := ""
	if queryParams["searchCity"] != nil && queryParams["searchCity"] != "0" {
		qFilterCityID = fmt.Sprintf("AND t.house_city_id = %v", queryParams["searchCity"])
	}

	qFilterStatus := ""
	if queryParams["searchStatus"] != nil {
		qFilterStatus = fmt.Sprintf("AND t.house_status ILIKE '%v'", queryParams["searchStatus"])
	}

	format := "2006-01-02"

	qFilterCheckInDate := ""
	if queryParams["searchCheckInDate"] != nil && queryParams["searchCheckInDate"] != "" {
		qFilterCheckInDate = fmt.Sprintf("%v", queryParams["searchCheckInDate"])
	} else {
		// if check in date is nil but check out date is not nil then check in date = check in date - 1
		if queryParams["searchCheckOutDate"] != nil && queryParams["searchCheckOutDate"] != "" {
			checkOutDate, _ := time.Parse(format, fmt.Sprintf("%v", queryParams["searchCheckOutDate"]))
			checkInDate := checkOutDate.AddDate(0, 0, -1)
			qFilterCheckInDate = checkInDate.Format(format)
		}
	}

	qFilterCheckOutDate := ""
	if queryParams["searchCheckOutDate"] != nil && queryParams["searchCheckOutDate"] != "" {
		qFilterCheckOutDate = fmt.Sprintf("%v", queryParams["searchCheckOutDate"])
	} else {
		// if check out date is nil but check in date is not nil then check out date = check in date + 1
		if queryParams["searchCheckInDate"] != nil && queryParams["searchCheckInDate"] != "" {
			checkInDate, _ := time.Parse(format, fmt.Sprintf("%v", queryParams["searchCheckInDate"]))
			checkOutDate := checkInDate.AddDate(0, 0, 1)
			qFilterCheckOutDate = checkOutDate.Format(format)
		}
	}

	qFilterReservation := ""
	if (queryParams["searchCheckInDate"] != nil && queryParams["searchCheckInDate"] != "") || (queryParams["searchCheckOutDate"] != nil && queryParams["searchCheckOutDate"] != "") {
		qFilterReservation = fmt.Sprintf(`AND t.house_id NOT IN (SELECT DISTINCT house_id FROM reservations WHERE '%s'::timestamp::date BETWEEN check_in_date::timestamp::date AND check_out_date::timestamp::date 
			OR '%s'::timestamp::date BETWEEN check_in_date::timestamp::date AND check_out_date::timestamp::date 
			OR check_in_date::timestamp::date BETWEEN '%s'::timestamp::date AND '%s'::timestamp::date
			OR check_out_date::timestamp::date BETWEEN '%s'::timestamp::date AND '%s'::timestamp::date)`, qFilterCheckInDate, qFilterCheckOutDate, qFilterCheckInDate, qFilterCheckOutDate, qFilterCheckInDate, qFilterCheckOutDate)
	}

	qSortBy := fmt.Sprintf("ORDER BY %s", queryParams["sortBy"])

	qSort := ""
	if queryParams["sort"] != nil {
		qSort = fmt.Sprintf("%v", queryParams["sort"])
	}

	query := fmt.Sprintf("SELECT * FROM (?) t WHERE t.house_name ILIKE '%s' %s %s %s %s %s", fmt.Sprintf("%%%s%%", qFilterHouseName), qFilterCityID, qFilterStatus, qFilterReservation, qSortBy, qSort)

	result := r.db.Raw(query, r.db.Model(&models.House{}).
		Select("houses.id AS house_id, houses.name AS house_name, houses.price AS house_price, houses.description AS house_description, cities.id AS house_city_id, cities.name AS house_city_name, houses.max_guest AS house_max_guest, house_statuses.status AS house_status, users.id AS host_id, users.name AS host_name, users.email AS host_email").
		Joins("JOIN cities ON houses.city_id = cities.id").
		Joins("JOIN users ON houses.user_id = users.id").
		Joins("JOIN house_statuses ON houses.house_status_id = house_statuses.id")).
		Scan(&houseDetails)

	if result.Error != nil {
		return nil, result.Error
	}

	return houseDetails, nil
}

func (r *houseRepository) GetHouseByHostID(hostID int) ([]*dtos.HouseDetailsDTO, error) {
	var houseDetails []*dtos.HouseDetailsDTO

	result := r.db.Model(&models.House{}).
		Select("houses.id AS house_id, houses.name AS house_name, houses.price AS house_price, houses.description AS house_description, cities.id AS house_city_id, cities.name AS house_city_name, houses.max_guest AS house_max_guest, house_statuses.status AS house_status, users.id AS host_id, users.name AS host_name, users.email AS host_email").
		Joins("JOIN cities ON houses.city_id = cities.id").
		Joins("JOIN users ON houses.user_id = users.id").
		Joins("JOIN house_statuses ON houses.house_status_id = house_statuses.id").
		Where("houses.user_id = ?", hostID).
		Find(&houseDetails)

	if result.Error != nil {
		return nil, result.Error
	}

	return houseDetails, nil
}

func (r *houseRepository) GetHousePhotoByHouseID(houseID int) ([]*dtos.HousePhotosDTO, error) {
	var housePhotos []*dtos.HousePhotosDTO

	result := r.db.Model(&models.HousePhoto{}).
		Select("id AS photo_id, photo").
		Where("house_id = ?", houseID).
		Find(&housePhotos)

	if result.Error != nil {
		return nil, result.Error
	}

	return housePhotos, nil
}

func (r *houseRepository) GetUnscopedHousePhotoByHouseID(houseID int) ([]*dtos.HousePhotosDTO, error) {
	var housePhotos []*dtos.HousePhotosDTO

	result := r.db.Model(&models.HousePhoto{}).
		Unscoped().
		Select("id AS photo_id, photo").
		Where("house_id = ?", houseID).
		Find(&housePhotos)

	if result.Error != nil {
		return nil, result.Error
	}

	return housePhotos, nil
}

func (r *houseRepository) GetHouseByHouseID(houseID int) (*dtos.HouseDetailsDTO, error) {
	var houseDetails *dtos.HouseDetailsDTO

	result := r.db.Model(&models.House{}).
		Select("houses.id AS house_id, houses.name AS house_name, houses.price AS house_price, houses.description AS house_description, cities.id AS house_city_id, cities.name AS house_city_name, houses.max_guest AS house_max_guest, house_statuses.status AS house_status, users.id AS host_id, users.name AS host_name, users.email AS host_email").
		Joins("JOIN cities ON houses.city_id = cities.id").
		Joins("JOIN users ON houses.user_id = users.id").
		Joins("JOIN house_statuses ON houses.house_status_id = house_statuses.id").
		Where("houses.id = ?", houseID).
		Find(&houseDetails)

	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, e.ErrIdNotFoundError
	}

	return houseDetails, nil
}

func (r *houseRepository) GetUnscopedHouseByHouseID(houseID int) (*dtos.HouseDetailsDTO, error) {
	var houseDetails *dtos.HouseDetailsDTO

	result := r.db.Model(&models.House{}).
		Unscoped().
		Select("houses.id AS house_id, houses.name AS house_name, houses.price AS house_price, houses.description AS house_description, cities.id AS house_city_id, cities.name AS house_city_name, houses.max_guest AS house_max_guest, house_statuses.status AS house_status, users.id AS host_id, users.name AS host_name, users.email AS host_email").
		Joins("JOIN cities ON houses.city_id = cities.id").
		Joins("JOIN users ON houses.user_id = users.id").
		Joins("JOIN house_statuses ON houses.house_status_id = house_statuses.id").
		Where("houses.id = ?", houseID).
		Find(&houseDetails)

	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, e.ErrIdNotFoundError
	}

	return houseDetails, nil
}

func (r *houseRepository) GetHouseStatus() ([]*dtos.HouseStatusDTO, error) {
	var houseStatus []*dtos.HouseStatusDTO

	result := r.db.
		Model(&models.HouseStatus{}).
		Select("id AS house_status_id, status AS house_status_name").
		Find(&houseStatus)

	if result.Error != nil {
		return nil, result.Error
	}

	return houseStatus, nil
}

func (r *houseRepository) GetHouseStatusByID(houseStatusID int) (*models.HouseStatus, error) {
	var houseStatus *models.HouseStatus

	result := r.db.Model(&models.HouseStatus{}).
		Where("id = ?", houseStatusID).
		First(&houseStatus)

	if result.Error != nil {
		return nil, result.Error
	}

	return houseStatus, nil
}

func (r *houseRepository) GetHouseStatusByName(status string) (*models.HouseStatus, error) {
	var houseStatus *models.HouseStatus

	result := r.db.Model(&models.HouseStatus{}).
		Where("status ILIKE ?", status).
		First(&houseStatus)

	if result.Error != nil {
		return nil, result.Error
	}

	return houseStatus, nil
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
		Updates(&houseNewData)

	if result.Error != nil {
		return nil, result.Error
	}

	return &house, nil
}

func (r *houseRepository) GetHouseByHousePhotoID(housePhotoID int) (*dtos.HouseDetailsDTO, error) {
	var houseDetails *dtos.HouseDetailsDTO

	result := r.db.Model(&models.HousePhoto{}).
		Select("houses.id AS house_id, houses.name AS house_name, houses.price AS house_price, houses.description AS house_description, cities.id AS house_city_id, cities.name AS house_city_name, houses.max_guest AS house_max_guest, house_statuses.status AS house_status, users.id AS host_id, users.name AS host_name, users.email AS host_email").
		Joins("JOIN houses ON house_photos.house_id = houses.id").
		Joins("JOIN cities ON houses.city_id = cities.id").
		Joins("JOIN users ON houses.user_id = users.id").
		Joins("JOIN house_statuses ON houses.house_status_id = house_statuses.id").
		Where("house_photos.id = ?", housePhotoID).
		Find(&houseDetails)

	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, e.ErrIdNotFoundError
	}

	return houseDetails, nil
}

func (r *houseRepository) DeleteHousePhotoByHouseID(houseID int) error {
	var housePhoto *models.HousePhoto

	result := r.db.
		Where("house_id = ?", houseID).
		Delete(&housePhoto)

	if result.Error != nil {
		return result.Error
	}

	return nil

}

func (r *houseRepository) DeleteHouseByID(houseID int) (*models.House, error) {
	var house *models.House

	result := r.db.
		Clauses(clause.Returning{}).
		Where("id = ?", houseID).
		Delete(&house)

	if result.Error != nil {
		return nil, result.Error
	}

	return house, nil
}

func (r *houseRepository) ValidateTotalHousePhoto(houseID int) (int64, error) {
	var totalRows int64

	result := r.db.
		Model(&models.HousePhoto{}).
		Where("house_id = ?", houseID).
		Count(&totalRows)

	if result.Error != nil {
		return 0, result.Error
	}

	return totalRows, nil
}

func (r *houseRepository) DeleteHousePhotoByID(housePhotoID int) (*models.HousePhoto, error) {
	var housePhoto *models.HousePhoto

	result := r.db.
		Clauses(clause.Returning{}).
		Where("id = ?", housePhotoID).
		Delete(&housePhoto)

	if result.Error != nil {
		return nil, result.Error
	}

	return housePhoto, nil
}

// func (r *houseRepository) GetHouseReservationsByHouseID(houseID int) (*dtos.HouseReservationsDTO, error) {
// 	var houseReservation *dtos.HouseReservationsDTO

// 	result := r.db.Model(&models.House{}).
// 		Select("houses.id AS house_id, reservations.id AS reservation_id, reservations.reservation_code AS reservation_code, reservations.check_in_date AS check_in_date, reservations.check_out_date AS check_out_date, reservations.total_price AS total_price").
// 		Joins("JOIN reservations ON houses.id = reservations.house_id").
// 		Where("house.id = ?", houseID).
// 		Find(&houseReservation)

// 	if result.Error != nil {
// 		return nil, result.Error
// 	}

// 	return houseReservation, nil
// }

// func (r *houseRepository) GetUpcomingHouseReservationsByHouseID(houseID int) (*dtos.HouseReservationsDTO, error) {
// 	var houseReservation *dtos.HouseReservationsDTO

// 	result := r.db.Model(&models.House{}).
// 		Select("houses.id AS house_id, reservations.id AS reservation_id, reservations.reservation_code AS reservation_code, reservations.check_in_date AS check_in_date, reservations.check_out_date AS check_out_date, reservations.total_price AS total_price").
// 		Joins("JOIN reservations ON houses.id = reservations.house_id").
// 		Where("houses.id = ?", houseID).
// 		Where("reservations.check_in_date > ?", time.Now()).
// 		Find(&houseReservation)

// 	if result.Error != nil {
// 		return nil, result.Error
// 	}

// 	if result.RowsAffected > 0 {
// 		return nil, e.ErrDeleteUpcomingReservationError
// 	}

// 	return houseReservation, nil
// }
