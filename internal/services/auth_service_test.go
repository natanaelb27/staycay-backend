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

func Test_cityService_NewAuthService(t *testing.T) {
	userRepository := mocks.NewUserRepository(t)
	walletRepository := mocks.NewWalletRepository(t)
	gameRepository := mocks.NewGameRepository(t)
	hashHelper := mocks.NewHashHelper(t)
	authService := NewAuthService(
		&ASConfig{
			UserRepository:   userRepository,
			WalletRepository: walletRepository,
			GameRepository:   gameRepository,
			HashHelper:       hashHelper,
		},
	)

	assert.NotNil(t, authService)
}

func Test_authService_Register(t *testing.T) {
	type args struct {
		userData dtos.RegisterRequestDTO
	}
	tests := []struct {
		name             string
		userRepository   *mocks.UserRepository
		walletRepository *mocks.WalletRepository
		gameRepository   *mocks.GameRepository
		hashHelper       *mocks.HashHelper
		mock             func(m *mocks.UserRepository, m2 *mocks.WalletRepository, m3 *mocks.GameRepository, m4 *mocks.HashHelper)
		args             args
		want             *models.User
		want2            *string
		wantErr          bool
	}{
		{
			name:             "Failed - Invalid Email Address",
			userRepository:   mocks.NewUserRepository(t),
			walletRepository: mocks.NewWalletRepository(t),
			gameRepository:   mocks.NewGameRepository(t),
			hashHelper:       mocks.NewHashHelper(t),
			mock: func(m *mocks.UserRepository, m2 *mocks.WalletRepository, m3 *mocks.GameRepository, m4 *mocks.HashHelper) {
			},
			args: args{
				userData: dtos.RegisterRequestDTO{
					Email: "invalid_email",
				},
			},
			want:    nil,
			want2:   nil,
			wantErr: true,
		},
		{
			name:             "Failed - Invalid Password",
			userRepository:   mocks.NewUserRepository(t),
			walletRepository: mocks.NewWalletRepository(t),
			gameRepository:   mocks.NewGameRepository(t),
			hashHelper:       mocks.NewHashHelper(t),
			mock: func(m *mocks.UserRepository, m2 *mocks.WalletRepository, m3 *mocks.GameRepository, m4 *mocks.HashHelper) {
			},
			args: args{
				userData: dtos.RegisterRequestDTO{
					Email:    "valid_email@email.com",
					Password: "invalid",
				},
			},
			want:    nil,
			want2:   nil,
			wantErr: true,
		},
		{
			name:             "Failed - Encrypt Password",
			userRepository:   mocks.NewUserRepository(t),
			walletRepository: mocks.NewWalletRepository(t),
			gameRepository:   mocks.NewGameRepository(t),
			hashHelper:       mocks.NewHashHelper(t),
			mock: func(m *mocks.UserRepository, m2 *mocks.WalletRepository, m3 *mocks.GameRepository, m4 *mocks.HashHelper) {
				m4.On("Encrypt", "Password123").Return("", errors.New("error"))
			},
			args: args{
				userData: dtos.RegisterRequestDTO{
					Email:    "valid_email@email.com",
					Password: "Password123",
				},
			},
			want:    nil,
			want2:   nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &authService{
				userRepository:   tt.userRepository,
				walletRepository: tt.walletRepository,
				gameRepository:   tt.gameRepository,
				hashHelper:       tt.hashHelper,
			}

			tt.mock(tt.userRepository, tt.walletRepository, tt.gameRepository, tt.hashHelper)

			got, got2, err := s.Register(tt.args.userData)

			if (err != nil) != tt.wantErr {
				t.Errorf("authService.Register() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("authService.Register() = %v, want %v", got, tt.want)
			}

			if !reflect.DeepEqual(got2, tt.want2) {
				t.Errorf("authService.Register() = %v, want2 %v", got2, tt.want2)
			}
		})
	}
}
