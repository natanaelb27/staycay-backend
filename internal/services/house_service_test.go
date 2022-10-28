package services

import (
	"errors"
	"reflect"
	"testing"

	"git.garena.com/sea-labs-id/batch-02/natanael-benediktus/final-project-backend/internal/dtos"
	"git.garena.com/sea-labs-id/batch-02/natanael-benediktus/final-project-backend/internal/helpers"
	"git.garena.com/sea-labs-id/batch-02/natanael-benediktus/final-project-backend/internal/mocks"
	"git.garena.com/sea-labs-id/batch-02/natanael-benediktus/final-project-backend/internal/models"
	"github.com/stretchr/testify/assert"
)

func Test_houseService_NewHouseService(t *testing.T) {
	houseRepository := mocks.NewHouseRepository(t)
	userRepository := mocks.NewUserRepository(t)
	reservationRepository := mocks.NewReservationRepository(t)
	houseService := NewHouseService(
		&HSConfig{
			HouseRepository:       houseRepository,
			UserRepository:        userRepository,
			ReservationRepository: reservationRepository,
		},
	)

	assert.NotNil(t, houseService)
}

func Test_houseService_GetAllHouses(t *testing.T) {
	type args struct {
		queryParams map[string]interface{}
	}
	tests := []struct {
		name            string
		houseRepository *mocks.HouseRepository
		mock            func(m *mocks.HouseRepository)
		args            args
		want            []*dtos.HouseDetailsWithPhotoDTO
		want2           *helpers.Pagination
		wantErr         bool
	}{
		{
			name:            "Failed - GetHouse",
			houseRepository: mocks.NewHouseRepository(t),
			mock: func(m *mocks.HouseRepository) {
				m.On("GetHouse", map[string]interface{}{}).Return(nil, nil, errors.New("error"))
			},
			args: args{
				queryParams: map[string]interface{}{},
			},
			want:    nil,
			want2:   nil,
			wantErr: true,
		},
		{
			name:            "Failed - GetHousePhotoByHouseID",
			houseRepository: mocks.NewHouseRepository(t),
			mock: func(m *mocks.HouseRepository) {
				m.On("GetHouse", map[string]interface{}{}).Return([]*dtos.HouseDetailsDTO{
					{
						HouseID: 1,
					},
				}, &helpers.Pagination{}, nil)
				m.On("GetHousePhotoByHouseID", 1).Return(nil, errors.New("error"))
			},
			args: args{
				queryParams: map[string]interface{}{},
			},
			want:    nil,
			want2:   nil,
			wantErr: true,
		},
		{
			name:            "Success",
			houseRepository: mocks.NewHouseRepository(t),
			mock: func(m *mocks.HouseRepository) {
				m.On("GetHouse", map[string]interface{}{}).Return([]*dtos.HouseDetailsDTO{
					{
						HouseID: 1,
					},
				}, &helpers.Pagination{}, nil)
				m.On("GetHousePhotoByHouseID", 1).Return([]*dtos.HousePhotosDTO{}, nil)
			},
			args: args{
				queryParams: map[string]interface{}{},
			},
			want: []*dtos.HouseDetailsWithPhotoDTO{
				{
					HouseID:    1,
					HousePhoto: []*dtos.HousePhotosDTO{},
				},
			},
			want2:   &helpers.Pagination{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &houseService{
				houseRepository: tt.houseRepository,
			}

			tt.mock(tt.houseRepository)

			got, got2, err := s.GetAllHouses(tt.args.queryParams)

			if (err != nil) != tt.wantErr {
				t.Errorf("houseService.GetAllHouses() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("houseService.GetAllHouses() = %v, want %v", got, tt.want)
			}

			if !reflect.DeepEqual(got2, tt.want2) {
				t.Errorf("houseService.GetAllHouses() = %v, want2 %v", got2, tt.want2)
			}
		})
	}
}

func Test_houseService_GetHouseByID(t *testing.T) {
	type args struct {
		houseID     int
		queryParams map[string]interface{}
	}
	tests := []struct {
		name            string
		houseRepository *mocks.HouseRepository
		mock            func(m *mocks.HouseRepository)
		args            args
		want            *dtos.HouseDetailsWithPhotoDTO
		wantErr         bool
	}{
		{
			name:            "Failed - unscoped is true: GetUnscopedHouseByHouseID",
			houseRepository: mocks.NewHouseRepository(t),
			mock: func(m *mocks.HouseRepository) {
				m.On("GetUnscopedHouseByHouseID", 1).Return(nil, errors.New("error"))
			},
			args: args{
				houseID: 1,
				queryParams: map[string]interface{}{
					"unscoped": true,
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name:            "Failed - unscoped is true: GetUnscopedHousePhotoByHouseID",
			houseRepository: mocks.NewHouseRepository(t),
			mock: func(m *mocks.HouseRepository) {
				m.On("GetUnscopedHouseByHouseID", 1).Return(&dtos.HouseDetailsDTO{}, nil)
				m.On("GetUnscopedHousePhotoByHouseID", 1).Return(nil, errors.New("error"))
			},
			args: args{
				houseID: 1,
				queryParams: map[string]interface{}{
					"unscoped": true,
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name:            "Failed - unscoped is false: GetHouseByHouseID",
			houseRepository: mocks.NewHouseRepository(t),
			mock: func(m *mocks.HouseRepository) {
				m.On("GetHouseByHouseID", 1).Return(nil, errors.New("error"))
			},
			args: args{
				houseID: 1,
				queryParams: map[string]interface{}{
					"unscoped": false,
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name:            "Failed - unscoped is false: GetHousePhotoByHouseID",
			houseRepository: mocks.NewHouseRepository(t),
			mock: func(m *mocks.HouseRepository) {
				m.On("GetHouseByHouseID", 1).Return(&dtos.HouseDetailsDTO{}, nil)
				m.On("GetHousePhotoByHouseID", 1).Return(nil, errors.New("error"))
			},
			args: args{
				houseID: 1,
				queryParams: map[string]interface{}{
					"unscoped": false,
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name:            "Success - unscoped is true",
			houseRepository: mocks.NewHouseRepository(t),
			mock: func(m *mocks.HouseRepository) {
				m.On("GetUnscopedHouseByHouseID", 1).Return(&dtos.HouseDetailsDTO{}, nil)
				m.On("GetUnscopedHousePhotoByHouseID", 1).Return([]*dtos.HousePhotosDTO{}, nil)
			},
			args: args{
				houseID: 1,
				queryParams: map[string]interface{}{
					"unscoped": true,
				},
			},
			want: &dtos.HouseDetailsWithPhotoDTO{
				HousePhoto: []*dtos.HousePhotosDTO{},
			},
			wantErr: false,
		},
		{
			name:            "Success - unscoped is false",
			houseRepository: mocks.NewHouseRepository(t),
			mock: func(m *mocks.HouseRepository) {
				m.On("GetHouseByHouseID", 1).Return(&dtos.HouseDetailsDTO{}, nil)
				m.On("GetHousePhotoByHouseID", 1).Return([]*dtos.HousePhotosDTO{}, nil)
			},
			args: args{
				houseID: 1,
				queryParams: map[string]interface{}{
					"unscoped": false,
				},
			},
			want: &dtos.HouseDetailsWithPhotoDTO{
				HousePhoto: []*dtos.HousePhotosDTO{},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &houseService{
				houseRepository: tt.houseRepository,
			}

			tt.mock(tt.houseRepository)

			got, err := s.GetHouseByID(tt.args.houseID, tt.args.queryParams)

			if (err != nil) != tt.wantErr {
				t.Errorf("houseService.GetHouseByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("houseService.GetHouseByID() = %v, want %v", got, tt.want)
			}

		})
	}
}

func Test_houseService_GetHostHouses(t *testing.T) {
	type args struct {
		hostID      int
		queryParams map[string]interface{}
	}
	tests := []struct {
		name            string
		houseRepository *mocks.HouseRepository
		userRepository  *mocks.UserRepository
		mock            func(m *mocks.HouseRepository, m2 *mocks.UserRepository)
		args            args
		want            []*dtos.HouseDetailsWithPhotoDTO
		want2           *helpers.Pagination
		wantErr         bool
	}{
		{
			name:            "Failed - GetUserDetails",
			houseRepository: mocks.NewHouseRepository(t),
			userRepository:  mocks.NewUserRepository(t),
			mock: func(m *mocks.HouseRepository, m2 *mocks.UserRepository) {
				m2.On("GetUserDetails", 1).Return(nil, errors.New("error"))
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
			name:            "Failed - User is not a Host",
			houseRepository: mocks.NewHouseRepository(t),
			userRepository:  mocks.NewUserRepository(t),
			mock: func(m *mocks.HouseRepository, m2 *mocks.UserRepository) {
				m2.On("GetUserDetails", 1).Return(&dtos.UserDetailsDTO{
					Role: "user",
				}, nil)
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
			name:            "Failed - GetHouseByHostID",
			houseRepository: mocks.NewHouseRepository(t),
			userRepository:  mocks.NewUserRepository(t),
			mock: func(m *mocks.HouseRepository, m2 *mocks.UserRepository) {
				m2.On("GetUserDetails", 1).Return(&dtos.UserDetailsDTO{
					Role: "host",
				}, nil)
				m.On("GetHouseByHostID", 1, map[string]interface{}{}).Return(nil, nil, errors.New("error"))
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
			name:            "Failed - GetHousePhotoByHouseID",
			houseRepository: mocks.NewHouseRepository(t),
			userRepository:  mocks.NewUserRepository(t),
			mock: func(m *mocks.HouseRepository, m2 *mocks.UserRepository) {
				m2.On("GetUserDetails", 1).Return(&dtos.UserDetailsDTO{
					Role: "host",
				}, nil)
				m.On("GetHouseByHostID", 1, map[string]interface{}{}).Return([]*dtos.HouseDetailsDTO{
					{
						HouseID: 1,
					},
				}, &helpers.Pagination{}, nil)
				m.On("GetHousePhotoByHouseID", 1).Return(nil, errors.New("error"))
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
			name:            "Failed - GetHousePhotoByHouseID",
			houseRepository: mocks.NewHouseRepository(t),
			userRepository:  mocks.NewUserRepository(t),
			mock: func(m *mocks.HouseRepository, m2 *mocks.UserRepository) {
				m2.On("GetUserDetails", 1).Return(&dtos.UserDetailsDTO{
					Role: "host",
				}, nil)
				m.On("GetHouseByHostID", 1, map[string]interface{}{}).Return([]*dtos.HouseDetailsDTO{
					{
						HouseID: 1,
					},
				}, &helpers.Pagination{}, nil)
				m.On("GetHousePhotoByHouseID", 1).Return(nil, errors.New("error"))
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
			name:            "Success",
			houseRepository: mocks.NewHouseRepository(t),
			userRepository:  mocks.NewUserRepository(t),
			mock: func(m *mocks.HouseRepository, m2 *mocks.UserRepository) {
				m2.On("GetUserDetails", 1).Return(&dtos.UserDetailsDTO{
					Role: "host",
				}, nil)
				m.On("GetHouseByHostID", 1, map[string]interface{}{}).Return([]*dtos.HouseDetailsDTO{
					{
						HouseID: 1,
					},
				}, &helpers.Pagination{}, nil)
				m.On("GetHousePhotoByHouseID", 1).Return([]*dtos.HousePhotosDTO{}, nil)
			},
			args: args{
				hostID:      1,
				queryParams: map[string]interface{}{},
			},
			want: []*dtos.HouseDetailsWithPhotoDTO{
				{
					HouseID:    1,
					HousePhoto: []*dtos.HousePhotosDTO{},
				},
			},
			want2:   &helpers.Pagination{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &houseService{
				houseRepository: tt.houseRepository,
				userRepository:  tt.userRepository,
			}

			tt.mock(tt.houseRepository, tt.userRepository)

			got, got2, err := s.GetHostHouses(tt.args.hostID, tt.args.queryParams)

			if (err != nil) != tt.wantErr {
				t.Errorf("houseService.GetHostHouses() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("houseService.GetHostHouses() = %v, want %v", got, tt.want)
			}

			if !reflect.DeepEqual(got2, tt.want2) {
				t.Errorf("houseService.GetHostHouses() = %v, want2 %v", got2, tt.want2)
			}
		})
	}
}

func Test_houseService_GetAllHouseStatuses(t *testing.T) {
	type args struct{}
	tests := []struct {
		name            string
		houseRepository *mocks.HouseRepository
		mock            func(m *mocks.HouseRepository)
		args            args
		want            []*dtos.HouseStatusDTO
		wantErr         bool
	}{
		{
			name:            "Failed - GetAllHouseStatuses",
			houseRepository: mocks.NewHouseRepository(t),
			mock: func(m *mocks.HouseRepository) {
				m.On("GetHouseStatus").Return(nil, errors.New("error"))
			},
			args:    args{},
			want:    nil,
			wantErr: true,
		},
		{
			name:            "Success",
			houseRepository: mocks.NewHouseRepository(t),
			mock: func(m *mocks.HouseRepository) {
				m.On("GetHouseStatus").Return([]*dtos.HouseStatusDTO{}, nil)
			},
			args:    args{},
			want:    []*dtos.HouseStatusDTO{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &houseService{
				houseRepository: tt.houseRepository,
			}

			tt.mock(tt.houseRepository)

			got, err := s.GetAllHouseStatuses()

			if (err != nil) != tt.wantErr {
				t.Errorf("houseService.GetAllHouseStatuses() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("houseService.GetAllHouseStatuses() = %v, want %v", got, tt.want)
			}

		})
	}
}

func Test_houseService_AddNewHouse(t *testing.T) {
	type args struct {
		userID  int
		request *dtos.AddNewHouseRequestDTO
	}
	tests := []struct {
		name            string
		houseRepository *mocks.HouseRepository
		userRepository  *mocks.UserRepository
		mock            func(m *mocks.HouseRepository, m2 *mocks.UserRepository)
		args            args
		want            *models.House
		want2           []string
		wantErr         bool
	}{
		{
			name:            "Failed - GetUserDetails",
			houseRepository: mocks.NewHouseRepository(t),
			userRepository:  mocks.NewUserRepository(t),
			mock: func(m *mocks.HouseRepository, m2 *mocks.UserRepository) {
				m2.On("GetUserDetails", 1).Return(nil, errors.New("error"))
			},
			args: args{
				userID:  1,
				request: &dtos.AddNewHouseRequestDTO{},
			},
			want:    nil,
			want2:   nil,
			wantErr: true,
		},
		{
			name:            "Failed - User is not a Host",
			houseRepository: mocks.NewHouseRepository(t),
			userRepository:  mocks.NewUserRepository(t),
			mock: func(m *mocks.HouseRepository, m2 *mocks.UserRepository) {
				m2.On("GetUserDetails", 1).Return(&dtos.UserDetailsDTO{
					Role: "user",
				}, nil)
			},
			args: args{
				userID:  1,
				request: &dtos.AddNewHouseRequestDTO{},
			},
			want:    nil,
			want2:   nil,
			wantErr: true,
		},
		{
			name:            "Failed - GetHouseStatusByName",
			houseRepository: mocks.NewHouseRepository(t),
			userRepository:  mocks.NewUserRepository(t),
			mock: func(m *mocks.HouseRepository, m2 *mocks.UserRepository) {
				m2.On("GetUserDetails", 1).Return(&dtos.UserDetailsDTO{
					Role: "host",
				}, nil)
				m.On("GetHouseStatusByName", "Active").Return(nil, errors.New("error"))
			},
			args: args{
				userID:  1,
				request: &dtos.AddNewHouseRequestDTO{},
			},
			want:    nil,
			want2:   nil,
			wantErr: true,
		},
		{
			name:            "Failed - CreateNewHouse",
			houseRepository: mocks.NewHouseRepository(t),
			userRepository:  mocks.NewUserRepository(t),
			mock: func(m *mocks.HouseRepository, m2 *mocks.UserRepository) {
				m2.On("GetUserDetails", 1).Return(&dtos.UserDetailsDTO{
					Role: "host",
				}, nil)
				m.On("GetHouseStatusByName", "Active").Return(&models.HouseStatus{}, nil)
				m.On("CreateNewHouse", &models.House{
					UserID: 1,
				}).Return(nil, errors.New("error"))
			},
			args: args{
				userID:  1,
				request: &dtos.AddNewHouseRequestDTO{},
			},
			want:    nil,
			want2:   nil,
			wantErr: true,
		},
		{
			name:            "Failed - CreateNewHouse",
			houseRepository: mocks.NewHouseRepository(t),
			userRepository:  mocks.NewUserRepository(t),
			mock: func(m *mocks.HouseRepository, m2 *mocks.UserRepository) {
				m2.On("GetUserDetails", 1).Return(&dtos.UserDetailsDTO{
					Role: "host",
				}, nil)
				m.On("GetHouseStatusByName", "Active").Return(&models.HouseStatus{}, nil)
				m.On("CreateNewHouse", &models.House{
					UserID: 1,
				}).Return(&models.House{}, nil)
				m.On("AddHousePhoto", &models.HousePhoto{
					Photo: "image1",
				}).Return(nil, errors.New("error"))
			},
			args: args{
				userID: 1,
				request: &dtos.AddNewHouseRequestDTO{
					Photo: []string{
						"image1",
					},
				},
			},
			want:    nil,
			want2:   nil,
			wantErr: true,
		},
		{
			name:            "Success",
			houseRepository: mocks.NewHouseRepository(t),
			userRepository:  mocks.NewUserRepository(t),
			mock: func(m *mocks.HouseRepository, m2 *mocks.UserRepository) {
				m2.On("GetUserDetails", 1).Return(&dtos.UserDetailsDTO{
					Role: "host",
				}, nil)
				m.On("GetHouseStatusByName", "Active").Return(&models.HouseStatus{}, nil)
				m.On("CreateNewHouse", &models.House{
					UserID: 1,
				}).Return(&models.House{}, nil)
				m.On("AddHousePhoto", &models.HousePhoto{
					Photo: "image1",
				}).Return(&models.HousePhoto{
					Photo: "image1",
				}, nil)
			},
			args: args{
				userID: 1,
				request: &dtos.AddNewHouseRequestDTO{
					Photo: []string{
						"image1",
					},
				},
			},
			want: &models.House{},
			want2: []string{
				"image1",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &houseService{
				houseRepository: tt.houseRepository,
				userRepository:  tt.userRepository,
			}

			tt.mock(tt.houseRepository, tt.userRepository)

			got, got2, err := s.AddNewHouse(tt.args.userID, tt.args.request)

			if (err != nil) != tt.wantErr {
				t.Errorf("houseService.AddNewHouse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("houseService.AddNewHouse() = %v, want %v", got, tt.want)
			}

			if !reflect.DeepEqual(got2, tt.want2) {
				t.Errorf("houseService.AddNewHouse() = %v, want2 %v", got2, tt.want2)
			}
		})
	}
}

func Test_houseService_UpdateHouse(t *testing.T) {
	type args struct {
		houseID int
		userID  int
		request *dtos.UpdateHouseRequestDTO
	}
	tests := []struct {
		name            string
		houseRepository *mocks.HouseRepository
		userRepository  *mocks.UserRepository
		mock            func(m *mocks.HouseRepository, m2 *mocks.UserRepository)
		args            args
		want            *models.House
		wantErr         bool
	}{
		{
			name:            "Failed - GetUserDetails",
			houseRepository: mocks.NewHouseRepository(t),
			userRepository:  mocks.NewUserRepository(t),
			mock: func(m *mocks.HouseRepository, m2 *mocks.UserRepository) {
				m2.On("GetUserDetails", 1).Return(nil, errors.New("error"))
			},
			args: args{
				houseID: 1,
				userID:  1,
				request: &dtos.UpdateHouseRequestDTO{},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name:            "Failed - User is not Host or Admin",
			houseRepository: mocks.NewHouseRepository(t),
			userRepository:  mocks.NewUserRepository(t),
			mock: func(m *mocks.HouseRepository, m2 *mocks.UserRepository) {
				m2.On("GetUserDetails", 1).Return(&dtos.UserDetailsDTO{
					Role: "user",
				}, nil)
			},
			args: args{
				houseID: 1,
				userID:  1,
				request: &dtos.UpdateHouseRequestDTO{},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name:            "Failed - GetHouseByHouseID",
			houseRepository: mocks.NewHouseRepository(t),
			userRepository:  mocks.NewUserRepository(t),
			mock: func(m *mocks.HouseRepository, m2 *mocks.UserRepository) {
				m2.On("GetUserDetails", 1).Return(&dtos.UserDetailsDTO{
					Role: "host",
				}, nil)
				m.On("GetHouseByHouseID", 1).Return(nil, errors.New("error"))
			},
			args: args{
				houseID: 1,
				userID:  1,
				request: &dtos.UpdateHouseRequestDTO{},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name:            "Failed - User is not the Host or an Admin",
			houseRepository: mocks.NewHouseRepository(t),
			userRepository:  mocks.NewUserRepository(t),
			mock: func(m *mocks.HouseRepository, m2 *mocks.UserRepository) {
				m2.On("GetUserDetails", 1).Return(&dtos.UserDetailsDTO{
					Role: "host",
				}, nil)
				m.On("GetHouseByHouseID", 1).Return(&dtos.HouseDetailsDTO{
					HostID: 2,
				}, nil)
			},
			args: args{
				houseID: 1,
				userID:  1,
				request: &dtos.UpdateHouseRequestDTO{},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name:            "Failed - UpdateHouseByHouseID",
			houseRepository: mocks.NewHouseRepository(t),
			userRepository:  mocks.NewUserRepository(t),
			mock: func(m *mocks.HouseRepository, m2 *mocks.UserRepository) {
				m2.On("GetUserDetails", 1).Return(&dtos.UserDetailsDTO{
					Role: "host",
				}, nil)
				m.On("GetHouseByHouseID", 1).Return(&dtos.HouseDetailsDTO{
					HostID: 1,
				}, nil)
				m.On("UpdateHouseByHouseID", 1, &models.House{}).Return(nil, errors.New("error"))
			},
			args: args{
				houseID: 1,
				userID:  1,
				request: &dtos.UpdateHouseRequestDTO{},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name:            "Success - House Status: Deactivated by Admin",
			houseRepository: mocks.NewHouseRepository(t),
			userRepository:  mocks.NewUserRepository(t),
			mock: func(m *mocks.HouseRepository, m2 *mocks.UserRepository) {
				m2.On("GetUserDetails", 1).Return(&dtos.UserDetailsDTO{
					Role: "host",
				}, nil)
				m.On("GetHouseByHouseID", 1).Return(&dtos.HouseDetailsDTO{
					HostID:      1,
					HouseStatus: "Deactivated by Admin",
				}, nil)
				m.On("UpdateHouseByHouseID", 1, &models.House{}).Return(&models.House{}, nil)
			},
			args: args{
				houseID: 1,
				userID:  1,
				request: &dtos.UpdateHouseRequestDTO{},
			},
			want:    &models.House{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &houseService{
				houseRepository: tt.houseRepository,
				userRepository:  tt.userRepository,
			}

			tt.mock(tt.houseRepository, tt.userRepository)

			got, err := s.UpdateHouse(tt.args.houseID, tt.args.userID, tt.args.request)

			if (err != nil) != tt.wantErr {
				t.Errorf("houseService.UpdateHouse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("houseService.UpdateHouse() = %v, want %v", got, tt.want)
			}

		})
	}
}

func Test_houseService_DeleteHouse(t *testing.T) {
	type args struct {
		houseID int
		userID  int
	}
	tests := []struct {
		name                  string
		houseRepository       *mocks.HouseRepository
		userRepository        *mocks.UserRepository
		reservationRepository *mocks.ReservationRepository
		mock                  func(m *mocks.HouseRepository, m2 *mocks.UserRepository, m3 *mocks.ReservationRepository)
		args                  args
		want                  *models.House
		wantErr               bool
	}{
		{
			name:                  "Failed - GetHouseByHouseID",
			houseRepository:       mocks.NewHouseRepository(t),
			userRepository:        mocks.NewUserRepository(t),
			reservationRepository: mocks.NewReservationRepository(t),
			mock: func(m *mocks.HouseRepository, m2 *mocks.UserRepository, m3 *mocks.ReservationRepository) {
				m.On("GetHouseByHouseID", 1).Return(nil, errors.New("error"))
			},
			args: args{
				houseID: 1,
				userID:  1,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name:                  "Failed - GetUserByID",
			houseRepository:       mocks.NewHouseRepository(t),
			userRepository:        mocks.NewUserRepository(t),
			reservationRepository: mocks.NewReservationRepository(t),
			mock: func(m *mocks.HouseRepository, m2 *mocks.UserRepository, m3 *mocks.ReservationRepository) {
				m.On("GetHouseByHouseID", 1).Return(&dtos.HouseDetailsDTO{}, nil)
				m2.On("GetUserByID", 1).Return(nil, errors.New("error"))
			},
			args: args{
				houseID: 1,
				userID:  1,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name:                  "Failed - User is not Host or Admin",
			houseRepository:       mocks.NewHouseRepository(t),
			userRepository:        mocks.NewUserRepository(t),
			reservationRepository: mocks.NewReservationRepository(t),
			mock: func(m *mocks.HouseRepository, m2 *mocks.UserRepository, m3 *mocks.ReservationRepository) {
				m.On("GetHouseByHouseID", 1).Return(&dtos.HouseDetailsDTO{}, nil)
				m2.On("GetUserByID", 1).Return(&models.User{
					Role: "user",
				}, nil)
			},
			args: args{
				houseID: 1,
				userID:  1,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name:                  "Failed - User is not the Host or Admin",
			houseRepository:       mocks.NewHouseRepository(t),
			userRepository:        mocks.NewUserRepository(t),
			reservationRepository: mocks.NewReservationRepository(t),
			mock: func(m *mocks.HouseRepository, m2 *mocks.UserRepository, m3 *mocks.ReservationRepository) {
				m.On("GetHouseByHouseID", 1).Return(&dtos.HouseDetailsDTO{
					HostID: 2,
				}, nil)
				m2.On("GetUserByID", 1).Return(&models.User{
					Role: "host",
				}, nil)
			},
			args: args{
				houseID: 1,
				userID:  1,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name:                  "Failed - GetHouseStatusByName",
			houseRepository:       mocks.NewHouseRepository(t),
			userRepository:        mocks.NewUserRepository(t),
			reservationRepository: mocks.NewReservationRepository(t),
			mock: func(m *mocks.HouseRepository, m2 *mocks.UserRepository, m3 *mocks.ReservationRepository) {
				m.On("GetHouseByHouseID", 1).Return(&dtos.HouseDetailsDTO{
					HostID: 1,
				}, nil)
				m2.On("GetUserByID", 1).Return(&models.User{
					Role: "host",
				}, nil)
				m.On("GetHouseStatusByName", "inactive").Return(nil, errors.New("error"))
			},
			args: args{
				houseID: 1,
				userID:  1,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name:                  "Failed - UpdateHouseByHouseID",
			houseRepository:       mocks.NewHouseRepository(t),
			userRepository:        mocks.NewUserRepository(t),
			reservationRepository: mocks.NewReservationRepository(t),
			mock: func(m *mocks.HouseRepository, m2 *mocks.UserRepository, m3 *mocks.ReservationRepository) {
				m.On("GetHouseByHouseID", 1).Return(&dtos.HouseDetailsDTO{
					HostID: 1,
				}, nil)
				m2.On("GetUserByID", 1).Return(&models.User{
					Role: "host",
				}, nil)
				m.On("GetHouseStatusByName", "inactive").Return(&models.HouseStatus{
					ID: 2,
				}, nil)
				m.On("UpdateHouseByHouseID", 1, &models.House{
					HouseStatusID: 2,
				}).Return(nil, errors.New("error"))
			},
			args: args{
				houseID: 1,
				userID:  1,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name:                  "Failed - GetHouseReservationsByHouseID",
			houseRepository:       mocks.NewHouseRepository(t),
			userRepository:        mocks.NewUserRepository(t),
			reservationRepository: mocks.NewReservationRepository(t),
			mock: func(m *mocks.HouseRepository, m2 *mocks.UserRepository, m3 *mocks.ReservationRepository) {
				m.On("GetHouseByHouseID", 1).Return(&dtos.HouseDetailsDTO{
					HostID: 1,
				}, nil)
				m2.On("GetUserByID", 1).Return(&models.User{
					Role: "host",
				}, nil)
				m.On("GetHouseStatusByName", "inactive").Return(&models.HouseStatus{
					ID: 2,
				}, nil)
				m.On("UpdateHouseByHouseID", 1, &models.House{
					HouseStatusID: 2,
				}).Return(nil, nil)
				m3.On("GetHouseReservationsByHouseID", 1, map[string]interface{}{
					"filterBy": "upcoming",
				}).Return(nil, nil, errors.New("error"))
			},
			args: args{
				houseID: 1,
				userID:  1,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name:                  "Failed - Delete House but have Upcoming Reservation",
			houseRepository:       mocks.NewHouseRepository(t),
			userRepository:        mocks.NewUserRepository(t),
			reservationRepository: mocks.NewReservationRepository(t),
			mock: func(m *mocks.HouseRepository, m2 *mocks.UserRepository, m3 *mocks.ReservationRepository) {
				m.On("GetHouseByHouseID", 1).Return(&dtos.HouseDetailsDTO{
					HostID: 1,
				}, nil)
				m2.On("GetUserByID", 1).Return(&models.User{
					Role: "host",
				}, nil)
				m.On("GetHouseStatusByName", "inactive").Return(&models.HouseStatus{
					ID: 2,
				}, nil)
				m.On("UpdateHouseByHouseID", 1, &models.House{
					HouseStatusID: 2,
				}).Return(nil, nil)
				m3.On("GetHouseReservationsByHouseID", 1, map[string]interface{}{
					"filterBy": "upcoming",
				}).Return([]*dtos.HouseReservationsDTO{
					{
						ReservationID: 1,
					},
				}, nil, nil)
			},
			args: args{
				houseID: 1,
				userID:  1,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name:                  "Failed - DeleteHousePhotoByHouseID",
			houseRepository:       mocks.NewHouseRepository(t),
			userRepository:        mocks.NewUserRepository(t),
			reservationRepository: mocks.NewReservationRepository(t),
			mock: func(m *mocks.HouseRepository, m2 *mocks.UserRepository, m3 *mocks.ReservationRepository) {
				m.On("GetHouseByHouseID", 1).Return(&dtos.HouseDetailsDTO{
					HostID: 1,
				}, nil)
				m2.On("GetUserByID", 1).Return(&models.User{
					Role: "host",
				}, nil)
				m.On("GetHouseStatusByName", "inactive").Return(&models.HouseStatus{
					ID: 2,
				}, nil)
				m.On("UpdateHouseByHouseID", 1, &models.House{
					HouseStatusID: 2,
				}).Return(nil, nil)
				m3.On("GetHouseReservationsByHouseID", 1, map[string]interface{}{
					"filterBy": "upcoming",
				}).Return([]*dtos.HouseReservationsDTO{}, nil, nil)
				m.On("DeleteHousePhotoByHouseID", 1).Return(errors.New("error"))
			},
			args: args{
				houseID: 1,
				userID:  1,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name:                  "Failed - DeleteHouseByID",
			houseRepository:       mocks.NewHouseRepository(t),
			userRepository:        mocks.NewUserRepository(t),
			reservationRepository: mocks.NewReservationRepository(t),
			mock: func(m *mocks.HouseRepository, m2 *mocks.UserRepository, m3 *mocks.ReservationRepository) {
				m.On("GetHouseByHouseID", 1).Return(&dtos.HouseDetailsDTO{
					HostID: 1,
				}, nil)
				m2.On("GetUserByID", 1).Return(&models.User{
					Role: "host",
				}, nil)
				m.On("GetHouseStatusByName", "inactive").Return(&models.HouseStatus{
					ID: 2,
				}, nil)
				m.On("UpdateHouseByHouseID", 1, &models.House{
					HouseStatusID: 2,
				}).Return(nil, nil)
				m3.On("GetHouseReservationsByHouseID", 1, map[string]interface{}{
					"filterBy": "upcoming",
				}).Return([]*dtos.HouseReservationsDTO{}, nil, nil)
				m.On("DeleteHousePhotoByHouseID", 1).Return(nil)
				m.On("DeleteHouseByID", 1).Return(nil, errors.New("error"))
			},
			args: args{
				houseID: 1,
				userID:  1,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name:                  "Success - Delete House as Admin",
			houseRepository:       mocks.NewHouseRepository(t),
			userRepository:        mocks.NewUserRepository(t),
			reservationRepository: mocks.NewReservationRepository(t),
			mock: func(m *mocks.HouseRepository, m2 *mocks.UserRepository, m3 *mocks.ReservationRepository) {
				m.On("GetHouseByHouseID", 1).Return(&dtos.HouseDetailsDTO{
					HostID: 1,
				}, nil)
				m2.On("GetUserByID", 1).Return(&models.User{
					Role: "admin",
				}, nil)
				m.On("GetHouseStatusByName", "deactivated by admin").Return(&models.HouseStatus{
					ID: 3,
				}, nil)
				m.On("UpdateHouseByHouseID", 1, &models.House{
					HouseStatusID: 3,
				}).Return(nil, nil)
				m3.On("GetHouseReservationsByHouseID", 1, map[string]interface{}{
					"filterBy": "upcoming",
				}).Return([]*dtos.HouseReservationsDTO{}, nil, nil)
				m.On("DeleteHousePhotoByHouseID", 1).Return(nil)
				m.On("DeleteHouseByID", 1).Return(&models.House{}, nil)
			},
			args: args{
				houseID: 1,
				userID:  1,
			},
			want:    &models.House{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &houseService{
				houseRepository:       tt.houseRepository,
				userRepository:        tt.userRepository,
				reservationRepository: tt.reservationRepository,
			}

			tt.mock(tt.houseRepository, tt.userRepository, tt.reservationRepository)

			got, err := s.DeleteHouse(tt.args.houseID, tt.args.userID)

			if (err != nil) != tt.wantErr {
				t.Errorf("houseService.DeleteHouse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("houseService.DeleteHouse() = %v, want %v", got, tt.want)
			}

		})
	}
}

func Test_houseService_AddHousePhoto(t *testing.T) {
	type args struct {
		houseID int
		userID  int
		request *dtos.AddHousePhotoRequestDTO
	}
	tests := []struct {
		name            string
		houseRepository *mocks.HouseRepository
		userRepository  *mocks.UserRepository
		mock            func(m *mocks.HouseRepository, m2 *mocks.UserRepository)
		args            args
		want            []string
		wantErr         bool
	}{
		{
			name:            "Failed - GetUserDetails",
			houseRepository: mocks.NewHouseRepository(t),
			userRepository:  mocks.NewUserRepository(t),
			mock: func(m *mocks.HouseRepository, m2 *mocks.UserRepository) {
				m2.On("GetUserDetails", 1).Return(nil, errors.New("error"))
			},
			args: args{
				houseID: 1,
				userID:  1,
				request: &dtos.AddHousePhotoRequestDTO{},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name:            "Failed - User is not a Host",
			houseRepository: mocks.NewHouseRepository(t),
			userRepository:  mocks.NewUserRepository(t),
			mock: func(m *mocks.HouseRepository, m2 *mocks.UserRepository) {
				m2.On("GetUserDetails", 1).Return(&dtos.UserDetailsDTO{
					Role: "user",
				}, nil)
			},
			args: args{
				houseID: 1,
				userID:  1,
				request: &dtos.AddHousePhotoRequestDTO{},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name:            "Failed - AddHousePhoto",
			houseRepository: mocks.NewHouseRepository(t),
			userRepository:  mocks.NewUserRepository(t),
			mock: func(m *mocks.HouseRepository, m2 *mocks.UserRepository) {
				m2.On("GetUserDetails", 1).Return(&dtos.UserDetailsDTO{
					Role: "host",
				}, nil)
				m.On("AddHousePhoto", &models.HousePhoto{
					HouseID: 1,
					Photo:   "image1",
				}).Return(nil, errors.New("error"))
			},
			args: args{
				houseID: 1,
				userID:  1,
				request: &dtos.AddHousePhotoRequestDTO{
					Photo: []string{
						"image1",
					},
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name:            "Failed - AddHousePhoto",
			houseRepository: mocks.NewHouseRepository(t),
			userRepository:  mocks.NewUserRepository(t),
			mock: func(m *mocks.HouseRepository, m2 *mocks.UserRepository) {
				m2.On("GetUserDetails", 1).Return(&dtos.UserDetailsDTO{
					Role: "host",
				}, nil)
				m.On("AddHousePhoto", &models.HousePhoto{
					HouseID: 1,
					Photo:   "image1",
				}).Return(&models.HousePhoto{
					Photo: "image1",
				}, nil)
			},
			args: args{
				houseID: 1,
				userID:  1,
				request: &dtos.AddHousePhotoRequestDTO{
					Photo: []string{
						"image1",
					},
				},
			},
			want: []string{
				"image1",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &houseService{
				houseRepository: tt.houseRepository,
				userRepository:  tt.userRepository,
			}

			tt.mock(tt.houseRepository, tt.userRepository)

			got, err := s.AddHousePhoto(tt.args.houseID, tt.args.userID, tt.args.request)

			if (err != nil) != tt.wantErr {
				t.Errorf("houseService.AddHousePhoto() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("houseService.AddHousePhoto() = %v, want %v", got, tt.want)
			}

		})
	}
}

func Test_houseService_DeleteHousePhoto(t *testing.T) {
	type args struct {
		housePhotoID int
		userID       int
	}
	tests := []struct {
		name            string
		houseRepository *mocks.HouseRepository
		mock            func(m *mocks.HouseRepository)
		args            args
		want            *models.HousePhoto
		wantErr         bool
	}{
		{
			name:            "Failed - GetHouseByHousePhotoID",
			houseRepository: mocks.NewHouseRepository(t),
			mock: func(m *mocks.HouseRepository) {
				m.On("GetHouseByHousePhotoID", 1).Return(nil, errors.New("error"))
			},
			args: args{
				housePhotoID: 1,
				userID:       1,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name:            "Failed - User is not the Host",
			houseRepository: mocks.NewHouseRepository(t),
			mock: func(m *mocks.HouseRepository) {
				m.On("GetHouseByHousePhotoID", 1).Return(&dtos.HouseDetailsDTO{
					HostID: 2,
				}, nil)
			},
			args: args{
				housePhotoID: 1,
				userID:       1,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name:            "Failed - ValidateTotalHousePhoto",
			houseRepository: mocks.NewHouseRepository(t),
			mock: func(m *mocks.HouseRepository) {
				m.On("GetHouseByHousePhotoID", 1).Return(&dtos.HouseDetailsDTO{
					HouseID: 1,
					HostID:  1,
				}, nil)
				m.On("ValidateTotalHousePhoto", 1).Return(int64(0), errors.New("error"))
			},
			args: args{
				housePhotoID: 1,
				userID:       1,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name:            "Failed - House only have one or less Photo",
			houseRepository: mocks.NewHouseRepository(t),
			mock: func(m *mocks.HouseRepository) {
				m.On("GetHouseByHousePhotoID", 1).Return(&dtos.HouseDetailsDTO{
					HouseID: 1,
					HostID:  1,
				}, nil)
				m.On("ValidateTotalHousePhoto", 1).Return(int64(1), nil)
			},
			args: args{
				housePhotoID: 1,
				userID:       1,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name:            "Failed - DeleteHousePhotoByID",
			houseRepository: mocks.NewHouseRepository(t),
			mock: func(m *mocks.HouseRepository) {
				m.On("GetHouseByHousePhotoID", 1).Return(&dtos.HouseDetailsDTO{
					HouseID: 1,
					HostID:  1,
				}, nil)
				m.On("ValidateTotalHousePhoto", 1).Return(int64(2), nil)
				m.On("DeleteHousePhotoByID", 1).Return(nil, errors.New("error"))
			},
			args: args{
				housePhotoID: 1,
				userID:       1,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name:            "Success",
			houseRepository: mocks.NewHouseRepository(t),
			mock: func(m *mocks.HouseRepository) {
				m.On("GetHouseByHousePhotoID", 1).Return(&dtos.HouseDetailsDTO{
					HouseID: 1,
					HostID:  1,
				}, nil)
				m.On("ValidateTotalHousePhoto", 1).Return(int64(2), nil)
				m.On("DeleteHousePhotoByID", 1).Return(&models.HousePhoto{}, nil)
			},
			args: args{
				housePhotoID: 1,
				userID:       1,
			},
			want:    &models.HousePhoto{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &houseService{
				houseRepository: tt.houseRepository,
			}

			tt.mock(tt.houseRepository)

			got, err := s.DeleteHousePhoto(tt.args.housePhotoID, tt.args.userID)

			if (err != nil) != tt.wantErr {
				t.Errorf("houseService.DeleteHousePhoto() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("houseService.DeleteHousePhoto() = %v, want %v", got, tt.want)
			}

		})
	}
}
