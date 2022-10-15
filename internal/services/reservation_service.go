package services

import (
	"fmt"
	"strings"
	"time"

	"git.garena.com/sea-labs-id/batch-02/natanael-benediktus/final-project-backend/internal/dtos"
	"git.garena.com/sea-labs-id/batch-02/natanael-benediktus/final-project-backend/internal/models"
	r "git.garena.com/sea-labs-id/batch-02/natanael-benediktus/final-project-backend/internal/repositories"
)

type ReservationService interface {
	AddReservation(userID int, request *dtos.AddReservationRequestDTO) (*models.Reservation, error)
}

type reservationService struct {
	reservationRepository r.ReservationRepository
	houseRepository       r.HouseRepository
}

type RSConfig struct {
	ReservationRepository r.ReservationRepository
	HouseRepository       r.HouseRepository
}

func NewReservationService(c *RSConfig) ReservationService {
	return &reservationService{
		reservationRepository: c.ReservationRepository,
		houseRepository:       c.HouseRepository,
	}
}

func (s *reservationService) AddReservation(userID int, request *dtos.AddReservationRequestDTO) (*models.Reservation, error) {
	err := s.reservationRepository.ValidateReservation(request)

	if err != nil {
		return nil, err
	}

	house, err := s.houseRepository.GetHouseByHouseID(request.HouseID)

	if err != nil {
		return nil, err
	}

	lastReservationID, err := s.reservationRepository.GetLastReservationID()
	print(lastReservationID)
	if err != nil {
		return nil, err
	}

	reservationCode := fmt.Sprintf("RC%s%d%s", strings.ReplaceAll(request.CheckInDate, "-", ""), lastReservationID+1, strings.ReplaceAll(request.CheckOutDate, "-", ""))

	format := "2006-01-02"
	checkInDate, _ := time.Parse(format, request.CheckInDate)
	checkOutDate, _ := time.Parse(format, request.CheckOutDate)
	diffDay := int(checkOutDate.Sub(checkInDate).Hours() / 24)

	totalPrice := diffDay * house.HousePrice

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

	return reservation, nil
}
