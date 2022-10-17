package services

import (
	"git.garena.com/sea-labs-id/batch-02/natanael-benediktus/final-project-backend/internal/dtos"
	e "git.garena.com/sea-labs-id/batch-02/natanael-benediktus/final-project-backend/internal/errors"
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
	gameRepository   r.GameRepository
}

type ASConfig struct {
	UserRepository   r.UserRepository
	CityRepository   r.CityRepository
	WalletRepository r.WalletRepository
	GameRepository   r.GameRepository
}

func NewAuthService(c *ASConfig) AuthService {
	return &authService{
		userRepository:   c.UserRepository,
		cityRepository:   c.CityRepository,
		walletRepository: c.WalletRepository,
		gameRepository:   c.GameRepository,
	}
}

func (s *authService) Register(userData dtos.RegisterRequestDTO) (*models.User, *string, error) {
	if !helpers.VerifyEMailAddress(userData.Email) {
		return nil, nil, e.ErrInvalidEmailInputError
	}

	if !helpers.VerifyPassword(userData.Password) {
		return nil, nil, e.ErrInvalidPasswordInputError
	}

	// city, err := s.cityRepository.GetCityIDByName(userData.CityName)

	// if err != nil {
	// 	return nil, nil, err
	// }

	encryptedPassword, err := helpers.Encrypt(userData.Password)

	if err != nil {
		return nil, nil, err
	}

	print(userData.IsAdmin)

	role := ""
	if userData.IsAdmin {
		role = "admin"
	} else {
		role = "user"
	}

	user := &models.User{
		Name:     userData.Name,
		Email:    userData.Email,
		Password: encryptedPassword,
		Address:  userData.Address,
		CityID:   userData.CityID,
		Role:     role,
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

	game := &models.GameChance{
		UserID:  newUser.ID,
		Chance:  0,
		Counter: 0,
	}

	err = s.gameRepository.CreateGameUser(game)

	if err != nil {
		return nil, nil, err
	}

	dataToJwt := &dtos.JwtData{
		UserID: newUser.ID,
		Email:  newUser.Email,
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
		UserID: user.ID,
		Email:  user.Email,
	}

	token, err := helpers.GenerateJwtToken(dataToJwt)

	if err != nil {
		return nil, nil, err
	}

	return user, token, nil
}
