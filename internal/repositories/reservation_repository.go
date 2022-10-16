package repositories

import (
	"git.garena.com/sea-labs-id/batch-02/natanael-benediktus/final-project-backend/internal/dtos"
	e "git.garena.com/sea-labs-id/batch-02/natanael-benediktus/final-project-backend/internal/errors"
	"git.garena.com/sea-labs-id/batch-02/natanael-benediktus/final-project-backend/internal/models"
	"gorm.io/gorm"
)

type ReservationRepository interface {
	GetHouseReservationsByUserID(userID int, filterBy string) ([]*dtos.HouseReservationsDTO, error)
	GetHouseReservationsByHouseID(houseID int, filterBy string) ([]*dtos.HouseReservationsDTO, error)
	GetHouseReservationsByID(reservationID int) (*dtos.HouseReservationsDTO, error)
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

func (r *reservationRepository) GetHouseReservationsByUserID(userID int, filterBy string) ([]*dtos.HouseReservationsDTO, error) {
	var houseReservation []*dtos.HouseReservationsDTO

	result := r.db.Model(&models.Reservation{}).
		Select(`houses.id AS house_id, houses.city_id AS house_city_id, reservations.user_id AS user_id, reservations.id AS reservation_id, reservations.reservation_code AS reservation_code, reservations.check_in_date AS check_in_date, reservations.check_out_date AS check_out_date, reservations.total_price AS total_price,
		(CASE WHEN NOW()::timestamp::date = reservations.check_in_date::timestamp::date THEN 'today'
		WHEN NOW()::timestamp::date BETWEEN reservations.check_in_date::timestamp::date AND reservations.check_out_date::timestamp::date THEN 'ongoing'
		WHEN reservations.check_in_date::timestamp::date > NOW() THEN 'upcoming'
		WHEN reservations.check_out_date::timestamp::date < NOW() THEN 'past' END) AS reservation_status`).
		Joins("JOIN houses ON reservations.house_id = houses.id").
		Where("reservations.user_id = ?", userID)

	// result := r.db.Model(&models.House{}).
	// 	Select("houses.id AS house_id, reservations.user_id AS user_id, reservations.id AS reservation_id, reservations.reservation_code AS reservation_code, reservations.check_in_date AS check_in_date, reservations.check_out_date AS check_out_date, reservations.total_price AS total_price").
	// 	Joins("JOIN reservations ON houses.id = reservations.house_id").
	// 	Where("reservations.user_id = ?", userID)

	if filterBy == "today" {
		result = result.Where("NOW()::timestamp::date = reservations.check_in_date::timestamp::date")
	}

	if filterBy == "ongoing" {
		result = result.Where("NOW()::timestamp::date BETWEEN reservations.check_in_date::timestamp::date AND reservations.check_out_date::timestamp::date")
	}

	if filterBy == "upcoming" {
		result = result.Where("reservations.check_in_date::timestamp::date > NOW()")
	}

	if filterBy == "past" {
		result = result.Where("reservations.check_out_date::timestamp::date < NOW()")
	}

	result.Find(&houseReservation)

	if result.Error != nil {
		return nil, result.Error
	}

	return houseReservation, nil
}

func (r *reservationRepository) GetHouseReservationsByHouseID(houseID int, filterBy string) ([]*dtos.HouseReservationsDTO, error) {
	var houseReservation []*dtos.HouseReservationsDTO

	result := r.db.Model(&models.Reservation{}).
		Select(`houses.id AS house_id, houses.city_id AS house_city_id, reservations.user_id AS user_id, reservations.id AS reservation_id, reservations.reservation_code AS reservation_code, reservations.check_in_date AS check_in_date, reservations.check_out_date AS check_out_date, reservations.total_price AS total_price,
		(CASE WHEN NOW()::timestamp::date = reservations.check_in_date::timestamp::date THEN 'today'
		WHEN NOW()::timestamp::date BETWEEN reservations.check_in_date::timestamp::date AND reservations.check_out_date::timestamp::date THEN 'ongoing'
		WHEN reservations.check_in_date::timestamp::date > NOW() THEN 'upcoming'
		WHEN reservations.check_out_date::timestamp::date < NOW() THEN 'past' END) AS reservation_status`).
		Joins("JOIN houses ON reservations.house_id = houses.id").
		Where("reservations.house_id = ?", houseID)

	if filterBy == "today" {
		result = result.Where("NOW()::timestamp::date = reservations.check_in_date::timestamp::date")
	}

	if filterBy == "ongoing" {
		result = result.Where("NOW()::timestamp::date BETWEEN reservations.check_in_date::timestamp::date AND reservations.check_out_date::timestamp::date")
	}

	if filterBy == "upcoming" {
		result = result.Where("reservations.check_in_date::timestamp::date > NOW()")
	}

	if filterBy == "past" {
		result = result.Where("reservations.check_out_date::timestamp::date < NOW()")
	}

	result.Find(&houseReservation)

	if result.Error != nil {
		return nil, result.Error
	}

	return houseReservation, nil
}

func (r *reservationRepository) GetHouseReservationsByID(reservationID int) (*dtos.HouseReservationsDTO, error) {
	var houseReservation *dtos.HouseReservationsDTO

	result := r.db.Model(&models.Reservation{}).
		Select(`houses.id AS house_id, houses.city_id AS house_city_id, reservations.user_id AS user_id, reservations.id AS reservation_id, reservations.reservation_code AS reservation_code, reservations.check_in_date AS check_in_date, reservations.check_out_date AS check_out_date, reservations.total_price AS total_price,
		(CASE WHEN NOW()::timestamp::date = reservations.check_in_date::timestamp::date THEN 'today'
		WHEN NOW()::timestamp::date BETWEEN reservations.check_in_date::timestamp::date AND reservations.check_out_date::timestamp::date THEN 'ongoing'
		WHEN reservations.check_in_date::timestamp::date > NOW() THEN 'upcoming'
		WHEN reservations.check_out_date::timestamp::date < NOW() THEN 'past' END) AS reservation_status`).
		Joins("JOIN houses ON reservations.house_id = houses.id").
		Where("reservations.id = ?", reservationID).
		Find(&houseReservation)

	if result.Error != nil {
		return nil, result.Error
	}

	return houseReservation, nil
}

func (r *reservationRepository) ValidateReservation(request *dtos.AddReservationRequestDTO) error {
	result := r.db.Model(&models.Reservation{}).
		Where("house_id = ?", request.HouseID).
		Where("? BETWEEN check_in_date::timestamp::date AND check_out_date::timestamp::date", request.CheckInDate).
		Or("? BETWEEN check_in_date::timestamp::date AND check_out_date::timestamp::date", request.CheckOutDate).
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

func (r *reservationRepository) ValidateTodaysReservation(reservationID int) error {
	result := r.db.Model(&models.Reservation{}).
		Where("id = ?", reservationID).
		Where("check_in_date::timestamp::date = NOW()::timestamp::date").
		Find(&models.Reservation{})

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected > 0 {
		return e.ErrNotTodaysReservationError
	}

	return nil
}
