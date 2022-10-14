package services

import (
	"git.garena.com/sea-labs-id/batch-02/natanael-benediktus/final-project-backend/internal/models"
	r "git.garena.com/sea-labs-id/batch-02/natanael-benediktus/final-project-backend/internal/repositories"
)

type CityService interface {
	GetAllCities() ([]*models.City, error)
}

type cityService struct {
	cityRepository r.CityRepository
}

type CSConfig struct {
	CityRepository r.CityRepository
}

func NewCityService(c *CSConfig) CityService {
	return &cityService{
		cityRepository: c.CityRepository,
	}
}

func (s *cityService) GetAllCities() ([]*models.City, error) {
	city, err := s.cityRepository.GetCity()

	if err != nil {
		return nil, err
	}

	return city, nil
}
