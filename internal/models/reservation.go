package models

import (
	"time"

	"gorm.io/gorm"
)

type Reservation struct {
	ID              int            `json:"id" gorm:"primary_key"`
	HouseID         int            `json:"house_id"`
	House           House          `json:"house" gorm:"foreignKey:HouseID"`
	UserID          int            `json:"user_id"`
	User            User           `json:"user" gorm:"foreignKey:UserID"`
	ReservationCode string         `json:"reservation_code"`
	CheckInDate     time.Time      `json:"check_in_date"`
	CheckOutDate    time.Time      `json:"check_out_date"`
	TotalPrice      int            `json:"total_price"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
	DeletedAt       gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
