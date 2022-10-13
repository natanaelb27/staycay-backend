package models

import (
	"time"

	"gorm.io/gorm"
)

type Transaction struct {
	ID            int            `json:"id" gorm:"primary_key"`
	UserID        int            `json:"user_id"`
	User          User           `json:"user" gorm:"foreignKey:UserID"`
	HouseID       int            `json:"house_id"`
	House         House          `json:"house" gorm:"foreignKey:HouseID"`
	ReservationID int            `json:"reservation_id"`
	Reservation   Reservation    `json:"reservation" gorm:"foreignKey:ReservationID"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
