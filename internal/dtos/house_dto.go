package dtos

import "git.garena.com/sea-labs-id/batch-02/natanael-benediktus/final-project-backend/internal/models"

type GetAllHousesResponseDTO struct {
	House []*models.House `json:"house"`
}
