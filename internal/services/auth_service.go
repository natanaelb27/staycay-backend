package services

import (
	"github.com/natanaelb27/staycay-backend/internal/dtos"
	e "github.com/natanaelb27/staycay-backend/internal/errors"
	"github.com/natanaelb27/staycay-backend/internal/helpers"
	"github.com/natanaelb27/staycay-backend/internal/models"
	r "github.com/natanaelb27/staycay-backend/internal/repositories"
)

type AuthService interface {
	Register(userData *dtos.RegisterRequestDTO) (*models.User, *string, error)
	Login(userData *dtos.LoginRequestDTO) (*models.User, *string, error)
}

type authService struct {
	userRepository   r.UserRepository
	walletRepository r.WalletRepository
	gameRepository   r.GameRepository
	hashHelper       helpers.HashHelper
	jwtHelper        helpers.JwtHelper
}

type ASConfig struct {
	UserRepository   r.UserRepository
	WalletRepository r.WalletRepository
	GameRepository   r.GameRepository
	HashHelper       helpers.HashHelper
	JwtHelper        helpers.JwtHelper
}

func NewAuthService(c *ASConfig) AuthService {
	return &authService{
		userRepository:   c.UserRepository,
		walletRepository: c.WalletRepository,
		gameRepository:   c.GameRepository,
		hashHelper:       c.HashHelper,
		jwtHelper:        c.JwtHelper,
	}
}

func (s *authService) Register(userData *dtos.RegisterRequestDTO) (*models.User, *string, error) {
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

	encryptedPassword, err := s.hashHelper.Encrypt(userData.Password)

	if err != nil {
		return nil, nil, err
	}

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

	token, err := s.jwtHelper.GenerateJwtToken(dataToJwt)

	if err != nil {
		return nil, nil, err
	}

	return newUser, token, nil
}

func (s *authService) Login(userData *dtos.LoginRequestDTO) (*models.User, *string, error) {
	user, err := s.userRepository.GetUserByEmail(userData.Email)

	if err != nil {
		return nil, nil, err
	}

	if !s.hashHelper.CompareEncrypt(userData.Password, user.Password) {
		return nil, nil, e.ErrCredentialError
	}

	dataToJwt := &dtos.JwtData{
		UserID: user.ID,
		Email:  user.Email,
	}

	token, err := s.jwtHelper.GenerateJwtToken(dataToJwt)

	if err != nil {
		return nil, nil, err
	}

	return user, token, nil
}
