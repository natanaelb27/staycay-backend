package models

import (
	"time"

	"gorm.io/gorm"
)

type House struct {
	ID            int            `json:"id" gorm:"primary_key"`
	Name          string         `json:"name"`
	UserID        int            `json:"user_id"`
	User          User           `json:"user" gorm:"foreignKey:UserID"`
	Price         int            `json:"price"`
	Description   string         `json:"description"`
	CityID        int            `json:"city_id"`
	City          City           `json:"city" gorm:"foreignKey:CityID"`
	MaxGuest      int            `json:"max_guest"`
	HouseStatusID int            `json:"house_status_id"`
	HouseStatus   HouseStatus    `json:"house_status" gorm:"foreignKey:HouseStatusID"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
