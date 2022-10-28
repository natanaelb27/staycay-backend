package services

import (
	"errors"
	"reflect"
	"testing"

	"git.garena.com/sea-labs-id/batch-02/natanael-benediktus/final-project-backend/internal/mocks"
	"git.garena.com/sea-labs-id/batch-02/natanael-benediktus/final-project-backend/internal/models"
	"github.com/stretchr/testify/assert"
)

func Test_gameService_NewGameService(t *testing.T) {
	gameRepository := mocks.NewGameRepository(t)
	walletRepository := mocks.NewWalletRepository(t)
	gameService := NewGameService(
		&GSConfig{
			GameRepository:   gameRepository,
			WalletRepository: walletRepository,
		},
	)

	assert.NotNil(t, gameService)
}

func Test_gameService_PlayGame(t *testing.T) {
	type args struct {
		userGuess string
		userID    int
	}
	tests := []struct {
		name             string
		gameRepository   *mocks.GameRepository
		walletRepository *mocks.WalletRepository
		mock             func(m *mocks.GameRepository, m2 *mocks.WalletRepository)
		args             args
		want             string
		want2            int
		wantErr          bool
	}{
		{
			name:             "Failed - GetGameData",
			gameRepository:   mocks.NewGameRepository(t),
			walletRepository: mocks.NewWalletRepository(t),
			mock: func(m *mocks.GameRepository, m2 *mocks.WalletRepository) {
				m.On("GetGameData", 1).Return(nil, errors.New("error"))
			},
			args: args{
				userGuess: "heads",
				userID:    1,
			},
			want:    "",
			want2:   0,
			wantErr: true,
		},
		{
			name:             "Failed - Insufficient Game Chance",
			gameRepository:   mocks.NewGameRepository(t),
			walletRepository: mocks.NewWalletRepository(t),
			mock: func(m *mocks.GameRepository, m2 *mocks.WalletRepository) {
				m.On("GetGameData", 1).Return(&models.GameChance{
					Chance: 0,
				}, nil)
			},
			args: args{
				userGuess: "heads",
				userID:    1,
			},
			want:    "",
			want2:   0,
			wantErr: true,
		},
		{
			name:             "Failed - AddWalletBalance",
			gameRepository:   mocks.NewGameRepository(t),
			walletRepository: mocks.NewWalletRepository(t),
			mock: func(m *mocks.GameRepository, m2 *mocks.WalletRepository) {
				m.On("GetGameData", 1).Return(&models.GameChance{
					Chance: 10,
				}, nil)
				m2.On("AddWalletBalance", 1, 10000).Return(nil, errors.New("error"))
			},
			args: args{
				userGuess: "heads",
				userID:    1,
			},
			want:    "",
			want2:   0,
			wantErr: true,
		},
		{
			name:             "Failed - SubtractGameChance",
			gameRepository:   mocks.NewGameRepository(t),
			walletRepository: mocks.NewWalletRepository(t),
			mock: func(m *mocks.GameRepository, m2 *mocks.WalletRepository) {
				m.On("GetGameData", 1).Return(&models.GameChance{
					Chance: 10,
				}, nil)
				m2.On("AddWalletBalance", 1, 10000).Return(&models.Wallet{}, nil)
				m.On("SubtractGameChance", 1, 1).Return(errors.New("error"))
			},
			args: args{
				userGuess: "heads",
				userID:    1,
			},
			want:    "",
			want2:   0,
			wantErr: true,
		},
		{
			name:             "Failed - AddGameChance When Counter is 9",
			gameRepository:   mocks.NewGameRepository(t),
			walletRepository: mocks.NewWalletRepository(t),
			mock: func(m *mocks.GameRepository, m2 *mocks.WalletRepository) {
				m.On("GetGameData", 1).Return(&models.GameChance{
					Chance:  10,
					Counter: 9,
				}, nil)
				m2.On("AddWalletBalance", 1, 10000).Return(&models.Wallet{}, nil)
				m.On("SubtractGameChance", 1, 1).Return(nil)
				m.On("AddGameChance", 1, 1).Return(errors.New("error"))
			},
			args: args{
				userGuess: "heads",
				userID:    1,
			},
			want:    "",
			want2:   0,
			wantErr: true,
		},
		{
			name:             "Failed - ResetGameChanceCounter When Counter is 9",
			gameRepository:   mocks.NewGameRepository(t),
			walletRepository: mocks.NewWalletRepository(t),
			mock: func(m *mocks.GameRepository, m2 *mocks.WalletRepository) {
				m.On("GetGameData", 1).Return(&models.GameChance{
					Chance:  10,
					Counter: 9,
				}, nil)
				m2.On("AddWalletBalance", 1, 10000).Return(&models.Wallet{}, nil)
				m.On("SubtractGameChance", 1, 1).Return(nil)
				m.On("AddGameChance", 1, 1).Return(nil)
				m.On("ResetGameChanceCounter", 1).Return(errors.New("error"))
			},
			args: args{
				userGuess: "heads",
				userID:    1,
			},
			want:    "",
			want2:   0,
			wantErr: true,
		},
		{
			name:             "Failed - AddGameChanceCounter",
			gameRepository:   mocks.NewGameRepository(t),
			walletRepository: mocks.NewWalletRepository(t),
			mock: func(m *mocks.GameRepository, m2 *mocks.WalletRepository) {
				m.On("GetGameData", 1).Return(&models.GameChance{
					Chance:  10,
					Counter: 1,
				}, nil)
				m2.On("AddWalletBalance", 1, 10000).Return(&models.Wallet{}, nil)
				m.On("SubtractGameChance", 1, 1).Return(nil)
				m.On("AddGameChanceCounter", 1, 1).Return(errors.New("error"))
			},
			args: args{
				userGuess: "heads",
				userID:    1,
			},
			want:    "",
			want2:   0,
			wantErr: true,
		},
		{
			name:             "Success",
			gameRepository:   mocks.NewGameRepository(t),
			walletRepository: mocks.NewWalletRepository(t),
			mock: func(m *mocks.GameRepository, m2 *mocks.WalletRepository) {
				m.On("GetGameData", 1).Return(&models.GameChance{
					Chance:  10,
					Counter: 1,
				}, nil)
				m2.On("AddWalletBalance", 1, 50000).Return(&models.Wallet{}, nil)
				m.On("SubtractGameChance", 1, 1).Return(nil)
				m.On("AddGameChanceCounter", 1, 1).Return(nil)
			},
			args: args{
				userGuess: "heads",
				userID:    1,
			},
			want:    "heads",
			want2:   50000,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &gameService{
				gameRepository:   tt.gameRepository,
				walletRepository: tt.walletRepository,
			}

			tt.mock(tt.gameRepository, tt.walletRepository)

			got, got2, err := s.PlayGame(tt.args.userGuess, tt.args.userID)

			if (err != nil) != tt.wantErr {
				t.Errorf("gameService.PlayGame() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("gameService.PlayGame() = %v, want %v", got, tt.want)
			}

			if !reflect.DeepEqual(got2, tt.want2) {
				t.Errorf("gameService.PlayGame() = %v, want2 %v", got, tt.want2)
			}
		})
	}
}

func Test_gameService_GetGame(t *testing.T) {
	type args struct {
		userID int
	}
	tests := []struct {
		name           string
		gameRepository *mocks.GameRepository
		mock           func(m *mocks.GameRepository)
		args           args
		want           *models.GameChance
		wantErr        bool
	}{
		{
			name:           "Failed - GetGameData",
			gameRepository: mocks.NewGameRepository(t),
			mock: func(m *mocks.GameRepository) {
				m.On("GetGameData", 1).Return(nil, errors.New("error"))
			},
			args: args{
				userID: 1,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name:           "Success",
			gameRepository: mocks.NewGameRepository(t),
			mock: func(m *mocks.GameRepository) {
				m.On("GetGameData", 1).Return(&models.GameChance{}, nil)
			},
			args: args{
				userID: 1,
			},
			want:    &models.GameChance{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &gameService{
				gameRepository: tt.gameRepository,
			}

			tt.mock(tt.gameRepository)

			got, err := s.GetGame(tt.args.userID)

			if (err != nil) != tt.wantErr {
				t.Errorf("gameService.GetGame() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("gameService.GetGame() = %v, want %v", got, tt.want)
			}
		})
	}
}
