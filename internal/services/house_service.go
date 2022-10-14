package services

import (
	"git.garena.com/sea-labs-id/batch-02/natanael-benediktus/final-project-backend/internal/models"
	r "git.garena.com/sea-labs-id/batch-02/natanael-benediktus/final-project-backend/internal/repositories"
)

type HouseService interface {
	GetAllHouses() ([]*models.House, error)
}

type houseService struct {
	houseRepository r.HouseRepository
}

type HSConfig struct {
	HouseRepository r.HouseRepository
}

func NewHouseService(c *HSConfig) HouseService {
	return &houseService{
		houseRepository: c.HouseRepository,
	}
}

func (s *houseService) GetAllHouses() ([]*models.House, error) {
	house, err := s.houseRepository.GetHouse()

	if err != nil {
		return nil, err
	}

	return house, nil
}
