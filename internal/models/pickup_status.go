package models

import (
	"time"

	"gorm.io/gorm"
)

type PickupStatus struct {
	ID        int            `json:"id" gorm:"primary_key"`
	Status    string         `json:"status"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
