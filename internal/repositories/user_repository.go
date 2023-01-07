package repositories

import (
	"github.com/natanaelb27/staycay-backend/internal/dtos"
	e "github.com/natanaelb27/staycay-backend/internal/errors"
	"github.com/natanaelb27/staycay-backend/internal/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type UserRepository interface {
	CreateNewUser(user *models.User) (*models.User, error)
	GetUserDetails(userID int) (*dtos.UserDetailsDTO, error)
	GetUserByID(userID int) (*models.User, error)
	GetUserByEmail(email string) (*models.User, error)
	UpdateUserRole(userID int, newRole string) (*models.User, error)
	UpdateUser(userID int, userNewData *models.User) (*models.User, error)
}

type userRepository struct {
	db *gorm.DB
}

type URConfig struct {
	DB *gorm.DB
}

func NewUserRepository(c *URConfig) UserRepository {
	return &userRepository{
		db: c.DB,
	}
}

func (r *userRepository) CreateNewUser(user *models.User) (*models.User, error) {
	result := r.db.Where("email = ?", user.Email).FirstOrCreate(&user)

	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, e.ErrDuplicateEmailError
	}

	return user, nil
}

func (r *userRepository) GetUserDetails(userID int) (*dtos.UserDetailsDTO, error) {
	var userDetails *dtos.UserDetailsDTO

	result := r.db.Model(&models.User{}).
		Select("users.id AS user_id, users.name, users.email, users.address, cities.id AS city_id, cities.name AS city_name, users.role, wallets.id AS wallet_id, wallets.balance").
		Joins("LEFT JOIN cities ON users.city_id = cities.id").
		Joins("LEFT JOIN wallets ON users.id = wallets.user_id").
		Where("users.id = ?", userID).
		First(&userDetails)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, e.ErrUserNotFoundError
		}

		return nil, result.Error
	}

	return userDetails, nil
}

func (r *userRepository) GetUserByID(userID int) (*models.User, error) {
	var user *models.User

	result := r.db.
		Where("id = ?", userID).
		First(&user)

	// if result.RowsAffected == 0 {
	// 	return nil, gorm.ErrRecordNotFound
	// }

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, e.ErrUserNotFoundError
		}
		return nil, result.Error
	}

	return user, nil
}

func (r *userRepository) GetUserByEmail(email string) (*models.User, error) {
	var user *models.User

	result := r.db.
		Where("email = ?", email).
		First(&user)

	// if result.RowsAffected == 0 {
	// 	return nil, gorm.ErrRecordNotFound
	// }

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, e.ErrCredentialError
		}

		return nil, result.Error
	}

	return user, nil
}

func (r *userRepository) UpdateUser(userID int, userNewData *models.User) (*models.User, error) {
	var user models.User

	// result := r.db.
	// 	Model(&user).
	// 	Clauses(clause.Returning{}).
	// 	Where("id = ?", userID).
	// 	Updates(
	// 		map[string]interface{}{
	// 			"name":     userNewData.Name,
	// 			"password": userNewData.Password,
	// 			"address":  userNewData.Address,
	// 			"city_id":  userNewData.CityID,
	// 		},
	// 	)

	result := r.db.
		Model(&user).
		Clauses(clause.Returning{}).
		Where("id = ?", userID).
		Updates(&userNewData)

	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

func (r *userRepository) UpdateUserRole(userID int, newRole string) (*models.User, error) {
	var user models.User

	result := r.db.
		Model(&user).
		Clauses(clause.Returning{}).
		Where("id = ?", userID).
		Update("role", newRole)

	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}
