package dtos

type RequestPickupRequestDTO struct {
	PickupAddress string `json:"pickup_address" binding:"required"`
	CityID        int    `json:"city_id" binding:"required"`
}

type RequestPickupResponseDTO struct {
	PickupID       int    `json:"pickup_id"`
	ReservationID  int    `json:"reservation_id"`
	PickupStatusID int    `json:"pickup_status_id"`
	PickupAddress  string `json:"pickup_address"`
	CityID         int    `json:"city_id"`
	PickupPrice    int    `json:"pickup_price"`
}

type GetPickupStatusByReservationResponseDTO struct {
	ReservationID    int    `json:"reservation_id"`
	PickupStatusName string `json:"pickup_status_name"`
}
