package services

import (
	"github.com/natanaelb27/staycay-backend/internal/dtos"
	r "github.com/natanaelb27/staycay-backend/internal/repositories"
)

type CityService interface {
	GetAllCities() ([]*dtos.CitiesDTO, error)
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

func (s *cityService) GetAllCities() ([]*dtos.CitiesDTO, error) {
	cities, err := s.cityRepository.GetCity()

	if err != nil {
		return nil, err
	}

	return cities, nil
}
