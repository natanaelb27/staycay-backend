package dtos

import (
	"time"

	"git.garena.com/sea-labs-id/batch-02/natanael-benediktus/final-project-backend/internal/helpers"
)

type HouseReservationsDTO struct {
	HouseID           int       `json:"house_id"`
	HouseCityID       int       `json:"house_city_id"`
	UserID            int       `json:"user_id"`
	ReservationID     int       `json:"reservation_id"`
	ReservationCode   string    `json:"reservation_code"`
	CheckInDate       time.Time `json:"check_in_date"`
	CheckOutDate      time.Time `json:"check_out_date"`
	TotalPrice        int       `json:"total_price"`
	ReservationStatus string    `json:"reservation_status"`
}

type AddReservationRequestDTO struct {
	HouseID      int    `json:"house_id" binding:"required"`
	CheckInDate  string `json:"check_in_date" binding:"required"`
	CheckOutDate string `json:"check_out_date" binding:"required"`
}

type AddReservationResponseDTO struct {
	ReservationID   int       `json:"reservation_id"`
	HouseID         int       `json:"house_id"`
	UserID          int       `json:"user_id"`
	ReservationCode string    `json:"reservation_code" gorm:"auto_increment;embedded;embeddedPrefix:RC" `
	CheckInDate     time.Time `json:"check_in_date"`
	CheckOutDate    time.Time `json:"check_out_date"`
	TotalPrice      int       `json:"total_price"`
}

type GetHouseReservationsByUserResponseDTO struct {
	Pagination  *helpers.Pagination     `json:"pagination"`
	Reservation []*HouseReservationsDTO `json:"reservation"`
}

type GetHouseReservationsByHouseResponseDTO struct {
	Pagination  *helpers.Pagination     `json:"pagination"`
	Reservation []*HouseReservationsDTO `json:"reservation"`
}

type GetHouseReservationsByIDResponseDTO struct {
	Reservation *HouseReservationsDTO `json:"reservation"`
}
