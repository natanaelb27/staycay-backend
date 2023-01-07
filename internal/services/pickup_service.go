package services

import (
	"strings"
	"time"

	"github.com/natanaelb27/staycay-backend/internal/dtos"
	e "github.com/natanaelb27/staycay-backend/internal/errors"
	"github.com/natanaelb27/staycay-backend/internal/helpers"
	"github.com/natanaelb27/staycay-backend/internal/models"
	r "github.com/natanaelb27/staycay-backend/internal/repositories"
)

type PickupService interface {
	RequestPickup(reservationID int, userID int, request *dtos.RequestPickupRequestDTO) (*models.Pickup, error)
	GetPickupDetailsByReservation(reservationID int, userID int) (*dtos.GetPickupDetailsByReservationResponseDTO, error)
	AdminUpdatePickupStatus(pickupStatusID int, pickupID int, userID int) (*models.Pickup, error)
	GetAllPickupStatuses() ([]*dtos.PickupStatusesDTO, error)
	GetAllPickupDetails(queryParams map[string]interface{}) ([]*dtos.PickupDetailsDTO, *helpers.Pagination, error)
}

type pickupService struct {
	pickupRepository      r.PickupRepository
	reservationRepository r.ReservationRepository
	userRepository        r.UserRepository
	walletRepository      r.WalletRepository
}

type PSConfig struct {
	PickupRepository      r.PickupRepository
	ReservationRepository r.ReservationRepository
	UserRepository        r.UserRepository
	WalletRepository      r.WalletRepository
}

func NewPickupService(c *PSConfig) PickupService {
	return &pickupService{
		pickupRepository:      c.PickupRepository,
		reservationRepository: c.ReservationRepository,
		userRepository:        c.UserRepository,
		walletRepository:      c.WalletRepository,
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

	if houseReservation.ReservationStatus == "completed" {
		return nil, e.ErrPickupCompletedReservation
	}

	if houseReservation.ReservationStatus == "ongoing" {
		return nil, e.ErrPickupOngoingReservation
	}

	err = s.pickupRepository.ValidateOnlyOnePickupReservation(reservationID)

	if err != nil {
		return nil, err
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
		return nil, e.ErrInsufficientBalanceError
	}

	pickupStatus, err := s.pickupRepository.GetPickupStatusByName("Pending Admin")

	if err != nil {
		return nil, err
	}

	pickup := &models.Pickup{
		ReservationID:  reservationID,
		PickupStatusID: pickupStatus.ID,
		PickupAddress:  request.PickupAddress,
		CityID:         request.CityID,
		PickupPrice:    pickUpPrice,
	}
	pickupData, err := s.pickupRepository.CreatePickup(pickup)

	if err != nil {
		return nil, err
	}

	_, err = s.walletRepository.SubtractWalletBalance(userID, pickUpPrice)

	if err != nil {
		return nil, err
	}

	return pickupData, nil
}

func (s *pickupService) GetPickupDetailsByReservation(reservationID int, userID int) (*dtos.GetPickupDetailsByReservationResponseDTO, error) {
	// err := s.pickupRepository.ValidatePickupReservationExists(reservationID)

	// if err != nil {
	// 	return nil, err
	// }

	houseReservation, err := s.reservationRepository.GetHouseReservationsByID(reservationID)

	if err != nil {
		return nil, err
	}

	if houseReservation.UserID != userID {
		return nil, e.ErrInvalidUserReservationError
	}

	reservationPickup, err := s.pickupRepository.GetPickupDetailsByReservationID(reservationID)

	if err != nil {
		return nil, err
	}

	return reservationPickup, nil
}

func (s *pickupService) AdminUpdatePickupStatus(pickupStatusID int, pickupID int, userID int) (*models.Pickup, error) {
	user, err := s.userRepository.GetUserDetails(userID)

	if err != nil {
		return nil, err
	}

	if user.Role != "admin" {
		return nil, e.ErrInvalidRoleNotAdminError
	}

	// err = s.pickupRepository.ValidatePickupReservationExists(reservationID)

	// if err != nil {
	// 	return nil, err
	// }

	pickup, err := s.pickupRepository.GetPickupByID(pickupID)

	if err != nil {
		return nil, err
	}

	pickupStatus, err := s.pickupRepository.GetPickupStatusByID(pickupStatusID)

	if err != nil {
		return nil, err
	}

	if strings.ToLower(pickupStatus.Status) == "on the way to pickup" || strings.ToLower(pickupStatus.Status) == "on the way to reservation" || strings.ToLower(pickupStatus.Status) == "completed" {
		houseReservation, err := s.reservationRepository.GetHouseReservationsByID(pickup.ReservationID)

		if err != nil {
			return nil, err
		}

		format := "2006-01-02"
		now, _ := time.Parse(format, time.Now().Format(format))
		checkInDate, _ := time.Parse(format, houseReservation.CheckInDate.Format(format))

		if !checkInDate.Equal(now) {
			return nil, e.ErrInvalidCheckInDateUpdatePickupStatusError
		}
	}

	updatedPickup, err := s.pickupRepository.UpdatePickupStatusByID(pickupStatusID, pickupID)

	if err != nil {
		return nil, err
	}

	return updatedPickup, nil
}

func (s *pickupService) GetAllPickupStatuses() ([]*dtos.PickupStatusesDTO, error) {
	pickupStatuses, err := s.pickupRepository.GetPickupStatuses()

	if err != nil {
		return nil, err
	}

	return pickupStatuses, nil
}

func (s *pickupService) GetAllPickupDetails(queryParams map[string]interface{}) ([]*dtos.PickupDetailsDTO, *helpers.Pagination, error) {
	pickupStatuses, pagination, err := s.pickupRepository.GetPickupDetails(queryParams)

	if err != nil {
		return nil, nil, err
	}

	return pickupStatuses, pagination, nil
}
