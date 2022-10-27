package services

import (
	"strings"

	"git.garena.com/sea-labs-id/batch-02/natanael-benediktus/final-project-backend/internal/dtos"
	e "git.garena.com/sea-labs-id/batch-02/natanael-benediktus/final-project-backend/internal/errors"
	"git.garena.com/sea-labs-id/batch-02/natanael-benediktus/final-project-backend/internal/helpers"
	"git.garena.com/sea-labs-id/batch-02/natanael-benediktus/final-project-backend/internal/models"
	r "git.garena.com/sea-labs-id/batch-02/natanael-benediktus/final-project-backend/internal/repositories"
)

type HouseService interface {
	GetAllHouses(queryParams map[string]interface{}) ([]*dtos.HouseDetailsWithPhotoDTO, *helpers.Pagination, error)
	GetHouseByID(houseID int) (*dtos.HouseDetailsWithPhotoDTO, error)
	GetUnscopedHouseByID(houseID int) (*dtos.HouseDetailsWithPhotoDTO, error)
	GetHostHouses(hostID int, queryParams map[string]interface{}) ([]*dtos.HouseDetailsWithPhotoDTO, *helpers.Pagination, error)
	GetAllHouseStatuses() ([]*dtos.HouseStatusDTO, error)
	AddNewHouse(userID int, request *dtos.AddNewHouseRequestDTO) (*models.House, []string, error)
	UpdateHouse(houseID int, userID int, request *dtos.UpdateHouseRequestDTO) (*models.House, error)
	DeleteHouse(houseID int, userID int) (*models.House, error)
	AddHousePhoto(houseID int, userID int, request *dtos.AddHousePhotoRequestDTO) ([]string, error)
	DeleteHousePhoto(housePhotoID int, userID int) (*models.HousePhoto, error)
}

type houseService struct {
	houseRepository       r.HouseRepository
	cityRepository        r.CityRepository
	userRepository        r.UserRepository
	reservationRepository r.ReservationRepository
}

type HSConfig struct {
	HouseRepository       r.HouseRepository
	CityRepository        r.CityRepository
	UserRepository        r.UserRepository
	ReservationRepository r.ReservationRepository
}

func NewHouseService(c *HSConfig) HouseService {
	return &houseService{
		houseRepository:       c.HouseRepository,
		cityRepository:        c.CityRepository,
		userRepository:        c.UserRepository,
		reservationRepository: c.ReservationRepository,
	}
}

func (s *houseService) GetAllHouses(queryParams map[string]interface{}) ([]*dtos.HouseDetailsWithPhotoDTO, *helpers.Pagination, error) {

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

	houseDetails, pagination, err := s.houseRepository.GetHouse(queryParams)

	if err != nil {
		return nil, nil, err
	}

	houseData := []*dtos.HouseDetailsWithPhotoDTO{}
	for _, house := range houseDetails {
		housePhoto, err := s.houseRepository.GetHousePhotoByHouseID(house.HouseID)

		if err != nil {
			return nil, nil, err
		}

		houseData = append(houseData, &dtos.HouseDetailsWithPhotoDTO{
			HouseID:          house.HouseID,
			HouseName:        house.HouseName,
			HousePrice:       house.HousePrice,
			HouseDescription: house.HouseDescription,
			HouseCityID:      house.HouseCityID,
			HouseCityName:    house.HouseCityName,
			HouseMaxGuest:    house.HouseMaxGuest,
			HouseStatus:      house.HouseStatus,
			HousePhoto:       housePhoto,
			HostID:           house.HostID,
			HostName:         house.HostName,
			HostEmail:        house.HostEmail,
		})
	}

	return houseData, pagination, nil
}

func (s *houseService) GetHouseByID(houseID int) (*dtos.HouseDetailsWithPhotoDTO, error) {
	houseDetails, err := s.houseRepository.GetHouseByHouseID(houseID)

	if err != nil {
		return nil, err
	}

	housePhoto, err := s.houseRepository.GetHousePhotoByHouseID(houseID)

	if err != nil {
		return nil, err
	}

	houseData := &dtos.HouseDetailsWithPhotoDTO{
		HouseID:          houseDetails.HouseID,
		HouseName:        houseDetails.HouseName,
		HousePrice:       houseDetails.HousePrice,
		HouseDescription: houseDetails.HouseDescription,
		HouseCityID:      houseDetails.HouseCityID,
		HouseCityName:    houseDetails.HouseCityName,
		HouseMaxGuest:    houseDetails.HouseMaxGuest,
		HouseStatus:      houseDetails.HouseStatus,
		HousePhoto:       housePhoto,
		HostID:           houseDetails.HostID,
		HostName:         houseDetails.HostName,
		HostEmail:        houseDetails.HostEmail,
	}

	return houseData, nil
}

