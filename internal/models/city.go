package models

import (
	"time"

	"gorm.io/gorm"
)

type City struct {
	ID        int            `json:"id" gorm:"primary_key"`
	Name      string         `json:"name"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
