package dtos

import "git.garena.com/sea-labs-id/batch-02/natanael-benediktus/final-project-backend/internal/models"

type GetAllCitiesResponseDTO struct {
	City []*models.City `json:"city"`
}
