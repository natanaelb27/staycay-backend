package services

import (
	"errors"
	"reflect"
	"testing"

	"git.garena.com/sea-labs-id/batch-02/natanael-benediktus/final-project-backend/internal/dtos"
	"git.garena.com/sea-labs-id/batch-02/natanael-benediktus/final-project-backend/internal/mocks"
	"git.garena.com/sea-labs-id/batch-02/natanael-benediktus/final-project-backend/internal/models"
	"github.com/stretchr/testify/assert"
)

func Test_userService_NewUserService(t *testing.T) {
	userRepository := mocks.NewUserRepository(t)
	hashHelper := mocks.NewHashHelper(t)
	userService := NewUserService(
		&USConfig{
			UserRepository: userRepository,
			HashHelper:     hashHelper,
		},
	)

	assert.NotNil(t, userService)
}

func Test_userService_GetUserProfile(t *testing.T) {
	type args struct {
		userID int
	}
	tests := []struct {
		name           string
		userRepository *mocks.UserRepository
		mock           func(m *mocks.UserRepository)
		args           args
		want           *dtos.UserDetailsDTO
		wantErr        bool
	}{
		{
			name:           "Failed - GetUserDetails",
			userRepository: mocks.NewUserRepository(t),
			mock: func(m *mocks.UserRepository) {
				m.On("GetUserDetails", 1).Return(nil, errors.New("error"))
			},
			args: args{
				userID: 1,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name:           "Success",
			userRepository: mocks.NewUserRepository(t),
			mock: func(m *mocks.UserRepository) {
				m.On("GetUserDetails", 1).Return(&dtos.UserDetailsDTO{}, nil)
			},
			args: args{
				userID: 1,
			},
			want:    &dtos.UserDetailsDTO{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &userService{
				userRepository: tt.userRepository,
			}

			tt.mock(tt.userRepository)

			got, err := s.GetUserProfile(tt.args.userID)

			if (err != nil) != tt.wantErr {
				t.Errorf("userService.GetUserProfile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("userService.GetUserProfile() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_userService_GetUserProfileByUserID(t *testing.T) {
	type args struct {
		userID int
	}
	tests := []struct {
		name           string
		userRepository *mocks.UserRepository
		mock           func(m *mocks.UserRepository)
		args           args
		want           *dtos.UserDetailsDTO
		wantErr        bool
	}{
		{
			name:           "Failed - GetUserDetails",
			userRepository: mocks.NewUserRepository(t),
			mock: func(m *mocks.UserRepository) {
				m.On("GetUserDetails", 1).Return(nil, errors.New("error"))
			},
			args: args{
				userID: 1,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name:           "Success",
			userRepository: mocks.NewUserRepository(t),
			mock: func(m *mocks.UserRepository) {
				m.On("GetUserDetails", 1).Return(&dtos.UserDetailsDTO{}, nil)
			},
			args: args{
				userID: 1,
			},
			want:    &dtos.UserDetailsDTO{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &userService{
				userRepository: tt.userRepository,
			}

			tt.mock(tt.userRepository)

			got, err := s.GetUserProfileByUserID(tt.args.userID)

			if (err != nil) != tt.wantErr {
				t.Errorf("userService.GetUserProfileByUserID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("userService.GetUserProfileByUserID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_userService_UpdateUserRoleToHost(t *testing.T) {
	type args struct {
		userID int
	}
	tests := []struct {
		name           string
		userRepository *mocks.UserRepository
		mock           func(m *mocks.UserRepository)
		args           args
		want           *models.User
		wantErr        bool
	}{
		{
			name:           "Failed - UpdateUserRole",
			userRepository: mocks.NewUserRepository(t),
			mock: func(m *mocks.UserRepository) {
				m.On("UpdateUserRole", 1, "host").Return(nil, errors.New("error"))
			},
			args: args{
				userID: 1,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name:           "Success",
			userRepository: mocks.NewUserRepository(t),
			mock: func(m *mocks.UserRepository) {
				m.On("UpdateUserRole", 1, "host").Return(&models.User{}, nil)
			},
			args: args{
				userID: 1,
			},
			want:    &models.User{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &userService{
				userRepository: tt.userRepository,
			}

			tt.mock(tt.userRepository)

			got, err := s.UpdateUserRoleToHost(tt.args.userID)

			if (err != nil) != tt.wantErr {
				t.Errorf("userService.UpdateUserRoleToHost() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("userService.UpdateUserRoleToHost() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_userService_UpdateUserProfile(t *testing.T) {
	type args struct {
		userID  int
		request *dtos.UpdateUserProfileRequestDTO
	}
	tests := []struct {
		name           string
		userRepository *mocks.UserRepository
		hashHelper     *mocks.HashHelper
		mock           func(m *mocks.UserRepository, m2 *mocks.HashHelper)
		args           args
		want           *models.User
		wantErr        bool
	}{
		{
			name:           "Failed - GetUserByID",
			userRepository: mocks.NewUserRepository(t),
			hashHelper:     mocks.NewHashHelper(t),
			mock: func(m *mocks.UserRepository, m2 *mocks.HashHelper) {
				m.On("GetUserByID", 1).Return(nil, errors.New("error"))
			},
			args: args{
				userID:  1,
				request: &dtos.UpdateUserProfileRequestDTO{},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name:           "Failed - Invalid New Password",
			userRepository: mocks.NewUserRepository(t),
			hashHelper:     mocks.NewHashHelper(t),
			mock: func(m *mocks.UserRepository, m2 *mocks.HashHelper) {
				m.On("GetUserByID", 1).Return(&models.User{}, nil)
			},
			args: args{
				userID: 1,
				request: &dtos.UpdateUserProfileRequestDTO{
					Password: "invalid",
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name:           "Failed - Old Password Does Not Match",
			userRepository: mocks.NewUserRepository(t),
			hashHelper:     mocks.NewHashHelper(t),
			mock: func(m *mocks.UserRepository, m2 *mocks.HashHelper) {
				m.On("GetUserByID", 1).Return(&models.User{
					Password: "PasswordEncrypted",
				}, nil)
				m2.On("CompareEncrypt", "OldPassword123", "PasswordEncrypted").Return(false)
			},
			args: args{
				userID: 1,
				request: &dtos.UpdateUserProfileRequestDTO{
					OldPassword: "OldPassword123",
					Password:    "Password123",
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name:           "Failed - Encrypt Password",
			userRepository: mocks.NewUserRepository(t),
			hashHelper:     mocks.NewHashHelper(t),
			mock: func(m *mocks.UserRepository, m2 *mocks.HashHelper) {
				m.On("GetUserByID", 1).Return(&models.User{
					Password: "PasswordEncrypted",
				}, nil)
				m2.On("CompareEncrypt", "OldPassword123", "PasswordEncrypted").Return(true)
				m2.On("Encrypt", "Password123").Return("", errors.New("error"))
			},
			args: args{
				userID: 1,
				request: &dtos.UpdateUserProfileRequestDTO{
					OldPassword: "OldPassword123",
					Password:    "Password123",
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name:           "Failed - UpdateUser",
			userRepository: mocks.NewUserRepository(t),
			hashHelper:     mocks.NewHashHelper(t),
			mock: func(m *mocks.UserRepository, m2 *mocks.HashHelper) {
				m.On("GetUserByID", 1).Return(&models.User{
					Password: "PasswordEncrypted",
				}, nil)
				m2.On("CompareEncrypt", "OldPassword123", "PasswordEncrypted").Return(true)
				m2.On("Encrypt", "Password123").Return("NewPasswordEncrypted", nil)
				m.On("UpdateUser", 1, &models.User{
					Name:     "Name Updated",
					Password: "NewPasswordEncrypted",
					Address:  "Address Updated",
					CityID:   2,
				}).Return(nil, errors.New("error"))
			},
			args: args{
				userID: 1,
				request: &dtos.UpdateUserProfileRequestDTO{
					Name:        "Name Updated",
					OldPassword: "OldPassword123",
					Password:    "Password123",
					Address:     "Address Updated",
					CityID:      2,
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name:           "Success",
			userRepository: mocks.NewUserRepository(t),
			hashHelper:     mocks.NewHashHelper(t),
			mock: func(m *mocks.UserRepository, m2 *mocks.HashHelper) {
				m.On("GetUserByID", 1).Return(&models.User{
					Password: "PasswordEncrypted",
				}, nil)
				m2.On("CompareEncrypt", "OldPassword123", "PasswordEncrypted").Return(true)
				m2.On("Encrypt", "Password123").Return("NewPasswordEncrypted", nil)
				m.On("UpdateUser", 1, &models.User{
					Name:     "Name Updated",
					Password: "NewPasswordEncrypted",
					Address:  "Address Updated",
					CityID:   2,
				}).Return(&models.User{}, nil)
			},
			args: args{
				userID: 1,
				request: &dtos.UpdateUserProfileRequestDTO{
					Name:        "Name Updated",
					OldPassword: "OldPassword123",
					Password:    "Password123",
					Address:     "Address Updated",
					CityID:      2,
				},
			},
			want:    &models.User{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &userService{
				userRepository: tt.userRepository,
				hashHelper:     tt.hashHelper,
			}

			tt.mock(tt.userRepository, tt.hashHelper)

			got, err := s.UpdateUserProfile(tt.args.userID, tt.args.request)

			if (err != nil) != tt.wantErr {
				t.Errorf("userService.UpdateUserProfile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("userService.UpdateUserProfile() = %v, want %v", got, tt.want)
			}
		})
	}
}
