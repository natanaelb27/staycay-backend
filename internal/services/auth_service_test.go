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

func Test_authService_NewAuthService(t *testing.T) {
	userRepository := mocks.NewUserRepository(t)
	walletRepository := mocks.NewWalletRepository(t)
	gameRepository := mocks.NewGameRepository(t)
	hashHelper := mocks.NewHashHelper(t)
	jwtHelper := mocks.NewJwtHelper(t)
	authService := NewAuthService(
		&ASConfig{
			UserRepository:   userRepository,
			WalletRepository: walletRepository,
			GameRepository:   gameRepository,
			HashHelper:       hashHelper,
			JwtHelper:        jwtHelper,
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
		jwtHelper        *mocks.JwtHelper
		mock             func(m *mocks.UserRepository, m2 *mocks.WalletRepository, m3 *mocks.GameRepository, m4 *mocks.HashHelper, m5 *mocks.JwtHelper)
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
			jwtHelper:        mocks.NewJwtHelper(t),
			mock: func(m *mocks.UserRepository, m2 *mocks.WalletRepository, m3 *mocks.GameRepository, m4 *mocks.HashHelper, m5 *mocks.JwtHelper) {
			},
			args: args{
				userData: dtos.RegisterRequestDTO{
					Name:     "Name",
					Email:    "invalid_email",
					Password: "Password123",
					Address:  "Badak Street",
					CityID:   1,
					IsAdmin:  false,
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
			jwtHelper:        mocks.NewJwtHelper(t),
			mock: func(m *mocks.UserRepository, m2 *mocks.WalletRepository, m3 *mocks.GameRepository, m4 *mocks.HashHelper, m5 *mocks.JwtHelper) {
			},
			args: args{
				userData: dtos.RegisterRequestDTO{
					Name:     "Name",
					Email:    "valid@email.com",
					Password: "invalid",
					Address:  "Badak Street",
					CityID:   1,
					IsAdmin:  false,
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
			jwtHelper:        mocks.NewJwtHelper(t),
			mock: func(m *mocks.UserRepository, m2 *mocks.WalletRepository, m3 *mocks.GameRepository, m4 *mocks.HashHelper, m5 *mocks.JwtHelper) {
				m4.On("Encrypt", "Password123").Return("", errors.New("error"))
			},
			args: args{
				userData: dtos.RegisterRequestDTO{
					Name:     "Name",
					Email:    "valid@email.com",
					Password: "Password123",
					Address:  "Badak Street",
					CityID:   1,
					IsAdmin:  false,
				},
			},
			want:    nil,
			want2:   nil,
			wantErr: true,
		},
		{
			name:             "Failed - CreateNewUser",
			userRepository:   mocks.NewUserRepository(t),
			walletRepository: mocks.NewWalletRepository(t),
			gameRepository:   mocks.NewGameRepository(t),
			hashHelper:       mocks.NewHashHelper(t),
			jwtHelper:        mocks.NewJwtHelper(t),
			mock: func(m *mocks.UserRepository, m2 *mocks.WalletRepository, m3 *mocks.GameRepository, m4 *mocks.HashHelper, m5 *mocks.JwtHelper) {
				m4.On("Encrypt", "Password123").Return("PasswordEncrypted", nil)
				m.On("CreateNewUser", &models.User{
					Name:     "Name",
					Email:    "valid@email.com",
					Password: "PasswordEncrypted",
					Address:  "Badak Street",
					CityID:   1,
					Role:     "user",
				}).Return(nil, errors.New("error"))
			},
			args: args{
				userData: dtos.RegisterRequestDTO{
					Name:     "Name",
					Email:    "valid@email.com",
					Password: "Password123",
					Address:  "Badak Street",
					CityID:   1,
					IsAdmin:  false,
				},
			},
			want:    nil,
			want2:   nil,
			wantErr: true,
		},
		{
			name:             "Failed - CreateWallet",
			userRepository:   mocks.NewUserRepository(t),
			walletRepository: mocks.NewWalletRepository(t),
			gameRepository:   mocks.NewGameRepository(t),
			hashHelper:       mocks.NewHashHelper(t),
			jwtHelper:        mocks.NewJwtHelper(t),
			mock: func(m *mocks.UserRepository, m2 *mocks.WalletRepository, m3 *mocks.GameRepository, m4 *mocks.HashHelper, m5 *mocks.JwtHelper) {
				m4.On("Encrypt", "Password123").Return("PasswordEncrypted", nil)
				m.On("CreateNewUser", &models.User{
					Name:     "Name",
					Email:    "valid@email.com",
					Password: "PasswordEncrypted",
					Address:  "Badak Street",
					CityID:   1,
					Role:     "user",
				}).Return(&models.User{
					ID:       1,
					Name:     "Name",
					Email:    "valid@email.com",
					Password: "Password123",
					Address:  "Badak Street",
					CityID:   1,
					Role:     "user",
				}, nil)
				m2.On("CreateWallet", &models.Wallet{
					UserID:  1,
					Balance: 0,
				}).Return(errors.New("error"))
			},
			args: args{
				userData: dtos.RegisterRequestDTO{
					Name:     "Name",
					Email:    "valid@email.com",
					Password: "Password123",
					Address:  "Badak Street",
					CityID:   1,
					IsAdmin:  false,
				},
			},
			want:    nil,
			want2:   nil,
			wantErr: true,
		},
		{
			name:             "Failed - CreateGameUser",
			userRepository:   mocks.NewUserRepository(t),
			walletRepository: mocks.NewWalletRepository(t),
			gameRepository:   mocks.NewGameRepository(t),
			hashHelper:       mocks.NewHashHelper(t),
			jwtHelper:        mocks.NewJwtHelper(t),
			mock: func(m *mocks.UserRepository, m2 *mocks.WalletRepository, m3 *mocks.GameRepository, m4 *mocks.HashHelper, m5 *mocks.JwtHelper) {
				m4.On("Encrypt", "Password123").Return("PasswordEncrypted", nil)
				m.On("CreateNewUser", &models.User{
					Name:     "Name",
					Email:    "valid@email.com",
					Password: "PasswordEncrypted",
					Address:  "Badak Street",
					CityID:   1,
					Role:     "user",
				}).Return(&models.User{
					ID:       1,
					Name:     "Name",
					Email:    "valid@email.com",
					Password: "Password123",
					Address:  "Badak Street",
					CityID:   1,
					Role:     "user",
				}, nil)
				m2.On("CreateWallet", &models.Wallet{
					UserID:  1,
					Balance: 0,
				}).Return(nil)
				m3.On("CreateGameUser", &models.GameChance{
					UserID:  1,
					Chance:  0,
					Counter: 0,
				}).Return(errors.New("error"))
			},
			args: args{
				userData: dtos.RegisterRequestDTO{
					Name:     "Name",
					Email:    "valid@email.com",
					Password: "Password123",
					Address:  "Badak Street",
					CityID:   1,
					IsAdmin:  false,
				},
			},
			want:    nil,
			want2:   nil,
			wantErr: true,
		},
		{
			name:             "Failed - GenerateJwtToken",
			userRepository:   mocks.NewUserRepository(t),
			walletRepository: mocks.NewWalletRepository(t),
			gameRepository:   mocks.NewGameRepository(t),
			hashHelper:       mocks.NewHashHelper(t),
			jwtHelper:        mocks.NewJwtHelper(t),
			mock: func(m *mocks.UserRepository, m2 *mocks.WalletRepository, m3 *mocks.GameRepository, m4 *mocks.HashHelper, m5 *mocks.JwtHelper) {
				m4.On("Encrypt", "Password123").Return("PasswordEncrypted", nil)
				m.On("CreateNewUser", &models.User{
					Name:     "Name",
					Email:    "valid@email.com",
					Password: "PasswordEncrypted",
					Address:  "Badak Street",
					CityID:   1,
					Role:     "user",
				}).Return(&models.User{
					ID:       1,
					Name:     "Name",
					Email:    "valid@email.com",
					Password: "Password123",
					Address:  "Badak Street",
					CityID:   1,
					Role:     "user",
				}, nil)
				m2.On("CreateWallet", &models.Wallet{
					UserID:  1,
					Balance: 0,
				}).Return(nil)
				m3.On("CreateGameUser", &models.GameChance{
					UserID:  1,
					Chance:  0,
					Counter: 0,
				}).Return(nil)
				m5.On("GenerateJwtToken", &dtos.JwtData{
					UserID: 1,
					Email:  "valid@email.com",
				}).Return(nil, errors.New("error"))
			},
			args: args{
				userData: dtos.RegisterRequestDTO{
					Name:     "Name",
					Email:    "valid@email.com",
					Password: "Password123",
					Address:  "Badak Street",
					CityID:   1,
					IsAdmin:  false,
				},
			},
			want:    nil,
			want2:   nil,
			wantErr: true,
		},
		{
			name:             "Success - Register as Admin",
			userRepository:   mocks.NewUserRepository(t),
			walletRepository: mocks.NewWalletRepository(t),
			gameRepository:   mocks.NewGameRepository(t),
			hashHelper:       mocks.NewHashHelper(t),
			jwtHelper:        mocks.NewJwtHelper(t),
			mock: func(m *mocks.UserRepository, m2 *mocks.WalletRepository, m3 *mocks.GameRepository, m4 *mocks.HashHelper, m5 *mocks.JwtHelper) {
				m4.On("Encrypt", "Password123").Return("PasswordEncrypted", nil)
				m.On("CreateNewUser", &models.User{
					Name:     "Name",
					Email:    "valid@email.com",
					Password: "PasswordEncrypted",
					Address:  "Badak Street",
					CityID:   1,
					Role:     "admin",
				}).Return(&models.User{
					ID:       1,
					Name:     "Name",
					Email:    "valid@email.com",
					Password: "Password123",
					Address:  "Badak Street",
					CityID:   1,
					Role:     "admin",
				}, nil)
				m2.On("CreateWallet", &models.Wallet{
					UserID:  1,
					Balance: 0,
				}).Return(nil)
				m3.On("CreateGameUser", &models.GameChance{
					UserID:  1,
					Chance:  0,
					Counter: 0,
				}).Return(nil)
				m5.On("GenerateJwtToken", &dtos.JwtData{
					UserID: 1,
					Email:  "valid@email.com",
				}).Return(nil, nil)
			},
			args: args{
				userData: dtos.RegisterRequestDTO{
					Name:     "Name",
					Email:    "valid@email.com",
					Password: "Password123",
					Address:  "Badak Street",
					CityID:   1,
					IsAdmin:  true,
				},
			},
			want: &models.User{
				ID:       1,
				Name:     "Name",
				Email:    "valid@email.com",
				Password: "Password123",
				Address:  "Badak Street",
				CityID:   1,
				Role:     "admin",
			},
			want2:   nil,
			wantErr: false,
		},
		{
			name:             "Success - Register as User",
			userRepository:   mocks.NewUserRepository(t),
			walletRepository: mocks.NewWalletRepository(t),
			gameRepository:   mocks.NewGameRepository(t),
			hashHelper:       mocks.NewHashHelper(t),
			jwtHelper:        mocks.NewJwtHelper(t),
			mock: func(m *mocks.UserRepository, m2 *mocks.WalletRepository, m3 *mocks.GameRepository, m4 *mocks.HashHelper, m5 *mocks.JwtHelper) {
				m4.On("Encrypt", "Password123").Return("PasswordEncrypted", nil)
				m.On("CreateNewUser", &models.User{
					Name:     "Name",
					Email:    "valid@email.com",
					Password: "PasswordEncrypted",
					Address:  "Badak Street",
					CityID:   1,
					Role:     "user",
				}).Return(&models.User{
					ID:       1,
					Name:     "Name",
					Email:    "valid@email.com",
					Password: "Password123",
					Address:  "Badak Street",
					CityID:   1,
					Role:     "user",
				}, nil)
				m2.On("CreateWallet", &models.Wallet{
					UserID:  1,
					Balance: 0,
				}).Return(nil)
				m3.On("CreateGameUser", &models.GameChance{
					UserID:  1,
					Chance:  0,
					Counter: 0,
				}).Return(nil)
				m5.On("GenerateJwtToken", &dtos.JwtData{
					UserID: 1,
					Email:  "valid@email.com",
				}).Return(nil, nil)
			},
			args: args{
				userData: dtos.RegisterRequestDTO{
					Name:     "Name",
					Email:    "valid@email.com",
					Password: "Password123",
					Address:  "Badak Street",
					CityID:   1,
					IsAdmin:  false,
				},
			},
			want: &models.User{
				ID:       1,
				Name:     "Name",
				Email:    "valid@email.com",
				Password: "Password123",
				Address:  "Badak Street",
				CityID:   1,
				Role:     "user",
			},
			want2:   nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &authService{
				userRepository:   tt.userRepository,
				walletRepository: tt.walletRepository,
				gameRepository:   tt.gameRepository,
				hashHelper:       tt.hashHelper,
				jwtHelper:        tt.jwtHelper,
			}

			tt.mock(tt.userRepository, tt.walletRepository, tt.gameRepository, tt.hashHelper, tt.jwtHelper)

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

func Test_authService_Login(t *testing.T) {
	type args struct {
		userData dtos.LoginRequestDTO
	}
	tests := []struct {
		name           string
		userRepository *mocks.UserRepository
		hashHelper     *mocks.HashHelper
		jwtHelper      *mocks.JwtHelper
		mock           func(m *mocks.UserRepository, m2 *mocks.HashHelper, m3 *mocks.JwtHelper)
		args           args
		want           *models.User
		want2          *string
		wantErr        bool
	}{
		{
			name:           "Failed - GetUserByEmail",
			userRepository: mocks.NewUserRepository(t),
			hashHelper:     mocks.NewHashHelper(t),
			jwtHelper:      mocks.NewJwtHelper(t),
			mock: func(m *mocks.UserRepository, m2 *mocks.HashHelper, m3 *mocks.JwtHelper) {
				m.On("GetUserByEmail", "email123@email.com").Return(nil, errors.New("error"))
			},
			args: args{
				userData: dtos.LoginRequestDTO{
					Email:    "email123@email.com",
					Password: "Password123",
				},
			},
			want:    nil,
			want2:   nil,
			wantErr: true,
		},
		{
			name:           "Failed - Credential Error (Invalid Password)",
			userRepository: mocks.NewUserRepository(t),
			hashHelper:     mocks.NewHashHelper(t),
			jwtHelper:      mocks.NewJwtHelper(t),
			mock: func(m *mocks.UserRepository, m2 *mocks.HashHelper, m3 *mocks.JwtHelper) {
				m.On("GetUserByEmail", "email123@email.com").Return(&models.User{
					ID:       1,
					Email:    "email123@email.com",
					Password: "PasswordEncrypted",
				}, nil)
				m2.On("CompareEncrypt", "Password123", "PasswordEncrypted").Return(false)
			},
			args: args{
				userData: dtos.LoginRequestDTO{
					Email:    "email123@email.com",
					Password: "Password123",
				},
			},
			want:    nil,
			want2:   nil,
			wantErr: true,
		},
		{
			name:           "Failed - GenerateJwtToken",
			userRepository: mocks.NewUserRepository(t),
			hashHelper:     mocks.NewHashHelper(t),
			jwtHelper:      mocks.NewJwtHelper(t),
			mock: func(m *mocks.UserRepository, m2 *mocks.HashHelper, m3 *mocks.JwtHelper) {
				m.On("GetUserByEmail", "email123@email.com").Return(&models.User{
					ID:       1,
					Email:    "email123@email.com",
					Password: "PasswordEncrypted",
				}, nil)
				m2.On("CompareEncrypt", "Password123", "PasswordEncrypted").Return(true)
				m3.On("GenerateJwtToken", &dtos.JwtData{
					UserID: 1,
					Email:  "email123@email.com",
				}).Return(nil, errors.New("error"))
			},
			args: args{
				userData: dtos.LoginRequestDTO{
					Email:    "email123@email.com",
					Password: "Password123",
				},
			},
			want:    nil,
			want2:   nil,
			wantErr: true,
		},

		{
			name:           "Success",
			userRepository: mocks.NewUserRepository(t),
			hashHelper:     mocks.NewHashHelper(t),
			jwtHelper:      mocks.NewJwtHelper(t),
			mock: func(m *mocks.UserRepository, m2 *mocks.HashHelper, m3 *mocks.JwtHelper) {
				m.On("GetUserByEmail", "email123@email.com").Return(&models.User{
					ID:       1,
					Email:    "email123@email.com",
					Password: "PasswordEncrypted",
				}, nil)
				m2.On("CompareEncrypt", "Password123", "PasswordEncrypted").Return(true)
				m3.On("GenerateJwtToken", &dtos.JwtData{
					UserID: 1,
					Email:  "email123@email.com",
				}).Return(nil, nil)
			},
			args: args{
				userData: dtos.LoginRequestDTO{
					Email:    "email123@email.com",
					Password: "Password123",
				},
			},
			want: &models.User{
				ID:       1,
				Email:    "email123@email.com",
				Password: "PasswordEncrypted",
			},
			want2:   nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &authService{
				userRepository: tt.userRepository,
				hashHelper:     tt.hashHelper,
				jwtHelper:      tt.jwtHelper,
			}

			tt.mock(tt.userRepository, tt.hashHelper, tt.jwtHelper)

			got, got2, err := s.Login(tt.args.userData)

			if (err != nil) != tt.wantErr {
				t.Errorf("authService.Login() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("authService.Login() = %v, want %v", got, tt.want)
			}

			if !reflect.DeepEqual(got2, tt.want2) {
				t.Errorf("authService.Login() = %v, want2 %v", got2, tt.want2)
			}
		})
	}
}
