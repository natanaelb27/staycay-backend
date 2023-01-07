package repositories

import (
	"github.com/natanaelb27/staycay-backend/internal/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type WalletRepository interface {
	CreateWallet(wallet *models.Wallet) error
	AddWalletBalance(userID int, amount int) (*models.Wallet, error)
	SubtractWalletBalance(userID int, amount int) (*models.Wallet, error)
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
		return result.Error
	}

	return nil
}

func (r *walletRepository) AddWalletBalance(userID int, amount int) (*models.Wallet, error) {
	var wallet models.Wallet

	result := r.db.
		Model(&wallet).
		Clauses(clause.Returning{}).
		Where("user_id = ?", userID).
		Update("balance", gorm.Expr("balance + ?", amount))

	if result.Error != nil {
		return nil, result.Error
	}

	return &wallet, nil

}

func (r *walletRepository) SubtractWalletBalance(userID int, amount int) (*models.Wallet, error) {
	var wallet models.Wallet

	result := r.db.
		Model(&wallet).
		Clauses(clause.Returning{}).
		Where("user_id = ?", userID).
		Update("balance", gorm.Expr("balance - ?", amount))

	if result.Error != nil {
		return nil, result.Error
	}

	return &wallet, nil

}
