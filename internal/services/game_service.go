package services

import (
	e "git.garena.com/sea-labs-id/batch-02/natanael-benediktus/final-project-backend/internal/errors"
	"git.garena.com/sea-labs-id/batch-02/natanael-benediktus/final-project-backend/internal/helpers"
	"git.garena.com/sea-labs-id/batch-02/natanael-benediktus/final-project-backend/internal/models"
	r "git.garena.com/sea-labs-id/batch-02/natanael-benediktus/final-project-backend/internal/repositories"
)

type GameService interface {
	PlayGame(userGuess string, userID int) (string, int, error)
	GetGame(userID int) (*models.GameChance, error)
}

type gameService struct {
	gameRepository   r.GameRepository
	walletRepository r.WalletRepository
}

type GSConfig struct {
	GameRepository   r.GameRepository
	WalletRepository r.WalletRepository
}

func NewGameService(c *GSConfig) GameService {
	return &gameService{
		gameRepository:   c.GameRepository,
		walletRepository: c.WalletRepository,
	}
}

func (s *gameService) PlayGame(userGuess string, userID int) (string, int, error) {
	gameData, err := s.gameRepository.GetGameData(userID)

	if err != nil {
		return "", 0, err
	}

	if gameData.Chance <= 0 {
		return "", 0, e.ErrInsufficientGameChanceError
	}

	maxAmount := 50000
	minAmount := 10000

	coin := []string{"heads", "tails"}
	randomIndex := helpers.Randomizer(len(coin))
	pick := coin[randomIndex]

	amountWon := 0
	if pick == userGuess {
		amountWon = maxAmount
	} else {
		amountWon = minAmount
	}

	_, err = s.walletRepository.AddWalletBalance(userID, amountWon)

	if err != nil {
		return "", 0, err
	}

	err = s.gameRepository.SubtractGameChance(userID, 1)

	if err != nil {
		return "", 0, err
	}

	if gameData.Counter == 9 {
		err = s.gameRepository.AddGameChance(userID, 1)

		if err != nil {
			return "", 0, err
		}

		err = s.gameRepository.ResetGameChanceCounter(userID)

		if err != nil {
			return "", 0, err
		}
	} else {
		err = s.gameRepository.AddGameChanceCounter(userID, 1)

		if err != nil {
			return "", 0, err
		}
	}

	return pick, amountWon, nil
}

func (s *gameService) GetGame(userID int) (*models.GameChance, error) {
	gameData, err := s.gameRepository.GetGameData(userID)

	if err != nil {
		return nil, err
	}

	return gameData, nil
}
