package services

import (
	"errors"
	"reflect"
	"testing"

	"git.garena.com/sea-labs-id/batch-02/natanael-benediktus/final-project-backend/internal/mocks"
	"git.garena.com/sea-labs-id/batch-02/natanael-benediktus/final-project-backend/internal/models"
	"github.com/stretchr/testify/assert"
)

func Test_walletService_NewWalletService(t *testing.T) {
	walletRepository := mocks.NewWalletRepository(t)
	gameRepository := mocks.NewGameRepository(t)
	walletService := NewWalletService(
		&WSConfig{
			WalletRepository: walletRepository,
			GameRepository:   gameRepository,
		},
	)

	assert.NotNil(t, walletService)
}

func Test_walletService_TopUpWallet(t *testing.T) {
	type args struct {
		userID int
		amount int
	}
	tests := []struct {
		name             string
		walletRepository *mocks.WalletRepository
		gameRepository   *mocks.GameRepository
		mock             func(m *mocks.WalletRepository, m2 *mocks.GameRepository)
		args             args
		want             *models.Wallet
		wantErr          bool
	}{
		{
			name:             "Failed - Top Up Amount Too Small",
			walletRepository: mocks.NewWalletRepository(t),
			gameRepository:   mocks.NewGameRepository(t),
			mock: func(m *mocks.WalletRepository, m2 *mocks.GameRepository) {
			},
			args: args{
				userID: 1,
				amount: 5000,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name:             "Failed - AddWalletBalance",
			walletRepository: mocks.NewWalletRepository(t),
			gameRepository:   mocks.NewGameRepository(t),
			mock: func(m *mocks.WalletRepository, m2 *mocks.GameRepository) {
				m.On("AddWalletBalance", 1, 500000).Return(nil, errors.New("error"))
			},
			args: args{
				userID: 1,
				amount: 500000,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name:             "Failed - AddGameChance",
			walletRepository: mocks.NewWalletRepository(t),
			gameRepository:   mocks.NewGameRepository(t),
			mock: func(m *mocks.WalletRepository, m2 *mocks.GameRepository) {
				m.On("AddWalletBalance", 1, 500000).Return(&models.Wallet{}, nil)
				m2.On("AddGameChance", 1, 1).Return(errors.New("error"))
			},
			args: args{
				userID: 1,
				amount: 500000,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name:             "Success",
			walletRepository: mocks.NewWalletRepository(t),
			gameRepository:   mocks.NewGameRepository(t),
			mock: func(m *mocks.WalletRepository, m2 *mocks.GameRepository) {
				m.On("AddWalletBalance", 1, 500000).Return(&models.Wallet{}, nil)
				m2.On("AddGameChance", 1, 1).Return(nil)
			},
			args: args{
				userID: 1,
				amount: 500000,
			},
			want:    &models.Wallet{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &walletService{
				walletRepository: tt.walletRepository,
				gameRepository:   tt.gameRepository,
			}

			tt.mock(tt.walletRepository, tt.gameRepository)

			got, err := s.TopUpWallet(tt.args.userID, tt.args.amount)

			if (err != nil) != tt.wantErr {
				t.Errorf("walletService.TopUpWallet() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("walletService.TopUpWallet() = %v, want %v", got, tt.want)
			}

		})
	}
}
