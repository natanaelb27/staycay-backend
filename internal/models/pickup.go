package models

import (
	"time"

	"gorm.io/gorm"
)

type Pickup struct {
	ID             int            `json:"id" gorm:"primary_key"`
	UserID         int            `json:"user_id"`
	User           User           `json:"user" gorm:"foreignKey:UserID"`
	ReservationID  int            `json:"reservation_id"`
	Reservation    Reservation    `json:"reservation" gorm:"foreignKey:ReservationID"`
	PickupStatusID int            `json:"pickup_status_id"`
	PickupStatus   PickupStatus   `json:"pickup_status" gorm:"foreignKey:PickupStatusID"`
	PickupAddress  string         `json:"pickup_address"`
	CityID         int            `json:"city_id"`
	City           City           `json:"city" gorm:"foreignKey:CityID"`
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
	DeletedAt      gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