func (s *houseService) GetUnscopedHouseByID(houseID int) (*dtos.HouseDetailsWithPhotoDTO, error) {
	houseDetails, err := s.houseRepository.GetUnscopedHouseByHouseID(houseID)

	if err != nil {
		return nil, err
	}

	housePhoto, err := s.houseRepository.GetUnscopedHousePhotoByHouseID(houseID)

	if err != nil {
		return nil, err
	}

	houseData := &dtos.HouseDetailsWithPhotoDTO{
		HouseID:          houseDetails.HouseID,
		HouseName:        houseDetails.HouseName,
		HousePrice:       houseDetails.HousePrice,
		HouseDescription: houseDetails.HouseDescription,
		HouseCityID:      houseDetails.HouseCityID,
		HouseCityName:    houseDetails.HouseCityName,
		HouseMaxGuest:    houseDetails.HouseMaxGuest,
		HouseStatus:      houseDetails.HouseStatus,
		HousePhoto:       housePhoto,
		HostID:           houseDetails.HostID,
		HostName:         houseDetails.HostName,
		HostEmail:        houseDetails.HostEmail,
	}

	return houseData, nil
}

func (s *houseService) GetHostHouses(hostID int, queryParams map[string]interface{}) ([]*dtos.HouseDetailsWithPhotoDTO, *helpers.Pagination, error) {
	userDetails, err := s.userRepository.GetUserDetails(hostID)

	if err != nil {
		return nil, nil, err
	}

	if userDetails.Role != "host" {
		return nil, nil, e.ErrInvalidRoleNotHostError
	}

	houseDetails, pagination, err := s.houseRepository.GetHouseByHostID(hostID, queryParams)

	if err != nil {
		return nil, nil, err
	}

	houseData := []*dtos.HouseDetailsWithPhotoDTO{}
	for _, house := range houseDetails {
		housePhoto, err := s.houseRepository.GetHousePhotoByHouseID(house.HouseID)

		if err != nil {
			return nil, nil, err
		}

		houseData = append(houseData, &dtos.HouseDetailsWithPhotoDTO{
			HouseID:          house.HouseID,
			HouseName:        house.HouseName,
			HousePrice:       house.HousePrice,
			HouseDescription: house.HouseDescription,
			HouseCityID:      house.HouseCityID,
			HouseCityName:    house.HouseCityName,
			HouseMaxGuest:    house.HouseMaxGuest,
			HouseStatus:      house.HouseStatus,
			HousePhoto:       housePhoto,
			HostID:           house.HostID,
			HostName:         house.HostName,
			HostEmail:        house.HostEmail,
		})
	}

	return houseData, pagination, nil
}

func (s *houseService) GetAllHouseStatuses() ([]*dtos.HouseStatusDTO, error) {
	houseStatus, err := s.houseRepository.GetHouseStatus()

	if err != nil {
		return nil, err
	}

	return houseStatus, nil
}

