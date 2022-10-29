package services

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"testing"
	"time"

	"git.garena.com/sea-labs-id/batch-02/natanael-benediktus/final-project-backend/internal/dtos"
	"git.garena.com/sea-labs-id/batch-02/natanael-benediktus/final-project-backend/internal/helpers"
	"git.garena.com/sea-labs-id/batch-02/natanael-benediktus/final-project-backend/internal/mocks"
	"git.garena.com/sea-labs-id/batch-02/natanael-benediktus/final-project-backend/internal/models"
	"github.com/stretchr/testify/assert"
)

func Test_reservationService_NewReservationService(t *testing.T) {
	reservationRepository := mocks.NewReservationRepository(t)
	houseRepository := mocks.NewHouseRepository(t)
	userRepository := mocks.NewUserRepository(t)
	walletRepository := mocks.NewWalletRepository(t)
	reservationService := NewReservationService(
		&RSConfig{
			ReservationRepository: reservationRepository,
			HouseRepository:       houseRepository,
			UserRepository:        userRepository,
			WalletRepository:      walletRepository,
		},
	)

	assert.NotNil(t, reservationService)
}

func Test_reservationService_AddReservation(t *testing.T) {
	type args struct {
		userID  int
		request *dtos.AddReservationRequestDTO
	}
	successCheckInDate, _ := time.Parse("2006-01-02", time.Now().Format("2006-01-02"))
	successCheckOutDate, _ := time.Parse("2006-01-02", time.Now().AddDate(0, 0, 1).Format("2006-01-02"))
	tests := []struct {
		name                  string
		reservationRepository *mocks.ReservationRepository
		houseRepository       *mocks.HouseRepository
		userRepository        *mocks.UserRepository
		walletRepository      *mocks.WalletRepository
		mock                  func(m *mocks.ReservationRepository, m2 *mocks.HouseRepository, m3 *mocks.UserRepository, m4 *mocks.WalletRepository)
		args                  args
		want                  *models.Reservation
		wantErr               bool
	}{
		{
			name:                  "Failed - GetHouseByHouseID",
			reservationRepository: mocks.NewReservationRepository(t),
			houseRepository:       mocks.NewHouseRepository(t),
			userRepository:        mocks.NewUserRepository(t),
			walletRepository:      mocks.NewWalletRepository(t),
			mock: func(m *mocks.ReservationRepository, m2 *mocks.HouseRepository, m3 *mocks.UserRepository, m4 *mocks.WalletRepository) {
				m2.On("GetHouseByHouseID", 1).Return(nil, errors.New("error"))
			},
			args: args{
				userID: 1,
				request: &dtos.AddReservationRequestDTO{
					HouseID: 1,
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name:                  "Failed - User the Host",
			reservationRepository: mocks.NewReservationRepository(t),
			houseRepository:       mocks.NewHouseRepository(t),
			userRepository:        mocks.NewUserRepository(t),
			walletRepository:      mocks.NewWalletRepository(t),
			mock: func(m *mocks.ReservationRepository, m2 *mocks.HouseRepository, m3 *mocks.UserRepository, m4 *mocks.WalletRepository) {
				m2.On("GetHouseByHouseID", 1).Return(&dtos.HouseDetailsDTO{
					HostID: 1,
				}, nil)
			},
			args: args{
				userID: 1,
				request: &dtos.AddReservationRequestDTO{
					HouseID: 1,
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name:                  "Failed - Check In Date is the same as the Check Out Date",
			reservationRepository: mocks.NewReservationRepository(t),
			houseRepository:       mocks.NewHouseRepository(t),
			userRepository:        mocks.NewUserRepository(t),
			walletRepository:      mocks.NewWalletRepository(t),
			mock: func(m *mocks.ReservationRepository, m2 *mocks.HouseRepository, m3 *mocks.UserRepository, m4 *mocks.WalletRepository) {
				m2.On("GetHouseByHouseID", 1).Return(&dtos.HouseDetailsDTO{
					HostID: 2,
				}, nil)
			},
			args: args{
				userID: 1,
				request: &dtos.AddReservationRequestDTO{
					HouseID:      1,
					CheckInDate:  time.Now().Format("2006-01-02"),
					CheckOutDate: time.Now().Format("2006-01-02"),
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name:                  "Failed - Check In Date is later than the Check Out Date",
			reservationRepository: mocks.NewReservationRepository(t),
			houseRepository:       mocks.NewHouseRepository(t),
			userRepository:        mocks.NewUserRepository(t),
			walletRepository:      mocks.NewWalletRepository(t),
			mock: func(m *mocks.ReservationRepository, m2 *mocks.HouseRepository, m3 *mocks.UserRepository, m4 *mocks.WalletRepository) {
				m2.On("GetHouseByHouseID", 1).Return(&dtos.HouseDetailsDTO{
					HostID: 2,
				}, nil)
			},
			args: args{
				userID: 1,
				request: &dtos.AddReservationRequestDTO{
					HouseID:      1,
					CheckInDate:  time.Now().AddDate(0, 0, 1).Format("2006-01-02"),
					CheckOutDate: time.Now().Format("2006-01-02"),
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name:                  "Failed - Check In Date is earlier than today",
			reservationRepository: mocks.NewReservationRepository(t),
			houseRepository:       mocks.NewHouseRepository(t),
			userRepository:        mocks.NewUserRepository(t),
			walletRepository:      mocks.NewWalletRepository(t),
			mock: func(m *mocks.ReservationRepository, m2 *mocks.HouseRepository, m3 *mocks.UserRepository, m4 *mocks.WalletRepository) {
				m2.On("GetHouseByHouseID", 1).Return(&dtos.HouseDetailsDTO{
					HostID: 2,
				}, nil)
			},
			args: args{
				userID: 1,
				request: &dtos.AddReservationRequestDTO{
					HouseID:      1,
					CheckInDate:  time.Now().AddDate(0, 0, -1).Format("2006-01-02"),
					CheckOutDate: time.Now().Format("2006-01-02"),
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name:                  "Failed - ValidateReservation",
			reservationRepository: mocks.NewReservationRepository(t),
			houseRepository:       mocks.NewHouseRepository(t),
			userRepository:        mocks.NewUserRepository(t),
			walletRepository:      mocks.NewWalletRepository(t),
			mock: func(m *mocks.ReservationRepository, m2 *mocks.HouseRepository, m3 *mocks.UserRepository, m4 *mocks.WalletRepository) {
				m2.On("GetHouseByHouseID", 1).Return(&dtos.HouseDetailsDTO{
					HostID: 2,
				}, nil)
				m.On("ValidateReservation", &dtos.AddReservationRequestDTO{
					HouseID:      1,
					CheckInDate:  time.Now().Format("2006-01-02"),
					CheckOutDate: time.Now().AddDate(0, 0, 1).Format("2006-01-02"),
				}).Return(errors.New("error"))
			},
			args: args{
				userID: 1,
				request: &dtos.AddReservationRequestDTO{
					HouseID:      1,
					CheckInDate:  time.Now().Format("2006-01-02"),
					CheckOutDate: time.Now().AddDate(0, 0, 1).Format("2006-01-02"),
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name:                  "Failed - GetUserDetails",
			reservationRepository: mocks.NewReservationRepository(t),
			houseRepository:       mocks.NewHouseRepository(t),
			userRepository:        mocks.NewUserRepository(t),
			walletRepository:      mocks.NewWalletRepository(t),
			mock: func(m *mocks.ReservationRepository, m2 *mocks.HouseRepository, m3 *mocks.UserRepository, m4 *mocks.WalletRepository) {
				m2.On("GetHouseByHouseID", 1).Return(&dtos.HouseDetailsDTO{
					HostID: 2,
				}, nil)
				m.On("ValidateReservation", &dtos.AddReservationRequestDTO{
					HouseID:      1,
					CheckInDate:  time.Now().Format("2006-01-02"),
					CheckOutDate: time.Now().AddDate(0, 0, 1).Format("2006-01-02"),
				}).Return(nil)
				m3.On("GetUserDetails", 1).Return(nil, errors.New("error"))
			},
			args: args{
				userID: 1,
				request: &dtos.AddReservationRequestDTO{
					HouseID:      1,
					CheckInDate:  time.Now().Format("2006-01-02"),
					CheckOutDate: time.Now().AddDate(0, 0, 1).Format("2006-01-02"),
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name:                  "Failed - Insufficient Balance",
			reservationRepository: mocks.NewReservationRepository(t),
			houseRepository:       mocks.NewHouseRepository(t),
			userRepository:        mocks.NewUserRepository(t),
			walletRepository:      mocks.NewWalletRepository(t),
			mock: func(m *mocks.ReservationRepository, m2 *mocks.HouseRepository, m3 *mocks.UserRepository, m4 *mocks.WalletRepository) {
				m2.On("GetHouseByHouseID", 1).Return(&dtos.HouseDetailsDTO{
					HostID:     2,
					HousePrice: 100000,
				}, nil)
				m.On("ValidateReservation", &dtos.AddReservationRequestDTO{
					HouseID:      1,
					CheckInDate:  time.Now().Format("2006-01-02"),
					CheckOutDate: time.Now().AddDate(0, 0, 1).Format("2006-01-02"),
				}).Return(nil)
				m3.On("GetUserDetails", 1).Return(&dtos.UserDetailsDTO{
					Balance: 50000,
				}, nil)
			},
			args: args{
				userID: 1,
				request: &dtos.AddReservationRequestDTO{
					HouseID:      1,
					CheckInDate:  time.Now().Format("2006-01-02"),
					CheckOutDate: time.Now().AddDate(0, 0, 1).Format("2006-01-02"),
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name:                  "Failed - GetLastReservationID",
			reservationRepository: mocks.NewReservationRepository(t),
			houseRepository:       mocks.NewHouseRepository(t),
			userRepository:        mocks.NewUserRepository(t),
			walletRepository:      mocks.NewWalletRepository(t),
			mock: func(m *mocks.ReservationRepository, m2 *mocks.HouseRepository, m3 *mocks.UserRepository, m4 *mocks.WalletRepository) {
				m2.On("GetHouseByHouseID", 1).Return(&dtos.HouseDetailsDTO{
					HostID:     2,
					HousePrice: 100000,
				}, nil)
				m.On("ValidateReservation", &dtos.AddReservationRequestDTO{
					HouseID:      1,
					CheckInDate:  time.Now().Format("2006-01-02"),
					CheckOutDate: time.Now().AddDate(0, 0, 1).Format("2006-01-02"),
				}).Return(nil)
				m3.On("GetUserDetails", 1).Return(&dtos.UserDetailsDTO{
					Balance: 500000,
				}, nil)
				m.On("GetLastReservationID").Return(0, errors.New("error"))
			},
			args: args{
				userID: 1,
				request: &dtos.AddReservationRequestDTO{
					HouseID:      1,
					CheckInDate:  time.Now().Format("2006-01-02"),
					CheckOutDate: time.Now().AddDate(0, 0, 1).Format("2006-01-02"),
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name:                  "Failed - CreateReservation",
			reservationRepository: mocks.NewReservationRepository(t),
			houseRepository:       mocks.NewHouseRepository(t),
			userRepository:        mocks.NewUserRepository(t),
			walletRepository:      mocks.NewWalletRepository(t),
			mock: func(m *mocks.ReservationRepository, m2 *mocks.HouseRepository, m3 *mocks.UserRepository, m4 *mocks.WalletRepository) {
				m2.On("GetHouseByHouseID", 1).Return(&dtos.HouseDetailsDTO{
					HostID:     2,
					HousePrice: 100000,
				}, nil)
				m.On("ValidateReservation", &dtos.AddReservationRequestDTO{
					HouseID:      1,
					CheckInDate:  time.Now().Format("2006-01-02"),
					CheckOutDate: time.Now().AddDate(0, 0, 1).Format("2006-01-02"),
				}).Return(nil)
				m3.On("GetUserDetails", 1).Return(&dtos.UserDetailsDTO{
					Balance: 500000,
				}, nil)
				m.On("GetLastReservationID").Return(1, nil)
				checkInDate, _ := time.Parse("2006-01-02", time.Now().Format("2006-01-02"))
				checkOutDate, _ := time.Parse("2006-01-02", time.Now().AddDate(0, 0, 1).Format("2006-01-02"))
				m.On("CreateReservation", 1, &models.Reservation{
					HouseID: 1,
					UserID:  1,
					ReservationCode: fmt.Sprintf("RC%s%d%s",
						strings.ReplaceAll(time.Now().Format("2006-01-02"), "-", ""),
						2,
						strings.ReplaceAll(time.Now().AddDate(0, 0, 1).Format("2006-01-02"), "-", "")),
					CheckInDate:  checkInDate,
					CheckOutDate: checkOutDate,
					TotalPrice:   100000,
				}).Return(nil, errors.New("error"))
			},
			args: args{
				userID: 1,
				request: &dtos.AddReservationRequestDTO{
					HouseID:      1,
					CheckInDate:  time.Now().Format("2006-01-02"),
					CheckOutDate: time.Now().AddDate(0, 0, 1).Format("2006-01-02"),
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name:                  "Failed - SubtractWalletBalance",
			reservationRepository: mocks.NewReservationRepository(t),
			houseRepository:       mocks.NewHouseRepository(t),
			userRepository:        mocks.NewUserRepository(t),
			walletRepository:      mocks.NewWalletRepository(t),
			mock: func(m *mocks.ReservationRepository, m2 *mocks.HouseRepository, m3 *mocks.UserRepository, m4 *mocks.WalletRepository) {
				m2.On("GetHouseByHouseID", 1).Return(&dtos.HouseDetailsDTO{
					HostID:     2,
					HousePrice: 100000,
				}, nil)
				m.On("ValidateReservation", &dtos.AddReservationRequestDTO{
					HouseID:      1,
					CheckInDate:  time.Now().Format("2006-01-02"),
					CheckOutDate: time.Now().AddDate(0, 0, 1).Format("2006-01-02"),
				}).Return(nil)
				m3.On("GetUserDetails", 1).Return(&dtos.UserDetailsDTO{
					Balance: 500000,
				}, nil)
				m.On("GetLastReservationID").Return(1, nil)
				checkInDate, _ := time.Parse("2006-01-02", time.Now().Format("2006-01-02"))
				checkOutDate, _ := time.Parse("2006-01-02", time.Now().AddDate(0, 0, 1).Format("2006-01-02"))
				m.On("CreateReservation", 1, &models.Reservation{
					HouseID: 1,
					UserID:  1,
					ReservationCode: fmt.Sprintf("RC%s%d%s",
						strings.ReplaceAll(time.Now().Format("2006-01-02"), "-", ""),
						2,
						strings.ReplaceAll(time.Now().AddDate(0, 0, 1).Format("2006-01-02"), "-", "")),
					CheckInDate:  checkInDate,
					CheckOutDate: checkOutDate,
					TotalPrice:   100000,
				}).Return(&models.Reservation{}, nil)
				m4.On("SubtractWalletBalance", 1, 100000).Return(nil, errors.New("error"))
			},
			args: args{
				userID: 1,
				request: &dtos.AddReservationRequestDTO{
					HouseID:      1,
					CheckInDate:  time.Now().Format("2006-01-02"),
					CheckOutDate: time.Now().AddDate(0, 0, 1).Format("2006-01-02"),
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name:                  "Success",
			reservationRepository: mocks.NewReservationRepository(t),
			houseRepository:       mocks.NewHouseRepository(t),
			userRepository:        mocks.NewUserRepository(t),
			walletRepository:      mocks.NewWalletRepository(t),
			mock: func(m *mocks.ReservationRepository, m2 *mocks.HouseRepository, m3 *mocks.UserRepository, m4 *mocks.WalletRepository) {
				m2.On("GetHouseByHouseID", 1).Return(&dtos.HouseDetailsDTO{
					HostID:     2,
					HousePrice: 100000,
				}, nil)
				m.On("ValidateReservation", &dtos.AddReservationRequestDTO{
					HouseID:      1,
					CheckInDate:  time.Now().Format("2006-01-02"),
					CheckOutDate: time.Now().AddDate(0, 0, 1).Format("2006-01-02"),
				}).Return(nil)
				m3.On("GetUserDetails", 1).Return(&dtos.UserDetailsDTO{
					Balance: 500000,
				}, nil)
				m.On("GetLastReservationID").Return(1, nil)
				checkInDate, _ := time.Parse("2006-01-02", time.Now().Format("2006-01-02"))
				checkOutDate, _ := time.Parse("2006-01-02", time.Now().AddDate(0, 0, 1).Format("2006-01-02"))
				m.On("CreateReservation", 1, &models.Reservation{
					HouseID: 1,
					UserID:  1,
					ReservationCode: fmt.Sprintf("RC%s%d%s",
						strings.ReplaceAll(time.Now().Format("2006-01-02"), "-", ""),
						2,
						strings.ReplaceAll(time.Now().AddDate(0, 0, 1).Format("2006-01-02"), "-", "")),
					CheckInDate:  checkInDate,
					CheckOutDate: checkOutDate,
					TotalPrice:   100000,
				}).Return(&models.Reservation{
					HouseID: 1,
					UserID:  1,
					ReservationCode: fmt.Sprintf("RC%s%d%s",
						strings.ReplaceAll(time.Now().Format("2006-01-02"), "-", ""),
						2,
						strings.ReplaceAll(time.Now().AddDate(0, 0, 1).Format("2006-01-02"), "-", "")),
					CheckInDate:  checkInDate,
					CheckOutDate: checkOutDate,
					TotalPrice:   100000,
				}, nil)
				m4.On("SubtractWalletBalance", 1, 100000).Return(nil, nil)
			},
			args: args{
				userID: 1,
				request: &dtos.AddReservationRequestDTO{
					HouseID:      1,
					CheckInDate:  time.Now().Format("2006-01-02"),
					CheckOutDate: time.Now().AddDate(0, 0, 1).Format("2006-01-02"),
				},
			},
			want: &models.Reservation{
				HouseID: 1,
				UserID:  1,
				ReservationCode: fmt.Sprintf("RC%s%d%s",
					strings.ReplaceAll(time.Now().Format("2006-01-02"), "-", ""),
					2,
					strings.ReplaceAll(time.Now().AddDate(0, 0, 1).Format("2006-01-02"), "-", "")),
				CheckInDate:  successCheckInDate,
				CheckOutDate: successCheckOutDate,
				TotalPrice:   100000,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &reservationService{
				reservationRepository: tt.reservationRepository,
				houseRepository:       tt.houseRepository,
				userRepository:        tt.userRepository,
				walletRepository:      tt.walletRepository,
			}

			tt.mock(tt.reservationRepository, tt.houseRepository, tt.userRepository, tt.walletRepository)

			got, err := s.AddReservation(tt.args.userID, tt.args.request)

			if (err != nil) != tt.wantErr {
				t.Errorf("reservationService.AddReservation() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reservationService.AddReservation() = %v, want %v", got, tt.want)
			}

		})
	}
}

func Test_reservationService_GetHouseReservationsByUser(t *testing.T) {
	type args struct {
		userID      int
		queryParams map[string]interface{}
	}
	tests := []struct {
		name                  string
		reservationRepository *mocks.ReservationRepository
		mock                  func(m *mocks.ReservationRepository)
		args                  args
		want                  []*dtos.HouseReservationsDTO
		want2                 *helpers.Pagination
		wantErr               bool
	}{
		{
			name:                  "Failed - GetHouseReservationsByUserID",
			reservationRepository: mocks.NewReservationRepository(t),
			mock: func(m *mocks.ReservationRepository) {
				m.On("GetHouseReservationsByUserID", 1, map[string]interface{}{}).Return(nil, nil, errors.New("error"))
			},
			args: args{
				userID:      1,
				queryParams: map[string]interface{}{},
			},
			want:    nil,
			want2:   nil,
			wantErr: true,
		},
		{
			name:                  "Success",
			reservationRepository: mocks.NewReservationRepository(t),
			mock: func(m *mocks.ReservationRepository) {
				m.On("GetHouseReservationsByUserID", 1, map[string]interface{}{}).Return([]*dtos.HouseReservationsDTO{}, &helpers.Pagination{}, nil)
			},
			args: args{
				userID:      1,
				queryParams: map[string]interface{}{},
			},
			want:    []*dtos.HouseReservationsDTO{},
			want2:   &helpers.Pagination{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &reservationService{
				reservationRepository: tt.reservationRepository,
			}

			tt.mock(tt.reservationRepository)

			got, got2, err := s.GetHouseReservationsByUser(tt.args.userID, tt.args.queryParams)

			if (err != nil) != tt.wantErr {
				t.Errorf("reservationService.GetHouseReservationsByUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reservationService.GetHouseReservationsByUser() = %v, want %v", got, tt.want)
			}

			if !reflect.DeepEqual(got2, tt.want2) {
				t.Errorf("reservationService.GetHouseReservationsByUser() = %v, want %v", got2, tt.want2)
			}

		})
	}
}

func Test_reservationService_GetHouseReservationsByHost(t *testing.T) {
	type args struct {
		hostID      int
		queryParams map[string]interface{}
	}
	tests := []struct {
		name                  string
		reservationRepository *mocks.ReservationRepository
		mock                  func(m *mocks.ReservationRepository)
		args                  args
		want                  []*dtos.HouseReservationsDTO
		want2                 *helpers.Pagination
		wantErr               bool
	}{
		{
			name:                  "Failed - GetHouseReservationsByHostID",
			reservationRepository: mocks.NewReservationRepository(t),
			mock: func(m *mocks.ReservationRepository) {
				m.On("GetHouseReservationsByHostID", 1, map[string]interface{}{}).Return(nil, nil, errors.New("error"))
			},
			args: args{
				hostID:      1,
				queryParams: map[string]interface{}{},
			},
			want:    nil,
			want2:   nil,
			wantErr: true,
		},
		{
			name:                  "Success",
			reservationRepository: mocks.NewReservationRepository(t),
			mock: func(m *mocks.ReservationRepository) {
				m.On("GetHouseReservationsByHostID", 1, map[string]interface{}{}).Return([]*dtos.HouseReservationsDTO{}, &helpers.Pagination{}, nil)
			},
			args: args{
				hostID:      1,
				queryParams: map[string]interface{}{},
			},
			want:    []*dtos.HouseReservationsDTO{},
			want2:   &helpers.Pagination{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &reservationService{
				reservationRepository: tt.reservationRepository,
			}

			tt.mock(tt.reservationRepository)

			got, got2, err := s.GetHouseReservationsByHost(tt.args.hostID, tt.args.queryParams)

			if (err != nil) != tt.wantErr {
				t.Errorf("reservationService.GetHouseReservationsByHost() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reservationService.GetHouseReservationsByHost() = %v, want %v", got, tt.want)
			}

			if !reflect.DeepEqual(got2, tt.want2) {
				t.Errorf("reservationService.GetHouseReservationsByHost() = %v, want %v", got2, tt.want2)
			}

		})
	}
}

func Test_reservationService_GetHouseReservationsByHouse(t *testing.T) {
	type args struct {
		houseID     int
		queryParams map[string]interface{}
	}
	tests := []struct {
		name                  string
		reservationRepository *mocks.ReservationRepository
		houseRepository       *mocks.HouseRepository
		mock                  func(m *mocks.ReservationRepository, m2 *mocks.HouseRepository)
		args                  args
		want                  []*dtos.HouseReservationsDTO
		wantErr               bool
	}{
		{
			name:                  "Failed - GetHouseByHouseID",
			reservationRepository: mocks.NewReservationRepository(t),
			houseRepository:       mocks.NewHouseRepository(t),
			mock: func(m *mocks.ReservationRepository, m2 *mocks.HouseRepository) {
				m2.On("GetHouseByHouseID", 1).Return(nil, errors.New("error"))
			},
			args: args{
				houseID:     1,
				queryParams: map[string]interface{}{},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name:                  "Failed - GetHouseReservationsByHouseID",
			reservationRepository: mocks.NewReservationRepository(t),
			houseRepository:       mocks.NewHouseRepository(t),
			mock: func(m *mocks.ReservationRepository, m2 *mocks.HouseRepository) {
				m2.On("GetHouseByHouseID", 1).Return(nil, nil)
				m.On("GetHouseReservationsByHouseID", 1, map[string]interface{}{}).Return(nil, errors.New("error"))
			},
			args: args{
				houseID:     1,
				queryParams: map[string]interface{}{},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name:                  "Success",
			reservationRepository: mocks.NewReservationRepository(t),
			houseRepository:       mocks.NewHouseRepository(t),
			mock: func(m *mocks.ReservationRepository, m2 *mocks.HouseRepository) {
				m2.On("GetHouseByHouseID", 1).Return(nil, nil)
				m.On("GetHouseReservationsByHouseID", 1, map[string]interface{}{}).Return([]*dtos.HouseReservationsDTO{}, nil)
			},
			args: args{
				houseID:     1,
				queryParams: map[string]interface{}{},
			},
			want:    []*dtos.HouseReservationsDTO{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &reservationService{
				reservationRepository: tt.reservationRepository,
				houseRepository:       tt.houseRepository,
			}

			tt.mock(tt.reservationRepository, tt.houseRepository)

			got, err := s.GetHouseReservationsByHouse(tt.args.houseID, tt.args.queryParams)

			if (err != nil) != tt.wantErr {
				t.Errorf("reservationService.GetHouseReservationsByHouse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reservationService.GetHouseReservationsByHouse() = %v, want %v", got, tt.want)
			}

		})
	}
}

func Test_reservationService_GetHouseReservationsByID(t *testing.T) {
	type args struct {
		reservationID int
		userID        int
	}
	tests := []struct {
		name                  string
		reservationRepository *mocks.ReservationRepository
		houseRepository       *mocks.HouseRepository
		userRepository        *mocks.UserRepository
		mock                  func(m *mocks.ReservationRepository, m2 *mocks.HouseRepository, m3 *mocks.UserRepository)
		args                  args
		want                  *dtos.HouseReservationsDTO
		wantErr               bool
	}{
		{
			name:                  "Failed - GetHouseReservationsByID",
			reservationRepository: mocks.NewReservationRepository(t),
			houseRepository:       mocks.NewHouseRepository(t),
			userRepository:        mocks.NewUserRepository(t),
			mock: func(m *mocks.ReservationRepository, m2 *mocks.HouseRepository, m3 *mocks.UserRepository) {
				m.On("GetHouseReservationsByID", 1).Return(nil, errors.New("error"))
			},
			args: args{
				reservationID: 1,
				userID:        1,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name:                  "Failed - GetUnscopedHouseByHouseID",
			reservationRepository: mocks.NewReservationRepository(t),
			houseRepository:       mocks.NewHouseRepository(t),
			userRepository:        mocks.NewUserRepository(t),
			mock: func(m *mocks.ReservationRepository, m2 *mocks.HouseRepository, m3 *mocks.UserRepository) {
				m.On("GetHouseReservationsByID", 1).Return(&dtos.HouseReservationsDTO{
					HouseID: 1,
				}, nil)
				m2.On("GetUnscopedHouseByHouseID", 1).Return(nil, errors.New("error"))
			},
			args: args{
				reservationID: 1,
				userID:        1,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name:                  "Failed - GetUserDetails",
			reservationRepository: mocks.NewReservationRepository(t),
			houseRepository:       mocks.NewHouseRepository(t),
			userRepository:        mocks.NewUserRepository(t),
			mock: func(m *mocks.ReservationRepository, m2 *mocks.HouseRepository, m3 *mocks.UserRepository) {
				m.On("GetHouseReservationsByID", 1).Return(&dtos.HouseReservationsDTO{
					HouseID: 1,
				}, nil)
				m2.On("GetUnscopedHouseByHouseID", 1).Return(&dtos.HouseDetailsDTO{}, nil)
				m3.On("GetUserDetails", 1).Return(nil, errors.New("error"))
			},
			args: args{
				reservationID: 1,
				userID:        1,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name:                  "Failed - User is not the host or the customer, and not an admin",
			reservationRepository: mocks.NewReservationRepository(t),
			houseRepository:       mocks.NewHouseRepository(t),
			userRepository:        mocks.NewUserRepository(t),
			mock: func(m *mocks.ReservationRepository, m2 *mocks.HouseRepository, m3 *mocks.UserRepository) {
				m.On("GetHouseReservationsByID", 1).Return(&dtos.HouseReservationsDTO{
					HouseID: 1,
					UserID:  2,
				}, nil)
				m2.On("GetUnscopedHouseByHouseID", 1).Return(&dtos.HouseDetailsDTO{
					HostID: 3,
				}, nil)
				m3.On("GetUserDetails", 1).Return(nil, errors.New("error"))
			},
			args: args{
				reservationID: 1,
				userID:        1,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name:                  "Failed - User is not the host or the customer, and not an admin",
			reservationRepository: mocks.NewReservationRepository(t),
			houseRepository:       mocks.NewHouseRepository(t),
			userRepository:        mocks.NewUserRepository(t),
			mock: func(m *mocks.ReservationRepository, m2 *mocks.HouseRepository, m3 *mocks.UserRepository) {
				m.On("GetHouseReservationsByID", 1).Return(&dtos.HouseReservationsDTO{
					HouseID: 1,
					UserID:  2,
				}, nil)
				m2.On("GetUnscopedHouseByHouseID", 1).Return(&dtos.HouseDetailsDTO{
					HostID: 3,
				}, nil)
				m3.On("GetUserDetails", 1).Return(&dtos.UserDetailsDTO{}, nil)
			},
			args: args{
				reservationID: 1,
				userID:        1,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name:                  "Success",
			reservationRepository: mocks.NewReservationRepository(t),
			houseRepository:       mocks.NewHouseRepository(t),
			userRepository:        mocks.NewUserRepository(t),
			mock: func(m *mocks.ReservationRepository, m2 *mocks.HouseRepository, m3 *mocks.UserRepository) {
				m.On("GetHouseReservationsByID", 1).Return(&dtos.HouseReservationsDTO{
					HouseID: 1,
					UserID:  1,
				}, nil)
				m2.On("GetUnscopedHouseByHouseID", 1).Return(&dtos.HouseDetailsDTO{
					HostID: 2,
				}, nil)
				m3.On("GetUserDetails", 1).Return(&dtos.UserDetailsDTO{}, nil)
			},
			args: args{
				reservationID: 1,
				userID:        1,
			},
			want: &dtos.HouseReservationsDTO{
				HouseID: 1,
				UserID:  1,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &reservationService{
				reservationRepository: tt.reservationRepository,
				houseRepository:       tt.houseRepository,
				userRepository:        tt.userRepository,
			}

			tt.mock(tt.reservationRepository, tt.houseRepository, tt.userRepository)

			got, err := s.GetHouseReservationsByID(tt.args.reservationID, tt.args.userID)

			if (err != nil) != tt.wantErr {
				t.Errorf("reservationService.GetHouseReservationsByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reservationService.GetHouseReservationsByID() = %v, want %v", got, tt.want)
			}
		})
	}
}
