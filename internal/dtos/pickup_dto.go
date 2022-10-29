package dtos

import "git.garena.com/sea-labs-id/batch-02/natanael-benediktus/final-project-backend/internal/helpers"

type RequestPickupRequestDTO struct {
	PickupAddress string `json:"pickup_address" binding:"required"`
	CityID        int    `json:"city_id" binding:"required"`
}

type RequestPickupResponseDTO struct {
	PickupID       int    `json:"pickup_id"`
	ReservationID  int    `json:"reservation_id"`
	PickupStatusID int    `json:"pickup_status_id"`
	PickupAddress  string `json:"pickup_address"`
	PickupCityID   int    `json:"pickup_city_id"`
	PickupPrice    int    `json:"pickup_price"`
}

// type GetPickupStatusByReservationResponseDTO struct {
// 	ReservationID    int    `json:"reservation_id"`
// 	PickupStatusName string `json:"pickup_status_name"`
// }

type GetPickupDetailsByReservationResponseDTO struct {
	PickupID         int    `json:"pickup_id"`
	ReservationID    int    `json:"reservation_id"`
	HouseID          int    `json:"house_id"`
	PickupStatusName string `json:"pickup_status_name"`
	PickupAddress    string `json:"pickup_address"`
	PickupCityID     int    `json:"pickup_city_id"`
	PickupCityName   string `json:"pickup_city_name"`
	PickupPrice      int    `json:"pickup_price"`
}

type AdminUpdatePickupStatusRequestDTO struct {
	PickupStatusID int `json:"pickup_status_id"`
}

type AdminUpdatePickupStatusResponseDTO struct {
	ReservationID  int `json:"reservation_id"`
	PickupStatusID int `json:"pickup_status_id"`
}

type PickupStatusesDTO struct {
	PickupStatusID   int    `json:"pickup_status_id"`
	PickupStatusName string `json:"pickup_status_name"`
}

type GetAllPickupStatusesResponseDTO struct {
	PickupStatus []*PickupStatusesDTO `json:"pickup_status"`
}

type PickupDetailsDTO struct {
	PickupID         int    `json:"pickup_id"`
	ReservationID    int    `json:"reservation_id"`
	HouseID          int    `json:"house_id"`
	PickupStatusID   int    `json:"pickup_status_id"`
	PickupStatusName string `json:"pickup_status_name"`
	PickupAddress    string `json:"pickup_address"`
	PickupCityID     int    `json:"pickup_city_id"`
	PickupCityName   string `json:"pickup_city_name"`
	PickupPrice      int    `json:"pickup_price"`
}

type GetAllPickupDetailsResponseDTO struct {
	Pagination *helpers.Pagination `json:"pagination"`
	Pickup     []*PickupDetailsDTO `json:"pickup"`
}
