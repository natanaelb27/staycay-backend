package repositories

import (
	e "git.garena.com/sea-labs-id/batch-02/natanael-benediktus/final-project-backend/internal/errors"
	"git.garena.com/sea-labs-id/batch-02/natanael-benediktus/final-project-backend/internal/models"
	"gorm.io/gorm"
)

type PickupRepository interface {
	CreatePickup(pickup *models.Pickup) (*models.Pickup, error)
	GetPickupStatusByName(status string) (*models.PickupStatus, error)
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

func (r *pickupRepository) CreatePickup(pickup *models.Pickup) (*models.Pickup, error) {
	result := r.db.Where("reservation_id = ?", pickup.ReservationID).FirstOrCreate(&pickup)

	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, e.ErrDuplicatePickupReservationError
	}

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
