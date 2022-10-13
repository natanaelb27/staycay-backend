package services

import (
	"git.garena.com/sea-labs-id/batch-02/natanael-benediktus/final-project-backend/internal/dtos"
	"git.garena.com/sea-labs-id/batch-02/natanael-benediktus/final-project-backend/internal/helpers"
	"git.garena.com/sea-labs-id/batch-02/natanael-benediktus/final-project-backend/internal/models"
	r "git.garena.com/sea-labs-id/batch-02/natanael-benediktus/final-project-backend/internal/repositories"
)

type AuthService interface {
	Register(auth dtos.RegisterRequestDTO) (*models.User, *string, error)
}

type authService struct {
	userRepository r.UserRepository
	cityRepository r.CityRepository
}

type ASConfig struct {
	UserRepository r.UserRepository
	CityRepository r.CityRepository
}

func NewAuthService(c *ASConfig) AuthService {
	return &authService{
		userRepository: c.UserRepository,
		cityRepository: c.CityRepository,
	}
}

func (s *authService) Register(auth dtos.RegisterRequestDTO) (*models.User, *string, error) {
	city, err := s.cityRepository.GetCityIDByName(auth.City)

	if err != nil {
		return nil, nil, err
	}

	encryptedPassword, err := helpers.Encrypt(auth.Password)

	if err != nil {
		return nil, nil, err
	}

	user := &models.User{
		Name:     auth.Name,
		Email:    auth.Email,
		Password: encryptedPassword,
		Address:  auth.Address,
		CityID:   city.ID,
		Role:     "user",
	}

	newUser, err := s.userRepository.CreateNewUser(user)

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
