package services

import (
	"git.garena.com/sea-labs-id/batch-02/natanael-benediktus/final-project-backend/internal/dtos"
	"git.garena.com/sea-labs-id/batch-02/natanael-benediktus/final-project-backend/internal/helpers"
	"git.garena.com/sea-labs-id/batch-02/natanael-benediktus/final-project-backend/internal/models"
	r "git.garena.com/sea-labs-id/batch-02/natanael-benediktus/final-project-backend/internal/repositories"
	"gorm.io/gorm"
)

type AuthService interface {
	Register(userData dtos.RegisterRequestDTO) (*models.User, *string, error)
	Login(userData dtos.LoginRequestDTO) (*models.User, *string, error)
}

type authService struct {
	userRepository   r.UserRepository
	cityRepository   r.CityRepository
	walletRepository r.WalletRepository
}

type ASConfig struct {
	UserRepository   r.UserRepository
	CityRepository   r.CityRepository
	WalletRepository r.WalletRepository
}

func NewAuthService(c *ASConfig) AuthService {
	return &authService{
		userRepository:   c.UserRepository,
		cityRepository:   c.CityRepository,
		walletRepository: c.WalletRepository,
	}
}

func (s *authService) Register(userData dtos.RegisterRequestDTO) (*models.User, *string, error) {
	city, err := s.cityRepository.GetCityIDByName(userData.City)

	if err != nil {
		return nil, nil, err
	}

	encryptedPassword, err := helpers.Encrypt(userData.Password)

	if err != nil {
		return nil, nil, err
	}

	user := &models.User{
		Name:     userData.Name,
		Email:    userData.Email,
		Password: encryptedPassword,
		Address:  userData.Address,
		CityID:   city.ID,
		Role:     "user",
	}

	newUser, err := s.userRepository.CreateNewUser(user)

	if err != nil {
		return nil, nil, err
	}

	wallet := &models.Wallet{
		UserID:  newUser.ID,
		Balance: 0,
	}

	err = s.walletRepository.CreateWallet(wallet)

	if err != nil {
		return nil, nil, err
	}

	dataToJwt := &dtos.JwtData{
		ID:    newUser.ID,
		Email: newUser.Email,
	}

	token, err := helpers.GenerateJwtToken(dataToJwt)

	if err != nil {
		return nil, nil, err
	}

	return newUser, token, nil
}

func (s *authService) Login(userData dtos.LoginRequestDTO) (*models.User, *string, error) {
	user, err := s.userRepository.GetUserByEmail(userData.Email)

	if err != nil {
		return nil, nil, err
	}

	if !helpers.CompareEncrypt(userData.Password, user.Password) {
		return nil, nil, gorm.ErrRecordNotFound
	}

	dataToJwt := &dtos.JwtData{
		ID:    user.ID,
		Email: user.Email,
	}

	token, err := helpers.GenerateJwtToken(dataToJwt)

	if err != nil {
		return nil, nil, err
	}

	return user, token, nil
}
