package dtos

import "git.garena.com/sea-labs-id/batch-02/natanael-benediktus/final-project-backend/internal/models"

type GetAllHousesResponseDTO struct {
	House []*models.House `json:"house"`
}

type AddNewHouseRequestDTO struct {
	Name        string   `json:"name" binding:"required"`
	Price       int      `json:"price" binding:"required"`
	Description string   `json:"description" binding:"required"`
	CityName    string   `json:"city_name" binding:"required"`
	MaxGuest    int      `json:"max_guest" binding:"required"`
	Photo       []string `json:"photo" binding:"required"`
}

type AddNewHouseResponseDTO struct {
	Name        string   `json:"name"`
	Price       int      `json:"price"`
	Description string   `json:"description"`
	CityID      int      `json:"city_id"`
	MaxGuest    int      `json:"max_guest"`
	Photo       []string `json:"photo"`
}
