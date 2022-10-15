package repositories

import (
	"time"

	"git.garena.com/sea-labs-id/batch-02/natanael-benediktus/final-project-backend/internal/dtos"
	e "git.garena.com/sea-labs-id/batch-02/natanael-benediktus/final-project-backend/internal/errors"
	"git.garena.com/sea-labs-id/batch-02/natanael-benediktus/final-project-backend/internal/models"
	"gorm.io/gorm"
)

type ReservationRepository interface {
	GetUpcomingHouseReservationsByHouseID(houseID int) (*dtos.HouseReservationsDTO, error)
	ValidateReservation(request *dtos.AddReservationRequestDTO) error
	GetLastReservationID() (int, error)
	CreateReservation(userID int, request *models.Reservation) (*models.Reservation, error)
}

type reservationRepository struct {
	db *gorm.DB
}

type RRConfig struct {
	DB *gorm.DB
}

func NewReservationRepository(c *RRConfig) ReservationRepository {
	return &reservationRepository{
		db: c.DB,
	}
}

func (r *reservationRepository) GetUpcomingHouseReservationsByHouseID(houseID int) (*dtos.HouseReservationsDTO, error) {
	var houseReservation *dtos.HouseReservationsDTO

	result := r.db.Model(&models.House{}).
		Select("houses.id AS house_id, reservations.id AS reservation_id, reservations.reservation_code AS reservation_code, reservations.check_in_date AS check_in_date, reservations.check_out_date AS check_out_date, reservations.total_price AS total_price").
		Joins("JOIN reservations ON houses.id = reservations.house_id").
		Where("houses.id = ?", houseID).
		Where("reservations.check_in_date > ?", time.Now()).
		Find(&houseReservation)

	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected > 0 {
		return nil, e.ErrDeleteUpcomingReservationError
	}

	return houseReservation, nil
}

func (r *reservationRepository) ValidateReservation(request *dtos.AddReservationRequestDTO) error {
	result := r.db.Model(&models.Reservation{}).
		Where("house_id = ?", request.HouseID).
		Where("check_in_date BETWEEN ? AND ?", request.CheckInDate, request.CheckOutDate).
		Or("check_out_date BETWEEN ? AND ?", request.CheckInDate, request.CheckOutDate).
		Find(&models.Reservation{})

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected > 0 {
		return e.ErrNotValidReservationDateConflict
	}

	return nil
}

func (r *reservationRepository) GetLastReservationID() (int, error) {
	var lastReservationID int
	result := r.db.Model(&models.Reservation{}).
		Select("CASE WHEN max(id) IS NULL THEN 0 ELSE max(id) END").
		Scan(&lastReservationID)

	if result.Error != nil {
		return 0, result.Error
	}

	return lastReservationID, nil
}

func (r *reservationRepository) CreateReservation(userID int, newReservation *models.Reservation) (*models.Reservation, error) {
	result := r.db.Create(&newReservation)

	if result.Error != nil {
		return nil, result.Error
	}

	return newReservation, nil
}
