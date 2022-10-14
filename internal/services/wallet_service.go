package services

import (
	"math"

	e "git.garena.com/sea-labs-id/batch-02/natanael-benediktus/final-project-backend/internal/errors"
	"git.garena.com/sea-labs-id/batch-02/natanael-benediktus/final-project-backend/internal/models"
	r "git.garena.com/sea-labs-id/batch-02/natanael-benediktus/final-project-backend/internal/repositories"
)

type WalletService interface {
	TopUpWallet(userID, amount int) (*models.Wallet, error)
}

type walletService struct {
	walletRepository r.WalletRepository
	gameRepository   r.GameRepository
}

type WSConfig struct {
	WalletRepository r.WalletRepository
	GameRepository   r.GameRepository
}

func NewWalletService(c *WSConfig) WalletService {
	return &walletService{
		walletRepository: c.WalletRepository,
		gameRepository:   c.GameRepository,
	}
}

func (s *walletService) TopUpWallet(userID, amount int) (*models.Wallet, error) {
	if amount < 10000 {
		return nil, e.ErrTopUpTooSmallError
	}

	wallet, err := s.walletRepository.AddWalletBalance(userID, amount)

	if err != nil {
		return nil, err
	}

	minTopupForGame := 500000.0

	gameChanceAmount := int(math.Floor(float64(amount) / minTopupForGame))

	if gameChanceAmount > 0 {
		err = s.gameRepository.AddGameChance(userID, gameChanceAmount)

		if err != nil {
			return nil, err
		}
	}

	return wallet, nil
}
