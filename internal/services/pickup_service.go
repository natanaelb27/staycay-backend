package services

import (
	"git.garena.com/sea-labs-id/batch-02/natanael-benediktus/final-project-backend/internal/dtos"
	e "git.garena.com/sea-labs-id/batch-02/natanael-benediktus/final-project-backend/internal/errors"
	"git.garena.com/sea-labs-id/batch-02/natanael-benediktus/final-project-backend/internal/models"
	r "git.garena.com/sea-labs-id/batch-02/natanael-benediktus/final-project-backend/internal/repositories"
)

type PickupService interface {
	RequestPickup(reservationID int, userID int, request *dtos.RequestPickupRequestDTO) (*models.Pickup, error)
}

type pickupService struct {
	pickupRepository      r.PickupRepository
	reservationRepository r.ReservationRepository
	userRepository        r.UserRepository
}

type PSConfig struct {
	PickupRepository      r.PickupRepository
	ReservationRepository r.ReservationRepository
	UserRepository        r.UserRepository
}

func NewPickupService(c *PSConfig) PickupService {
	return &pickupService{
		pickupRepository:      c.PickupRepository,
		reservationRepository: c.ReservationRepository,
		userRepository:        c.UserRepository,
	}
}

func (s *pickupService) RequestPickup(reservationID int, userID int, request *dtos.RequestPickupRequestDTO) (*models.Pickup, error) {
	houseReservation, err := s.reservationRepository.GetHouseReservationsByID(reservationID)

	if err != nil {
		return nil, err
	}

	if houseReservation.UserID != userID {
		return nil, e.ErrInvalidUserReservationError
	}

	if houseReservation.ReservationStatus == "past" {
		return nil, e.ErrPickupCompletedReservation
	}

	if houseReservation.ReservationStatus == "ongoing" {
		return nil, e.ErrPickupOngoingReservation
	}

	pickUpPrice := 0
	if houseReservation.HouseCityID != request.CityID {
		pickUpPrice = 300000
	} else {
		pickUpPrice = 100000
	}

	user, err := s.userRepository.GetUserDetails(userID)

	if err != nil {
		return nil, err
	}

	if user.Balance < pickUpPrice {
		return nil, e.ErrBalanceNotEnoughError
	}

	pickupStatus, err := s.pickupRepository.GetPickupStatusByName("pending admin")

	if err != nil {
		return nil, err
	}

	pickup := &models.Pickup{
		ReservationID:  reservationID,
		PickupStatusID: pickupStatus.ID,
		PickupAddress:  request.PickupAddress,
		CityID:         houseReservation.HouseCityID,
		PickupPrice:    pickUpPrice,
	}
	pickupData, err := s.pickupRepository.CreatePickup(pickup)

	if err != nil {
		return nil, err
	}

	return pickupData, nil
}
