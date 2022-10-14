package repositories

import (
	"git.garena.com/sea-labs-id/batch-02/natanael-benediktus/final-project-backend/internal/models"
	"gorm.io/gorm"
)

type GameRepository interface {
	CreateGameUser(game *models.GameChance) error
	AddGameChance(userID, gameChanceAmount int) error
}

type gameRepository struct {
	db *gorm.DB
}

type GRConfig struct {
	DB *gorm.DB
}

func NewGameRepository(c *GRConfig) GameRepository {
	return &gameRepository{
		db: c.DB,
	}
}

func (r *gameRepository) CreateGameUser(game *models.GameChance) error {
	result := r.db.Create(&game)

	if result.Error != nil {
		print(result.Error)
		return result.Error
	}

	return nil
}

func (r *gameRepository) AddGameChance(userID, gameChanceAmount int) error {
	result := r.db.
		Model(&models.GameChance{}).
		Where("user_id = ?", userID).
		Update("chance", gorm.Expr("chance + ?", gameChanceAmount))

	if result.Error != nil {
		return result.Error
	}

	return nil
}
