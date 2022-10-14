package models

import (
	"time"

	"gorm.io/gorm"
)

type GameChance struct {
	ID        int            `json:"id" gorm:"primary_key"`
	UserID    int            `json:"user_id"`
	User      User           `json:"user" gorm:"foreignKey:UserID"`
	Chance    int            `json:"chance"`
	Counter   int            `json:"counter"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
