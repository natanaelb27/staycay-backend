package services

import (
	"git.garena.com/sea-labs-id/batch-02/natanael-benediktus/final-project-backend/internal/dtos"
	e "git.garena.com/sea-labs-id/batch-02/natanael-benediktus/final-project-backend/internal/errors"
	"git.garena.com/sea-labs-id/batch-02/natanael-benediktus/final-project-backend/internal/helpers"
	"git.garena.com/sea-labs-id/batch-02/natanael-benediktus/final-project-backend/internal/models"
	r "git.garena.com/sea-labs-id/batch-02/natanael-benediktus/final-project-backend/internal/repositories"
)

type UserService interface {
	GetUserProfile(userID int) (*dtos.UserDetailsDTO, error)
	GetUserProfileByUserID(userID int) (*dtos.UserDetailsDTO, error)
	UpdateUserRoleToHost(userID int) (*models.User, error)
	UpdateUserProfile(userID int, request dtos.UpdateUserProfileRequestDTO) (*models.User, error)
}

type userService struct {
	userRepository r.UserRepository
	cityRepository r.CityRepository
}

type USConfig struct {
	UserRepository r.UserRepository
	CityRepository r.CityRepository
}

func NewUserService(c *USConfig) UserService {
	return &userService{
		userRepository: c.UserRepository,
		cityRepository: c.CityRepository,
	}
}

func (s *userService) GetUserProfile(userID int) (*dtos.UserDetailsDTO, error) {
	userDetails, err := s.userRepository.GetUserDetails(userID)

	if err != nil {
		return nil, err
	}

	return userDetails, nil
}

func (s *userService) GetUserProfileByUserID(userID int) (*dtos.UserDetailsDTO, error) {
	userDetails, err := s.userRepository.GetUserDetails(userID)

	if err != nil {
		return nil, err
	}

	return userDetails, nil
}

func (s *userService) UpdateUserRoleToHost(userID int) (*models.User, error) {
	userDetails, err := s.userRepository.UpdateUserRole(userID, "host")

	if err != nil {
		return nil, err
	}

	return userDetails, nil
}

func (s *userService) UpdateUserProfile(userID int, request dtos.UpdateUserProfileRequestDTO) (*models.User, error) {
	user, err := s.userRepository.GetUserByID(userID)

	if err != nil {
		return nil, err
	}

	encryptedPassword := ""
	if request.Password != "" {
		if !helpers.CompareEncrypt(request.OldPassword, user.Password) {
			return nil, e.ErrOldPasswordNotMatchError
		}

		encryptedPassword, err = helpers.Encrypt(request.Password)
		if err != nil {
			return nil, err
		}
	}

	city := &models.City{}
	if request.CityName != "" {
		city, err = s.cityRepository.GetCityIDByName(request.CityName)

		if err != nil {
			return nil, err
		}
	}

	newData := &models.User{
		Name:     request.Name,
		Password: encryptedPassword,
		Address:  request.Address,
		CityID:   city.ID,
	}

	userDetails, err := s.userRepository.UpdateUser(userID, newData)

	if err != nil {
		return nil, err
	}

	return userDetails, nil
}
