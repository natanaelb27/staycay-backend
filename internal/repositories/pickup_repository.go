package repositories

import (
	"fmt"
	"strconv"

	"github.com/natanaelb27/staycay-backend/internal/dtos"
	e "github.com/natanaelb27/staycay-backend/internal/errors"
	"github.com/natanaelb27/staycay-backend/internal/helpers"
	"github.com/natanaelb27/staycay-backend/internal/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type PickupRepository interface {
	ValidateOnlyOnePickupReservation(reservationID int) error
	// ValidatePickupReservationExists(reservationID int) error
	CreatePickup(pickup *models.Pickup) (*models.Pickup, error)
	GetPickupStatusByID(pickupStatusID int) (*models.PickupStatus, error)
	GetPickupStatusByName(status string) (*models.PickupStatus, error)
	GetPickupDetailsByReservationID(reservationID int) (*dtos.GetPickupDetailsByReservationResponseDTO, error)
	UpdatePickupStatusByID(pickupStatusID int, pickupID int) (*models.Pickup, error)
	GetPickupStatuses() ([]*dtos.PickupStatusesDTO, error)
	GetPickupByID(pickupID int) (*models.Pickup, error)
	GetPickupDetails(queryParams map[string]interface{}) ([]*dtos.PickupDetailsDTO, *helpers.Pagination, error)
}

type pickupRepository struct {
	db *gorm.DB
}

type PRConfig struct {
	DB *gorm.DB
}

func NewPickupRepository(c *PRConfig) PickupRepository {
	return &pickupRepository{
		db: c.DB,
	}
}

func (r *pickupRepository) ValidateOnlyOnePickupReservation(reservationID int) error {
	result := r.db.Model(&models.Pickup{}).
		Where("reservation_id = ?", reservationID).
		Find(&models.Pickup{})

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected > 0 {
		return e.ErrDuplicatePickupReservationError
	}

	return nil
}

// func (r *pickupRepository) ValidatePickupReservationExists(reservationID int) error {
// 	result := r.db.Model(&models.Pickup{}).
// 		Where("reservation_id = ?", reservationID).
// 		Find(&models.Pickup{})

// 	if result.Error != nil {
// 		return result.Error
// 	}

// 	if result.RowsAffected == 0 {
// 		return e.ErrIdNotFoundError
// 	}

// 	return nil
// }

func (r *pickupRepository) CreatePickup(pickup *models.Pickup) (*models.Pickup, error) {
	// result := r.db.Where("reservation_id = ?", pickup.ReservationID).FirstOrCreate(&pickup)

	result := r.db.Create(&pickup)

	if result.Error != nil {
		return nil, result.Error
	}

	// if result.RowsAffected == 0 {
	// 	return nil, e.ErrDuplicatePickupReservationError
	// }

	return pickup, nil
}

func (r *pickupRepository) GetPickupStatusByID(pickupStatusID int) (*models.PickupStatus, error) {
	var pickupStatus *models.PickupStatus

	result := r.db.Where("id = ?", pickupStatusID).First(&pickupStatus)

	if result.Error != nil {
		return nil, result.Error
	}

	return pickupStatus, nil
}

func (r *pickupRepository) GetPickupStatusByName(status string) (*models.PickupStatus, error) {
	var pickupStatus *models.PickupStatus

	result := r.db.Where("status ILIKE ?", status).First(&pickupStatus)

	if result.Error != nil {
		return nil, result.Error
	}

	return pickupStatus, nil
}

func (r *pickupRepository) GetPickupDetailsByReservationID(reservationID int) (*dtos.GetPickupDetailsByReservationResponseDTO, error) {
	var houseReservation *dtos.GetPickupDetailsByReservationResponseDTO

	result := r.db.Model(&models.Reservation{}).
		Select("pickups.id AS pickup_id, reservations.id AS reservation_id, reservations.house_id AS house_id, pickup_statuses.status AS pickup_status_name, pickups.pickup_address AS pickup_address, pickups.city_id AS pickup_city_id, cities.name AS pickup_city_name, pickups.pickup_price AS pickup_price").
		Joins("LEFT JOIN pickups ON reservations.id = pickups.reservation_id").
		Joins("LEFT JOIN pickup_statuses ON pickups.pickup_status_id = pickup_statuses.id").
		Joins("LEFT JOIN cities ON pickups.city_id = cities.id").
		Where("reservations.id = ?", reservationID).
		Find(&houseReservation)

	if result.Error != nil {
		return nil, result.Error
	}

	return houseReservation, nil
}

