package handlers

import (
	"git.garena.com/sea-labs-id/batch-02/natanael-benediktus/final-project-backend/db"
	r "git.garena.com/sea-labs-id/batch-02/natanael-benediktus/final-project-backend/internal/repositories"
	s "git.garena.com/sea-labs-id/batch-02/natanael-benediktus/final-project-backend/internal/services"
)

type Handler struct {
	authService  s.AuthService
	cityService  s.CityService
	userService  s.UserService
	houseService s.HouseService
}

type HandlerConfig struct {
	AuthService  s.AuthService
	CityService  s.CityService
	UserService  s.UserService
	HouseService s.HouseService
}

func New(c *HandlerConfig) *Handler {
	return &Handler{
		authService:  c.AuthService,
		cityService:  c.CityService,
		userService:  c.UserService,
		houseService: c.HouseService,
	}
}

func InitAllHandlers() *Handler {
	db := db.Get()

	userRepo := r.NewUserRepository(&r.URConfig{
		DB: db,
	})

	cityRepo := r.NewCityRepository(&r.CRConfig{
		DB: db,
	})

	walletRepo := r.NewWalletRepository(&r.WRConfig{
		DB: db,
	})

	houseRepo := r.NewHouseRepository(&r.HRConfig{
		DB: db,
	})

	authService := s.NewAuthService(&s.ASConfig{
		UserRepository:   userRepo,
		CityRepository:   cityRepo,
		WalletRepository: walletRepo,
	})

	cityService := s.NewCityService(&s.CSConfig{
		CityRepository: cityRepo,
	})

	userService := s.NewUserService(&s.USConfig{
		UserRepository: userRepo,
		CityRepository: cityRepo,
	})

	houseService := s.NewHouseService(&s.HSConfig{
		HouseRepository: houseRepo,
	})

	h := New(&HandlerConfig{
		AuthService:  authService,
		CityService:  cityService,
		UserService:  userService,
		HouseService: houseService,
	})

	return h
}
