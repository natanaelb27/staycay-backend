package services

import (
	"git.garena.com/sea-labs-id/batch-02/natanael-benediktus/final-project-backend/internal/dtos"
	e "git.garena.com/sea-labs-id/batch-02/natanael-benediktus/final-project-backend/internal/errors"
	"git.garena.com/sea-labs-id/batch-02/natanael-benediktus/final-project-backend/internal/models"
	r "git.garena.com/sea-labs-id/batch-02/natanael-benediktus/final-project-backend/internal/repositories"
)

type HouseService interface {
	GetAllHouses() ([]*dtos.HouseDetailsDTO, error)
	AddNewHouse(userID int, request *dtos.AddNewHouseRequestDTO) (*models.House, []string, error)
}

type houseService struct {
	houseRepository r.HouseRepository
	cityRepository  r.CityRepository
	userRepository  r.UserRepository
}

type HSConfig struct {
	HouseRepository r.HouseRepository
	CityRepository  r.CityRepository
	UserRepository  r.UserRepository
}

func NewHouseService(c *HSConfig) HouseService {
	return &houseService{
		houseRepository: c.HouseRepository,
		cityRepository:  c.CityRepository,
		userRepository:  c.UserRepository,
	}
}

func (s *houseService) GetAllHouses() ([]*dtos.HouseDetailsDTO, error) {
	house, err := s.houseRepository.GetHouse()

	if err != nil {
		return nil, err
	}

	return house, nil
}

func (s *houseService) AddNewHouse(userID int, request *dtos.AddNewHouseRequestDTO) (*models.House, []string, error) {
	userDetails, err := s.userRepository.GetUserDetails(userID)

	if err != nil {
		return nil, nil, err
	}

	if userDetails.Role != "host" {
		return nil, nil, e.ErrInvalidRoleAddHouseError
	}

	city, err := s.cityRepository.GetCityIDByName(request.CityName)

	if err != nil {
		return nil, nil, err
	}

	newHouse := &models.House{
		Name:        request.Name,
		UserID:      userID,
		Price:       request.Price,
		Description: request.Description,
		CityID:      city.ID,
		MaxGuest:    request.MaxGuest,
		Status:      "active",
	}

	house, err := s.houseRepository.CreateNewHouse(newHouse)

	if err != nil {
		return nil, nil, err
	}

	photoList := make([]string, 0)

	if len(request.Photo) > 0 {
		for _, photo := range request.Photo {
			newHousePhoto := &models.HousePhoto{
				HouseID: house.ID,
				Photo:   photo,
			}

			housePhoto, err := s.houseRepository.AddHousePhoto(newHousePhoto)

			if err != nil {
				return nil, nil, err
			}

			photoList = append(photoList, housePhoto.Photo)
		}
	}

	return house, photoList, nil
}