func (s *houseService) AddNewHouse(userID int, request *dtos.AddNewHouseRequestDTO) (*models.House, []string, error) {
	userDetails, err := s.userRepository.GetUserDetails(userID)

	if err != nil {
		return nil, nil, err
	}

	if userDetails.Role != "host" {
		return nil, nil, e.ErrInvalidRoleNotHostError
	}

	// city, err := s.cityRepository.GetCityIDByName(request.CityName)

	// if err != nil {
	// 	return nil, nil, err
	// }

	houseStatus, err := s.houseRepository.GetHouseStatusByName("Active")

	if err != nil {
		return nil, nil, err
	}

	newHouse := &models.House{
		Name:          request.Name,
		UserID:        userID,
		Price:         request.Price,
		Description:   request.Description,
		CityID:        request.CityID,
		MaxGuest:      request.MaxGuest,
		HouseStatusID: houseStatus.ID,
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

func (s *houseService) UpdateHouse(houseID int, userID int, request *dtos.UpdateHouseRequestDTO) (*models.House, error) {
	userDetails, err := s.userRepository.GetUserDetails(userID)

	if err != nil {
		return nil, err
	}

	houseDetails, err := s.houseRepository.GetHouseByHouseID(houseID)

	if err != nil {
		return nil, err
	}

	if strings.ToLower(userDetails.Role) != "host" && strings.ToLower(userDetails.Role) != "admin" {
		return nil, e.ErrInvalidRoleNotHostError
	}

	if houseDetails.HostID != userID && strings.ToLower(userDetails.Role) != "admin" {
		return nil, e.ErrInvalidHostError
	}

	if strings.ToLower(houseDetails.HouseStatus) == "deactivated by admin" && userDetails.Role != "admin" {
		request.HouseStatusID = 0
	}

	newData := &models.House{
		Name:          request.Name,
		Price:         request.Price,
		Description:   request.Description,
		CityID:        request.CityID,
		MaxGuest:      request.MaxGuest,
		HouseStatusID: request.HouseStatusID,
	}

	updatedHouse, err := s.houseRepository.UpdateHouseByHouseID(houseID, newData)

	if err != nil {
		return nil, err
	}

	return updatedHouse, nil
}

func (s *houseService) DeleteHouse(houseID int, userID int) (*models.House, error) {
	houseDetails, err := s.houseRepository.GetHouseByHouseID(houseID)

	if err != nil {
		return nil, err
	}

	user, err := s.userRepository.GetUserByID(userID)

	if err != nil {
		return nil, err
	}

	// ini ga perlu
	if strings.ToLower(user.Role) != "host" && strings.ToLower(user.Role) != "admin" {
		return nil, e.ErrInvalidRoleNotHostError
	}

	if houseDetails.HostID != userID && strings.ToLower(user.Role) != "admin" {
		return nil, e.ErrInvalidHostError
	}

	houseStatusName := ""
	if strings.ToLower(user.Role) == "admin" {
		houseStatusName = "deactivated by admin"
	} else {
		houseStatusName = "inactive"
	}

	houseStatus, err := s.houseRepository.GetHouseStatusByName(houseStatusName)

	if err != nil {
		return nil, err
	}

	newData := &models.House{
		HouseStatusID: houseStatus.ID,
	}

	_, err = s.houseRepository.UpdateHouseByHouseID(houseID, newData)

	if err != nil {
		return nil, err
	}

	queryParams := make(map[string]interface{})
	queryParams["filterBy"] = "upcoming"

	upcomingHouseReservation, _, err := s.reservationRepository.GetHouseReservationsByHouseID(houseID, queryParams)

	if err != nil {
		return nil, err

	}

	if len(upcomingHouseReservation) > 0 {
		return nil, e.ErrDeleteUpcomingReservationError
	}

	err = s.houseRepository.DeleteHousePhotoByHouseID(houseID)

	if err != nil {
		return nil, err
	}

	deletedHouse, err := s.houseRepository.DeleteHouseByID(houseID)

	if err != nil {
		return nil, err
	}

	return deletedHouse, nil
}

func (s *houseService) AddHousePhoto(houseID int, userID int, request *dtos.AddHousePhotoRequestDTO) ([]string, error) {
	userDetails, err := s.userRepository.GetUserDetails(userID)

	if err != nil {
		return nil, err
	}

	if userDetails.Role != "host" {
		return nil, e.ErrInvalidRoleNotHostError
	}

	photoList := make([]string, 0)

	if len(request.Photo) > 0 {
		for _, photo := range request.Photo {
			newHousePhoto := &models.HousePhoto{
				HouseID: houseID,
				Photo:   photo,
			}

			housePhoto, err := s.houseRepository.AddHousePhoto(newHousePhoto)

			if err != nil {
				return nil, err
			}

			photoList = append(photoList, housePhoto.Photo)
		}
	}

	return photoList, nil
}

func (s *houseService) DeleteHousePhoto(housePhotoID int, userID int) (*models.HousePhoto, error) {
	houseDetails, err := s.houseRepository.GetHouseByHousePhotoID(housePhotoID)

	if err != nil {
		return nil, err
	}

	if houseDetails.HostID != userID {
		return nil, e.ErrInvalidHostError
	}

	totalRows, err := s.houseRepository.ValidateTotalHousePhoto(houseDetails.HouseID)

	if totalRows <= 1 {
		return nil, e.ErrHousePhotoMustMoreThanOneError
	}

	housePhoto, err := s.houseRepository.DeleteHousePhotoByID(housePhotoID)
	if err != nil {
		return nil, err
	}

	return housePhoto, nil
}