func (r *pickupRepository) UpdatePickupStatusByID(pickupStatusID int, pickupID int) (*models.Pickup, error) {
	var pickup models.Pickup

	result := r.db.Model(&pickup).
		Clauses(clause.Returning{}).
		Where("id = ?", pickupID).
		Update("pickup_status_id", pickupStatusID)

	if result.Error != nil {
		return nil, result.Error
	}

	return &pickup, nil
}

func (r *pickupRepository) GetPickupStatuses() ([]*dtos.PickupStatusesDTO, error) {
	var pickupStatuses []*dtos.PickupStatusesDTO

	result := r.db.Model(&models.PickupStatus{}).
		Select("id AS pickup_status_id, status AS pickup_status_name").
		Find(&pickupStatuses)

	if result.Error != nil {
		return nil, result.Error
	}

	return pickupStatuses, nil

}

func (r *pickupRepository) GetPickupByID(pickupID int) (*models.Pickup, error) {
	var pickup *models.Pickup

	result := r.db.Model(&models.Pickup{}).
		Where("id = ?", pickupID).
		First(&pickup)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, e.ErrPickupNotFoundError
		}

		return nil, result.Error
	}

	return pickup, nil

}

func (r *pickupRepository) GetPickupDetails(queryParams map[string]interface{}) ([]*dtos.PickupDetailsDTO, *helpers.Pagination, error) {
	var pickupDetails []*dtos.PickupDetailsDTO

	filterPickupStatusName := ""
	if queryParams["filterPickupStatusName"] != nil {
		filterPickupStatusName = fmt.Sprintf("%v", queryParams["filterPickupStatusName"])
	}

	filterReservationCode := ""
	println(queryParams["filterReservationCode"])
	if queryParams["filterReservationCode"] != nil {
		filterReservationCode = fmt.Sprintf("%v", queryParams["filterReservationCode"])
	}
	println(filterReservationCode)

	result := r.db.Model(&models.Pickup{}).
		Select(`pickups.id AS pickup_id, pickups.reservation_id AS reservation_id, reservations.house_id AS house_id, pickups.pickup_status_id AS pickup_status_id,
		pickup_statuses.status AS pickup_status_name, pickups.pickup_address AS pickup_address, pickups.city_id AS pickup_city_id, 
		cities.name AS pickup_city_name, pickups.pickup_price AS pickup_price`).
		Joins("JOIN pickup_statuses ON pickups.pickup_status_id = pickup_statuses.id").
		Joins("JOIN reservations ON pickups.reservation_id = reservations.id").
		Joins("JOIN cities ON pickups.city_id = cities.id")

	if filterPickupStatusName != "all" {
		result.Where("pickup_statuses.status ILIKE ?", fmt.Sprintf("%%%s%%", filterPickupStatusName))
	}

	if filterReservationCode != "" {
		result.Where("reservations.reservation_code ILIKE ?", fmt.Sprintf("%%%s%%", filterReservationCode))
	}

	result.Order("pickup_id asc")

	limit := 0
	if queryParams["limit"] != nil {
		limit, _ = strconv.Atoi(queryParams["limit"].(string))
	}

	page := 0
	if queryParams["page"] != nil {
		page, _ = strconv.Atoi(queryParams["page"].(string))
	}

	pagination := helpers.Pagination{
		Limit: limit,
		Page:  page,
	}

	result.Scopes(helpers.Paginate(&pagination, result)).Find(&pickupDetails)

	if result.Error != nil {
		return nil, nil, result.Error
	}

	return pickupDetails, &pagination, nil

}
