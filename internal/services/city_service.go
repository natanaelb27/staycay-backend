package services

import (
	"git.garena.com/sea-labs-id/batch-02/natanael-benediktus/final-project-backend/internal/dtos"
	r "git.garena.com/sea-labs-id/batch-02/natanael-benediktus/final-project-backend/internal/repositories"
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

	cityList := make([]*dtos.CitiesDTO, 0)

	for _, city := range cities {
		cityList = append(cityList, &dtos.CitiesDTO{
			CityID:   city.ID,
			CityName: city.Name,
		})
	}

	return cityList, nil
}
