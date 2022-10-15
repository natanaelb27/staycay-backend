package services

import (
	"errors"

	"git.garena.com/sea-labs-id/batch-02/natanael-benediktus/final-project-backend/internal/dtos"
	e "git.garena.com/sea-labs-id/batch-02/natanael-benediktus/final-project-backend/internal/errors"
	"git.garena.com/sea-labs-id/batch-02/natanael-benediktus/final-project-backend/internal/models"
	r "git.garena.com/sea-labs-id/batch-02/natanael-benediktus/final-project-backend/internal/repositories"
)

type HouseService interface {
	GetAllHouses(queryParams map[string]interface{}, querySearchCity []string) ([]*dtos.HouseDetailsWithPhotoDTO, error)
	GetHostHouses(hostID int) ([]*dtos.HouseDetailsWithPhotoDTO, error)
	AddNewHouse(userID int, request *dtos.AddNewHouseRequestDTO) (*models.House, []string, error)
	UpdateHouse(houseID int, userID int, request *dtos.UpdateHouseRequestDTO) (*models.House, []string, error)
	DeleteHouse(houseID int, userID int) (*models.House, error)
	DeleteHousePhoto(housePhotoID int, userID int) (*models.HousePhoto, error)
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

func (s *houseService) GetAllHouses(queryParams map[string]interface{}, querySearchCity []string) ([]*dtos.HouseDetailsWithPhotoDTO, error) {

	// if rec, ok := query["searchCity"].(map[string]interface{}); ok {
	// 	for key, val := range rec {
	// 		print(key, val)
	// 	}
	// } else {
	// 	print("HAHaH")
	// }

	// print("SORT: ", query["sort"])
	// print(query["sortby"])
	// print("SEARCH: ", query["searchName"])
	// print(query["searchCity"])

	// house, err := s.houseRepository.GetHouse(queryParams, querySearchCity)

	// if err != nil {
	// 	return nil, err
	// }

	houseDetails, err := s.houseRepository.GetHouse(queryParams, querySearchCity)

	if err != nil {
		return nil, err
	}

	houseData := []*dtos.HouseDetailsWithPhotoDTO{}
	for _, house := range houseDetails {
		housePhoto, err := s.houseRepository.GetHousePhotoByHouseID(house.HouseID)

		if err != nil {
			return nil, err
		}

		houseData = append(houseData, &dtos.HouseDetailsWithPhotoDTO{
			HouseID:          house.HouseID,
			HouseName:        house.HouseName,
			HousePrice:       house.HousePrice,
			HouseDescription: house.HouseDescription,
			HouseCityName:    house.HouseCityName,
			HouseMaxGuest:    house.HouseMaxGuest,
			HouseStatus:      house.HouseStatus,
			HousePhoto:       housePhoto,
			HostID:           house.HostID,
			HostName:         house.HostName,
			HostEmail:        house.HostEmail,
		})
	}

	return houseData, nil
}

func (s *houseService) GetHostHouses(userID int) ([]*dtos.HouseDetailsWithPhotoDTO, error) {
	userDetails, err := s.userRepository.GetUserDetails(userID)

	if err != nil {
		return nil, err
	}

	if userDetails.Role != "host" {
		return nil, e.ErrInvalidRoleNotHostError
	}

	houseDetails, err := s.houseRepository.GetHouseByHostID(userID)

	if err != nil {
		return nil, err
	}

	houseData := []*dtos.HouseDetailsWithPhotoDTO{}
	for _, house := range houseDetails {
		housePhoto, err := s.houseRepository.GetHousePhotoByHouseID(house.HouseID)

		if err != nil {
			return nil, err
		}

		houseData = append(houseData, &dtos.HouseDetailsWithPhotoDTO{
			HouseID:          house.HouseID,
			HouseName:        house.HouseName,
			HousePrice:       house.HousePrice,
			HouseDescription: house.HouseDescription,
			HouseCityName:    house.HouseCityName,
			HouseMaxGuest:    house.HouseMaxGuest,
			HouseStatus:      house.HouseStatus,
			HousePhoto:       housePhoto,
			HostID:           house.HostID,
			HostName:         house.HostName,
			HostEmail:        house.HostEmail,
		})
	}

	return houseData, nil
}

func (s *houseService) AddNewHouse(userID int, request *dtos.AddNewHouseRequestDTO) (*models.House, []string, error) {
	userDetails, err := s.userRepository.GetUserDetails(userID)

	if err != nil {
		return nil, nil, err
	}

	if userDetails.Role != "host" {
		return nil, nil, e.ErrInvalidRoleNotHostError
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

func (s *houseService) UpdateHouse(houseID int, userID int, request *dtos.UpdateHouseRequestDTO) (*models.House, []string, error) {
	userDetails, err := s.userRepository.GetUserDetails(userID)

	if err != nil {
		return nil, nil, err
	}

	if userDetails.Role != "host" {
		return nil, nil, e.ErrInvalidRoleNotHostError
	}

	house, err := s.houseRepository.GetHouseByHouseID(houseID)

	if err != nil {
		return nil, nil, err
	}

	if house.HostID != userID {
		return nil, nil, e.ErrInvalidHostError
	}

	city := &models.City{}

	if request.CityName != "" {
		city, err = s.cityRepository.GetCityIDByName(request.CityName)

		if err != nil {
			return nil, nil, err
		}
	}

	newData := &models.House{
		Name:        request.Name,
		Price:       request.Price,
		Description: request.Description,
		CityID:      city.ID,
		MaxGuest:    request.MaxGuest,
		Status:      request.Status,
	}

	updatedHouse, err := s.houseRepository.UpdateHouseByHouseID(houseID, newData)

	if err != nil {
		return nil, nil, err
	}

	photoList := make([]string, 0)

	if len(request.Photo) > 0 {
		for _, photo := range request.Photo {
			newHousePhoto := &models.HousePhoto{
				HouseID: updatedHouse.ID,
				Photo:   photo,
			}

			housePhoto, err := s.houseRepository.AddHousePhoto(newHousePhoto)

			if err != nil {
				return nil, nil, err
			}

			photoList = append(photoList, housePhoto.Photo)
		}
	}

	return updatedHouse, photoList, nil
}

func (s *houseService) DeleteHouse(houseID int, userID int) (*models.House, error) {
	houseDetails, err := s.houseRepository.GetHouseByHouseID(houseID)

	if err != nil {
		return nil, err
	}

	if houseDetails.HostID != userID {
		return nil, e.ErrInvalidHostError
	}

	_, err = s.houseRepository.GetUpcomingHouseReservationsByHouseID(houseID)

	if err != nil {
		if errors.Is(err, e.ErrDeleteUpcomingReservationError) {

			newData := &models.House{
				Status: "inactive",
			}

			house, err := s.houseRepository.UpdateHouseByHouseID(houseID, newData)

			if err != nil {
				return nil, err
			}

			return house, e.ErrDeleteUpcomingReservationError

		} else {
			return nil, err
		}
	}

	deletedHouse, err := s.houseRepository.DeleteHouseByID(houseID)

	if err != nil {
		return nil, err
	}

	return deletedHouse, nil
}

func (s *houseService) DeleteHousePhoto(housePhotoID int, userID int) (*models.HousePhoto, error) {
	houseDetails, err := s.houseRepository.GetHouseByHousePhotoID(housePhotoID)

	if err != nil {
		return nil, err
	}

	if houseDetails.HostID != userID {
		return nil, e.ErrInvalidHostError
	}

	housePhoto, err := s.houseRepository.DeleteHousePhotoByID(housePhotoID)
	if err != nil {
		return nil, err
	}

	return housePhoto, nil
}
