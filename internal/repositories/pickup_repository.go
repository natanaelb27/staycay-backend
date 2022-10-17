package repositories

import (
	"git.garena.com/sea-labs-id/batch-02/natanael-benediktus/final-project-backend/internal/dtos"
	e "git.garena.com/sea-labs-id/batch-02/natanael-benediktus/final-project-backend/internal/errors"
	"git.garena.com/sea-labs-id/batch-02/natanael-benediktus/final-project-backend/internal/models"
	"gorm.io/gorm"
)

type PickupRepository interface {
	ValidateOnlyOnePickupReservation(reservationID int) error
	ValidatePickupReservationExists(reservationID int) error
	CreatePickup(pickup *models.Pickup) (*models.Pickup, error)
	GetPickupStatusByName(status string) (*models.PickupStatus, error)
	GetReservationPickupStatus(reservationID int) (*dtos.GetPickupStatusByReservationResponseDTO, error)
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

func (r *pickupRepository) ValidatePickupReservationExists(reservationID int) error {
	result := r.db.Model(&models.Pickup{}).
		Where("reservation_id = ?", reservationID).
		Find(&models.Pickup{})

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return e.ErrIdNotFoundError
	}

	return nil
}

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

func (r *pickupRepository) GetPickupStatusByName(status string) (*models.PickupStatus, error) {
	var pickupStatus *models.PickupStatus

	result := r.db.Where("status ILIKE ?", status).First(&pickupStatus)

	if result.Error != nil {
		return nil, result.Error
	}

	return pickupStatus, nil
}

func (r *pickupRepository) GetReservationPickupStatus(reservationID int) (*dtos.GetPickupStatusByReservationResponseDTO, error) {
	var houseReservationStatus *dtos.GetPickupStatusByReservationResponseDTO

	result := r.db.Model(&models.Pickup{}).
		Select("pickups.reservation_id AS reservation_id, pickup_statuses.status AS pickup_status_name").
		Joins("JOIN pickup_statuses ON pickups.pickup_status_id = pickup_statuses.id").
		Where("reservation_id = ?", reservationID).
		First(&houseReservationStatus)

	if result.Error != nil {
		return nil, result.Error
	}

	return houseReservationStatus, nil
}
