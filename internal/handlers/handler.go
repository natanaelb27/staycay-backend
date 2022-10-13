package handlers

import (
	"git.garena.com/sea-labs-id/batch-02/natanael-benediktus/final-project-backend/db"
	r "git.garena.com/sea-labs-id/batch-02/natanael-benediktus/final-project-backend/internal/repositories"
	"git.garena.com/sea-labs-id/batch-02/natanael-benediktus/final-project-backend/internal/services"
	s "git.garena.com/sea-labs-id/batch-02/natanael-benediktus/final-project-backend/internal/services"
)

type Handler struct {
	authService services.AuthService
}

type HandlerConfig struct {
	AuthService services.AuthService
}

func New(c *HandlerConfig) *Handler {
	return &Handler{
		authService: c.AuthService,
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

	authService := s.NewAuthService(&s.ASConfig{
		UserRepository:   userRepo,
		CityRepository:   cityRepo,
		WalletRepository: walletRepo,
	})

	h := New(&HandlerConfig{
		AuthService: authService,
	})

	return h
}
