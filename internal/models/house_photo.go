package models

import (
	"time"

	"gorm.io/gorm"
)

type HousePhoto struct {
	ID        int            `json:"id" gorm:"primary_key"`
	HouseID   int            `json:"house_id"`
	House     House          `json:"house" gorm:"foreignKey:HouseID"`
	Photo     string         `json:"photo"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
