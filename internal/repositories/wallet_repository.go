package repositories

import (
	"git.garena.com/sea-labs-id/batch-02/natanael-benediktus/final-project-backend/internal/models"
	"gorm.io/gorm"
)

type WalletRepository interface {
	CreateWallet(wallet *models.Wallet) error
}

type walletRepository struct {
	db *gorm.DB
}

type WRConfig struct {
	DB *gorm.DB
}

func NewWalletRepository(c *WRConfig) WalletRepository {
	return &walletRepository{
		db: c.DB,
	}
}

func (r *walletRepository) CreateWallet(wallet *models.Wallet) error {
	result := r.db.Create(&wallet)

	if result.Error != nil {
		print(result.Error)
		return result.Error
	}

	return nil
}
