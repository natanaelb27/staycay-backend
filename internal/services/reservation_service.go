package services

import (
	"fmt"
	"strings"
	"time"

	"git.garena.com/sea-labs-id/batch-02/natanael-benediktus/final-project-backend/internal/dtos"
	e "git.garena.com/sea-labs-id/batch-02/natanael-benediktus/final-project-backend/internal/errors"
	"git.garena.com/sea-labs-id/batch-02/natanael-benediktus/final-project-backend/internal/models"
	r "git.garena.com/sea-labs-id/batch-02/natanael-benediktus/final-project-backend/internal/repositories"
)

type ReservationService interface {
	AddReservation(userID int, request *dtos.AddReservationRequestDTO) (*models.Reservation, error)
	GetHouseReservationsByUser(userID int, queryParams map[string]interface{}) ([]*dtos.HouseReservationsDTO, error)
	GetHouseReservationsByHouse(houseID int, userID int, queryParams map[string]interface{}) ([]*dtos.HouseReservationsDTO, error)
}

type reservationService struct {
	reservationRepository r.ReservationRepository
	houseRepository       r.HouseRepository
	userRepository        r.UserRepository
	walletRepository      r.WalletRepository
}

type RSConfig struct {
	ReservationRepository r.ReservationRepository
	HouseRepository       r.HouseRepository
	UserRepository        r.UserRepository
	WalletRepository      r.WalletRepository
}

func NewReservationService(c *RSConfig) ReservationService {
	return &reservationService{
		reservationRepository: c.ReservationRepository,
		houseRepository:       c.HouseRepository,
		userRepository:        c.UserRepository,
		walletRepository:      c.WalletRepository,
	}
}

func (s *reservationService) AddReservation(userID int, request *dtos.AddReservationRequestDTO) (*models.Reservation, error) {
	format := "2006-01-02"
	checkInDate, _ := time.Parse(format, request.CheckInDate)
	checkOutDate, _ := time.Parse(format, request.CheckOutDate)
	diffDay := int(checkOutDate.Sub(checkInDate).Hours() / 24)

	if diffDay == 0 {
		return nil, e.ErrInvalidReservationDateSameError
	}

	if diffDay < 0 {
		return nil, e.ErrInvalidReservationDateLaterError
	}

	now, _ := time.Parse(format, time.Now().Format(format))

	if now.After(checkInDate) {
		return nil, e.ErrInvalidReservationBeforeNowError
	}

	err := s.reservationRepository.ValidateReservation(request)

	if err != nil {
		return nil, err
	}

	house, err := s.houseRepository.GetHouseByHouseID(request.HouseID)

	if err != nil {
		return nil, err
	}

	totalPrice := diffDay * house.HousePrice

	user, err := s.userRepository.GetUserDetails(userID)

	if err != nil {
		return nil, err
	}

	if user.Balance < totalPrice {
		return nil, e.ErrBalanceNotEnoughError
	}

	lastReservationID, err := s.reservationRepository.GetLastReservationID()

	if err != nil {
		return nil, err
	}

	reservationCode := fmt.Sprintf("RC%s%d%s",
		strings.ReplaceAll(request.CheckInDate, "-", ""),
		lastReservationID+1,
		strings.ReplaceAll(request.CheckOutDate, "-", ""))

	newReservation := &models.Reservation{
		HouseID:         request.HouseID,
		UserID:          userID,
		ReservationCode: reservationCode,
		CheckInDate:     checkInDate,
		CheckOutDate:    checkOutDate,
		TotalPrice:      totalPrice,
	}

	reservation, err := s.reservationRepository.CreateReservation(userID, newReservation)

	if err != nil {
		return nil, err
	}

	_, err = s.walletRepository.SubtractWalletBalance(userID, totalPrice)

	if err != nil {
		return nil, err
	}

	return reservation, nil
}

func (s *reservationService) GetHouseReservationsByUser(userID int, queryParams map[string]interface{}) ([]*dtos.HouseReservationsDTO, error) {
	filterBy := fmt.Sprintf("%v", queryParams["filterBy"])
	houseReservations, err := s.reservationRepository.GetHouseReservationsByUserID(userID, filterBy)

	if err != nil {
		return nil, err
	}

	return houseReservations, nil
}

func (s *reservationService) GetHouseReservationsByHouse(houseID int, userID int, queryParams map[string]interface{}) ([]*dtos.HouseReservationsDTO, error) {
	houseDetails, err := s.houseRepository.GetHouseByHouseID(houseID)

	if err != nil {
		return nil, err
	}

	if houseDetails.HostID != userID {
		return nil, e.ErrInvalidHostError
	}

	filterBy := fmt.Sprintf("%v", queryParams["filterBy"])
	houseReservations, err := s.reservationRepository.GetHouseReservationsByHouseID(houseID, filterBy)

	if err != nil {
		return nil, err
	}

	return houseReservations, nil
}
