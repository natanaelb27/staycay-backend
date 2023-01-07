package services

import (
	"errors"
	"reflect"
	"testing"
	"time"

	"github.com/natanaelb27/staycay-backend/internal/dtos"
	"github.com/natanaelb27/staycay-backend/internal/helpers"
	"github.com/natanaelb27/staycay-backend/internal/mocks"
	"github.com/natanaelb27/staycay-backend/internal/models"
	"github.com/stretchr/testify/assert"
)

func Test_pickupService_NewPickupService(t *testing.T) {
	pickupRepository := mocks.NewPickupRepository(t)
	reservationRepository := mocks.NewReservationRepository(t)
	userRepository := mocks.NewUserRepository(t)
	walletRepository := mocks.NewWalletRepository(t)
	pickupService := NewPickupService(
		&PSConfig{
			PickupRepository:      pickupRepository,
			ReservationRepository: reservationRepository,
			UserRepository:        userRepository,
			WalletRepository:      walletRepository,
		},
	)

	assert.NotNil(t, pickupService)
}

func Test_pickupService_RequestPickup(t *testing.T) {
	type args struct {
		reservationID int
		userID        int
		request       *dtos.RequestPickupRequestDTO
	}
	tests := []struct {
		name                  string
		pickupRepository      *mocks.PickupRepository
		reservationRepository *mocks.ReservationRepository
		userRepository        *mocks.UserRepository
		walletRepository      *mocks.WalletRepository
		mock                  func(m *mocks.PickupRepository, m2 *mocks.ReservationRepository, m3 *mocks.UserRepository, m4 *mocks.WalletRepository)
		args                  args
		want                  *models.Pickup
		wantErr               bool
	}{
		{
			name:                  "Failed - GetHouseReservationsByID",
			pickupRepository:      mocks.NewPickupRepository(t),
			reservationRepository: mocks.NewReservationRepository(t),
			userRepository:        mocks.NewUserRepository(t),
			walletRepository:      mocks.NewWalletRepository(t),
			mock: func(m *mocks.PickupRepository, m2 *mocks.ReservationRepository, m3 *mocks.UserRepository, m4 *mocks.WalletRepository) {
				m2.On("GetHouseReservationsByID", 1).Return(nil, errors.New("error"))
			},
			args: args{
				reservationID: 1,
				userID:        1,
				request:       &dtos.RequestPickupRequestDTO{},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name:                  "Failed - User is Not The Customer",
			pickupRepository:      mocks.NewPickupRepository(t),
			reservationRepository: mocks.NewReservationRepository(t),
			userRepository:        mocks.NewUserRepository(t),
			walletRepository:      mocks.NewWalletRepository(t),
			mock: func(m *mocks.PickupRepository, m2 *mocks.ReservationRepository, m3 *mocks.UserRepository, m4 *mocks.WalletRepository) {
				m2.On("GetHouseReservationsByID", 1).Return(&dtos.HouseReservationsDTO{
					UserID: 2,
				}, nil)
			},
			args: args{
				reservationID: 1,
				userID:        1,
				request:       &dtos.RequestPickupRequestDTO{},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name:                  "Failed - Request Pickup on Completed Reservation",
			pickupRepository:      mocks.NewPickupRepository(t),
			reservationRepository: mocks.NewReservationRepository(t),
			userRepository:        mocks.NewUserRepository(t),
			walletRepository:      mocks.NewWalletRepository(t),
			mock: func(m *mocks.PickupRepository, m2 *mocks.ReservationRepository, m3 *mocks.UserRepository, m4 *mocks.WalletRepository) {
				m2.On("GetHouseReservationsByID", 1).Return(&dtos.HouseReservationsDTO{
					UserID:            1,
					ReservationStatus: "completed",
				}, nil)
			},
			args: args{
				reservationID: 1,
				userID:        1,
				request:       &dtos.RequestPickupRequestDTO{},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name:                  "Failed - Request Pickup on Ongoing Reservation",
			pickupRepository:      mocks.NewPickupRepository(t),
			reservationRepository: mocks.NewReservationRepository(t),
			userRepository:        mocks.NewUserRepository(t),
			walletRepository:      mocks.NewWalletRepository(t),
			mock: func(m *mocks.PickupRepository, m2 *mocks.ReservationRepository, m3 *mocks.UserRepository, m4 *mocks.WalletRepository) {
				m2.On("GetHouseReservationsByID", 1).Return(&dtos.HouseReservationsDTO{
					UserID:            1,
					ReservationStatus: "ongoing",
				}, nil)
			},
			args: args{
				reservationID: 1,
				userID:        1,
				request:       &dtos.RequestPickupRequestDTO{},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name:                  "Failed - ValidateOnlyOnePickupReservation",
			pickupRepository:      mocks.NewPickupRepository(t),
			reservationRepository: mocks.NewReservationRepository(t),
			userRepository:        mocks.NewUserRepository(t),
			walletRepository:      mocks.NewWalletRepository(t),
			mock: func(m *mocks.PickupRepository, m2 *mocks.ReservationRepository, m3 *mocks.UserRepository, m4 *mocks.WalletRepository) {
				m2.On("GetHouseReservationsByID", 1).Return(&dtos.HouseReservationsDTO{
					UserID: 1,
				}, nil)
				m.On("ValidateOnlyOnePickupReservation", 1).Return(errors.New("error"))
			},
			args: args{
				reservationID: 1,
				userID:        1,
				request:       &dtos.RequestPickupRequestDTO{},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name:                  "Failed - GetUserDetails",
			pickupRepository:      mocks.NewPickupRepository(t),
			reservationRepository: mocks.NewReservationRepository(t),
			userRepository:        mocks.NewUserRepository(t),
			walletRepository:      mocks.NewWalletRepository(t),
			mock: func(m *mocks.PickupRepository, m2 *mocks.ReservationRepository, m3 *mocks.UserRepository, m4 *mocks.WalletRepository) {
				m2.On("GetHouseReservationsByID", 1).Return(&dtos.HouseReservationsDTO{
					UserID: 1,
				}, nil)
				m.On("ValidateOnlyOnePickupReservation", 1).Return(nil)
				m3.On("GetUserDetails", 1).Return(nil, errors.New("error"))
			},
			args: args{
				reservationID: 1,
				userID:        1,
				request:       &dtos.RequestPickupRequestDTO{},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name:                  "Failed - Insufficient Balance to Pay Pickup Fee",
			pickupRepository:      mocks.NewPickupRepository(t),
			reservationRepository: mocks.NewReservationRepository(t),
			userRepository:        mocks.NewUserRepository(t),
			walletRepository:      mocks.NewWalletRepository(t),
			mock: func(m *mocks.PickupRepository, m2 *mocks.ReservationRepository, m3 *mocks.UserRepository, m4 *mocks.WalletRepository) {
				m2.On("GetHouseReservationsByID", 1).Return(&dtos.HouseReservationsDTO{
					UserID:      1,
					HouseCityID: 1,
				}, nil)
				m.On("ValidateOnlyOnePickupReservation", 1).Return(nil)
				m3.On("GetUserDetails", 1).Return(&dtos.UserDetailsDTO{
					Balance: 50000,
				}, nil)
			},
			args: args{
				reservationID: 1,
				userID:        1,
				request: &dtos.RequestPickupRequestDTO{
					CityID: 1,
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name:                  "Failed - GetPickupStatusByName",
			pickupRepository:      mocks.NewPickupRepository(t),
			reservationRepository: mocks.NewReservationRepository(t),
			userRepository:        mocks.NewUserRepository(t),
			walletRepository:      mocks.NewWalletRepository(t),
			mock: func(m *mocks.PickupRepository, m2 *mocks.ReservationRepository, m3 *mocks.UserRepository, m4 *mocks.WalletRepository) {
				m2.On("GetHouseReservationsByID", 1).Return(&dtos.HouseReservationsDTO{
					UserID:      1,
					HouseCityID: 1,
				}, nil)
				m.On("ValidateOnlyOnePickupReservation", 1).Return(nil)
				m3.On("GetUserDetails", 1).Return(&dtos.UserDetailsDTO{
					Balance: 500000,
				}, nil)
				m.On("GetPickupStatusByName", "Pending Admin").Return(nil, errors.New("error"))
			},
			args: args{
				reservationID: 1,
				userID:        1,
				request: &dtos.RequestPickupRequestDTO{
					CityID: 1,
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name:                  "Failed - CreatePickup",
			pickupRepository:      mocks.NewPickupRepository(t),
			reservationRepository: mocks.NewReservationRepository(t),
			userRepository:        mocks.NewUserRepository(t),
			walletRepository:      mocks.NewWalletRepository(t),
			mock: func(m *mocks.PickupRepository, m2 *mocks.ReservationRepository, m3 *mocks.UserRepository, m4 *mocks.WalletRepository) {
				m2.On("GetHouseReservationsByID", 1).Return(&dtos.HouseReservationsDTO{
					UserID:      1,
					HouseCityID: 1,
				}, nil)
				m.On("ValidateOnlyOnePickupReservation", 1).Return(nil)
				m3.On("GetUserDetails", 1).Return(&dtos.UserDetailsDTO{
					Balance: 500000,
				}, nil)
				m.On("GetPickupStatusByName", "Pending Admin").Return(&models.PickupStatus{
					ID: 1,
				}, nil)
				m.On("CreatePickup", &models.Pickup{
					ReservationID:  1,
					PickupStatusID: 1,
					CityID:         1,
					PickupPrice:    100000,
				}).Return(nil, errors.New("error"))
			},
			args: args{
				reservationID: 1,
				userID:        1,
				request: &dtos.RequestPickupRequestDTO{
					CityID: 1,
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name:                  "Failed - SubtractWalletBalance",
			pickupRepository:      mocks.NewPickupRepository(t),
			reservationRepository: mocks.NewReservationRepository(t),
			userRepository:        mocks.NewUserRepository(t),
			walletRepository:      mocks.NewWalletRepository(t),
			mock: func(m *mocks.PickupRepository, m2 *mocks.ReservationRepository, m3 *mocks.UserRepository, m4 *mocks.WalletRepository) {
				m2.On("GetHouseReservationsByID", 1).Return(&dtos.HouseReservationsDTO{
					UserID:      1,
					HouseCityID: 1,
				}, nil)
				m.On("ValidateOnlyOnePickupReservation", 1).Return(nil)
				m3.On("GetUserDetails", 1).Return(&dtos.UserDetailsDTO{
					Balance: 500000,
				}, nil)
				m.On("GetPickupStatusByName", "Pending Admin").Return(&models.PickupStatus{
					ID: 1,
				}, nil)
				m.On("CreatePickup", &models.Pickup{
					ReservationID:  1,
					PickupStatusID: 1,
					CityID:         1,
					PickupPrice:    100000,
				}).Return(&models.Pickup{}, nil)
				m4.On("SubtractWalletBalance", 1, 100000).Return(nil, errors.New("error"))
			},
			args: args{
				reservationID: 1,
				userID:        1,
				request: &dtos.RequestPickupRequestDTO{
					CityID: 1,
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name:                  "Success - Pickup City and House City is Different",
			pickupRepository:      mocks.NewPickupRepository(t),
			reservationRepository: mocks.NewReservationRepository(t),
			userRepository:        mocks.NewUserRepository(t),
			walletRepository:      mocks.NewWalletRepository(t),
			mock: func(m *mocks.PickupRepository, m2 *mocks.ReservationRepository, m3 *mocks.UserRepository, m4 *mocks.WalletRepository) {
				m2.On("GetHouseReservationsByID", 1).Return(&dtos.HouseReservationsDTO{
					UserID:      1,
					HouseCityID: 2,
				}, nil)
				m.On("ValidateOnlyOnePickupReservation", 1).Return(nil)
				m3.On("GetUserDetails", 1).Return(&dtos.UserDetailsDTO{
					Balance: 500000,
				}, nil)
				m.On("GetPickupStatusByName", "Pending Admin").Return(&models.PickupStatus{
					ID: 1,
				}, nil)
				m.On("CreatePickup", &models.Pickup{
					ReservationID:  1,
					PickupStatusID: 1,
					CityID:         1,
					PickupPrice:    300000,
				}).Return(&models.Pickup{}, nil)
				m4.On("SubtractWalletBalance", 1, 300000).Return(nil, nil)
			},
			args: args{
				reservationID: 1,
				userID:        1,
				request: &dtos.RequestPickupRequestDTO{
					CityID: 1,
				},
			},
			want:    &models.Pickup{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &pickupService{
				pickupRepository:      tt.pickupRepository,
				reservationRepository: tt.reservationRepository,
				userRepository:        tt.userRepository,
				walletRepository:      tt.walletRepository,
			}

			tt.mock(tt.pickupRepository, tt.reservationRepository, tt.userRepository, tt.walletRepository)

			got, err := s.RequestPickup(tt.args.reservationID, tt.args.userID, tt.args.request)

			if (err != nil) != tt.wantErr {
				t.Errorf("pickupService.RequestPickup() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("pickupService.RequestPickup() = %v, want %v", got, tt.want)
			}

		})
	}
}

func Test_pickupService_GetPickupDetailsByReservation(t *testing.T) {
	type args struct {
		reservationID int
		userID        int
	}
	tests := []struct {
		name                  string
		pickupRepository      *mocks.PickupRepository
		reservationRepository *mocks.ReservationRepository
		mock                  func(m *mocks.PickupRepository, m2 *mocks.ReservationRepository)
		args                  args
		want                  *dtos.GetPickupDetailsByReservationResponseDTO
		wantErr               bool
	}{
		{
			name:                  "Failed - GetHouseReservationsByID",
			pickupRepository:      mocks.NewPickupRepository(t),
			reservationRepository: mocks.NewReservationRepository(t),
			mock: func(m *mocks.PickupRepository, m2 *mocks.ReservationRepository) {
				m2.On("GetHouseReservationsByID", 1).Return(nil, errors.New("error"))
			},
			args: args{
				reservationID: 1,
				userID:        1,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name:                  "Failed - User is not the Customer",
			pickupRepository:      mocks.NewPickupRepository(t),
			reservationRepository: mocks.NewReservationRepository(t),
			mock: func(m *mocks.PickupRepository, m2 *mocks.ReservationRepository) {
				m2.On("GetHouseReservationsByID", 1).Return(&dtos.HouseReservationsDTO{
					UserID: 2,
				}, nil)
			},
			args: args{
				reservationID: 1,
				userID:        1,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name:                  "Failed - GetPickupDetailsByReservationID",
			pickupRepository:      mocks.NewPickupRepository(t),
			reservationRepository: mocks.NewReservationRepository(t),
			mock: func(m *mocks.PickupRepository, m2 *mocks.ReservationRepository) {
				m2.On("GetHouseReservationsByID", 1).Return(&dtos.HouseReservationsDTO{
					UserID: 1,
				}, nil)
				m.On("GetPickupDetailsByReservationID", 1).Return(nil, errors.New("error"))
			},
			args: args{
				reservationID: 1,
				userID:        1,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name:                  "Failed - GetPickupDetailsByReservationID",
			pickupRepository:      mocks.NewPickupRepository(t),
			reservationRepository: mocks.NewReservationRepository(t),
			mock: func(m *mocks.PickupRepository, m2 *mocks.ReservationRepository) {
				m2.On("GetHouseReservationsByID", 1).Return(&dtos.HouseReservationsDTO{
					UserID: 1,
				}, nil)
				m.On("GetPickupDetailsByReservationID", 1).Return(&dtos.GetPickupDetailsByReservationResponseDTO{}, nil)
			},
			args: args{
				reservationID: 1,
				userID:        1,
			},
			want:    &dtos.GetPickupDetailsByReservationResponseDTO{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &pickupService{
				pickupRepository:      tt.pickupRepository,
				reservationRepository: tt.reservationRepository,
			}

			tt.mock(tt.pickupRepository, tt.reservationRepository)

			got, err := s.GetPickupDetailsByReservation(tt.args.reservationID, tt.args.userID)

			if (err != nil) != tt.wantErr {
				t.Errorf("pickupService.GetPickupDetailsByReservation() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("pickupService.GetPickupDetailsByReservation() = %v, want %v", got, tt.want)
			}

		})
	}
}

func Test_pickupService_AdminUpdatePickupStatus(t *testing.T) {
	type args struct {
		pickupStatusID int
		pickupID       int
		userID         int
	}
	tests := []struct {
		name                  string
		pickupRepository      *mocks.PickupRepository
		reservationRepository *mocks.ReservationRepository
		userRepository        *mocks.UserRepository
		mock                  func(m *mocks.PickupRepository, m2 *mocks.ReservationRepository, m3 *mocks.UserRepository)
		args                  args
		want                  *models.Pickup
		wantErr               bool
	}{
		{
			name:                  "Failed - GetUserDetails",
			pickupRepository:      mocks.NewPickupRepository(t),
			reservationRepository: mocks.NewReservationRepository(t),
			userRepository:        mocks.NewUserRepository(t),
			mock: func(m *mocks.PickupRepository, m2 *mocks.ReservationRepository, m3 *mocks.UserRepository) {
				m3.On("GetUserDetails", 1).Return(nil, errors.New("error"))
			},
			args: args{
				pickupStatusID: 2,
				pickupID:       1,
				userID:         1,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name:                  "Failed - User is not Admin",
			pickupRepository:      mocks.NewPickupRepository(t),
			reservationRepository: mocks.NewReservationRepository(t),
			userRepository:        mocks.NewUserRepository(t),
			mock: func(m *mocks.PickupRepository, m2 *mocks.ReservationRepository, m3 *mocks.UserRepository) {
				m3.On("GetUserDetails", 1).Return(&dtos.UserDetailsDTO{
					Role: "user",
				}, nil)
			},
			args: args{
				pickupStatusID: 2,
				pickupID:       1,
				userID:         1,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name:                  "Failed - GetPickupByID",
			pickupRepository:      mocks.NewPickupRepository(t),
			reservationRepository: mocks.NewReservationRepository(t),
			userRepository:        mocks.NewUserRepository(t),
			mock: func(m *mocks.PickupRepository, m2 *mocks.ReservationRepository, m3 *mocks.UserRepository) {
				m3.On("GetUserDetails", 1).Return(&dtos.UserDetailsDTO{
					Role: "admin",
				}, nil)
				m.On("GetPickupByID", 1).Return(nil, errors.New("error"))
			},
			args: args{
				pickupStatusID: 2,
				pickupID:       1,
				userID:         1,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name:                  "Failed - GetPickupStatusByID",
			pickupRepository:      mocks.NewPickupRepository(t),
			reservationRepository: mocks.NewReservationRepository(t),
			userRepository:        mocks.NewUserRepository(t),
			mock: func(m *mocks.PickupRepository, m2 *mocks.ReservationRepository, m3 *mocks.UserRepository) {
				m3.On("GetUserDetails", 1).Return(&dtos.UserDetailsDTO{
					Role: "admin",
				}, nil)
				m.On("GetPickupByID", 1).Return(&models.Pickup{}, nil)
				m.On("GetPickupStatusByID", 2).Return(nil, errors.New("error"))
			},
			args: args{
				pickupStatusID: 2,
				pickupID:       1,
				userID:         1,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name:                  "Failed - GetHouseReservationsByID",
			pickupRepository:      mocks.NewPickupRepository(t),
			reservationRepository: mocks.NewReservationRepository(t),
			userRepository:        mocks.NewUserRepository(t),
			mock: func(m *mocks.PickupRepository, m2 *mocks.ReservationRepository, m3 *mocks.UserRepository) {
				m3.On("GetUserDetails", 1).Return(&dtos.UserDetailsDTO{
					Role: "admin",
				}, nil)
				m.On("GetPickupByID", 1).Return(&models.Pickup{
					ReservationID: 1,
				}, nil)
				m.On("GetPickupStatusByID", 2).Return(&models.PickupStatus{
					Status: "on the way to pickup",
				}, nil)
				m2.On("GetHouseReservationsByID", 1).Return(nil, errors.New("error"))
			},
			args: args{
				pickupStatusID: 2,
				pickupID:       1,
				userID:         1,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name:                  "Failed - Cannot Update Pickup Status because Check In Date is not Today",
			pickupRepository:      mocks.NewPickupRepository(t),
			reservationRepository: mocks.NewReservationRepository(t),
			userRepository:        mocks.NewUserRepository(t),
			mock: func(m *mocks.PickupRepository, m2 *mocks.ReservationRepository, m3 *mocks.UserRepository) {
				m3.On("GetUserDetails", 1).Return(&dtos.UserDetailsDTO{
					Role: "admin",
				}, nil)
				m.On("GetPickupByID", 1).Return(&models.Pickup{
					ReservationID: 1,
				}, nil)
				m.On("GetPickupStatusByID", 3).Return(&models.PickupStatus{
					Status: "On the Way to Pickup",
				}, nil)
				m2.On("GetHouseReservationsByID", 1).Return(&dtos.HouseReservationsDTO{
					CheckInDate: time.Now().AddDate(0, 0, -1),
				}, nil)
			},
			args: args{
				pickupStatusID: 3,
				pickupID:       1,
				userID:         1,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name:                  "Failed - UpdatePickupStatusByID",
			pickupRepository:      mocks.NewPickupRepository(t),
			reservationRepository: mocks.NewReservationRepository(t),
			userRepository:        mocks.NewUserRepository(t),
			mock: func(m *mocks.PickupRepository, m2 *mocks.ReservationRepository, m3 *mocks.UserRepository) {
				m3.On("GetUserDetails", 1).Return(&dtos.UserDetailsDTO{
					Role: "admin",
				}, nil)
				m.On("GetPickupByID", 1).Return(&models.Pickup{
					ReservationID: 1,
				}, nil)
				m.On("GetPickupStatusByID", 3).Return(&models.PickupStatus{
					Status: "On the Way to Pickup",
				}, nil)
				m2.On("GetHouseReservationsByID", 1).Return(&dtos.HouseReservationsDTO{
					CheckInDate: time.Now(),
				}, nil)
				m.On("UpdatePickupStatusByID", 3, 1).Return(nil, errors.New("error"))
			},
			args: args{
				pickupStatusID: 3,
				pickupID:       1,
				userID:         1,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name:                  "Success",
			pickupRepository:      mocks.NewPickupRepository(t),
			reservationRepository: mocks.NewReservationRepository(t),
			userRepository:        mocks.NewUserRepository(t),
			mock: func(m *mocks.PickupRepository, m2 *mocks.ReservationRepository, m3 *mocks.UserRepository) {
				m3.On("GetUserDetails", 1).Return(&dtos.UserDetailsDTO{
					Role: "admin",
				}, nil)
				m.On("GetPickupByID", 1).Return(&models.Pickup{
					ReservationID: 1,
				}, nil)
				m.On("GetPickupStatusByID", 2).Return(&models.PickupStatus{
					Status: "Awaiting Check-in Date",
				}, nil)
				m.On("UpdatePickupStatusByID", 2, 1).Return(&models.Pickup{}, nil)
			},
			args: args{
				pickupStatusID: 2,
				pickupID:       1,
				userID:         1,
			},
			want:    &models.Pickup{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &pickupService{
				pickupRepository:      tt.pickupRepository,
				reservationRepository: tt.reservationRepository,
				userRepository:        tt.userRepository,
			}

			tt.mock(tt.pickupRepository, tt.reservationRepository, tt.userRepository)

			got, err := s.AdminUpdatePickupStatus(tt.args.pickupStatusID, tt.args.pickupID, tt.args.userID)

			if (err != nil) != tt.wantErr {
				t.Errorf("pickupService.AdminUpdatePickupStatus() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("pickupService.AdminUpdatePickupStatus() = %v, want %v", got, tt.want)
			}

		})
	}
}

func Test_pickupService_GetAllPickupStatuses(t *testing.T) {
	type args struct{}
	tests := []struct {
		name             string
		pickupRepository *mocks.PickupRepository
		mock             func(m *mocks.PickupRepository)
		args             args
		want             []*dtos.PickupStatusesDTO
		wantErr          bool
	}{
		{
			name:             "Failed - GetPickupStatuses",
			pickupRepository: mocks.NewPickupRepository(t),
			mock: func(m *mocks.PickupRepository) {
				m.On("GetPickupStatuses").Return(nil, errors.New("error"))
			},
			args:    args{},
			want:    nil,
			wantErr: true,
		},
		{
			name:             "Success",
			pickupRepository: mocks.NewPickupRepository(t),
			mock: func(m *mocks.PickupRepository) {
				m.On("GetPickupStatuses").Return([]*dtos.PickupStatusesDTO{}, nil)
			},
			args:    args{},
			want:    []*dtos.PickupStatusesDTO{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &pickupService{
				pickupRepository: tt.pickupRepository,
			}

			tt.mock(tt.pickupRepository)

			got, err := s.GetAllPickupStatuses()

			if (err != nil) != tt.wantErr {
				t.Errorf("pickupService.GetAllPickupStatuses() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("pickupService.GetAllPickupStatuses() = %v, want %v", got, tt.want)
			}

		})
	}
}

func Test_pickupService_GetAllPickupDetails(t *testing.T) {
	type args struct {
		queryParams map[string]interface{}
	}
	tests := []struct {
		name             string
		pickupRepository *mocks.PickupRepository
		mock             func(m *mocks.PickupRepository)
		args             args
		want             []*dtos.PickupDetailsDTO
		want2            *helpers.Pagination
		wantErr          bool
	}{
		{
			name:             "Failed - GetPickupDetails",
			pickupRepository: mocks.NewPickupRepository(t),
			mock: func(m *mocks.PickupRepository) {
				m.On("GetPickupDetails", map[string]interface{}{}).Return(nil, nil, errors.New("error"))
			},
			args: args{
				queryParams: map[string]interface{}{},
			},
			want:    nil,
			want2:   nil,
			wantErr: true,
		},
		{
			name:             "Success",
			pickupRepository: mocks.NewPickupRepository(t),
			mock: func(m *mocks.PickupRepository) {
				m.On("GetPickupDetails", map[string]interface{}{}).Return([]*dtos.PickupDetailsDTO{}, &helpers.Pagination{}, nil)
			},
			args: args{
				queryParams: map[string]interface{}{},
			},
			want:    []*dtos.PickupDetailsDTO{},
			want2:   &helpers.Pagination{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &pickupService{
				pickupRepository: tt.pickupRepository,
			}

			tt.mock(tt.pickupRepository)

			got, got2, err := s.GetAllPickupDetails(tt.args.queryParams)

			if (err != nil) != tt.wantErr {
				t.Errorf("pickupService.GetAllPickupDetails() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("pickupService.GetAllPickupDetails() = %v, want %v", got, tt.want)
			}

			if !reflect.DeepEqual(got2, tt.want2) {
				t.Errorf("pickupService.GetAllPickupDetails() = %v, want %v", got2, tt.want2)
			}

		})
	}
}
