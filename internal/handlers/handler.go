package handlers

import (
	"net/url"
	"strconv"

	"git.garena.com/sea-labs-id/batch-02/natanael-benediktus/final-project-backend/db"
	"git.garena.com/sea-labs-id/batch-02/natanael-benediktus/final-project-backend/internal/dtos"
	"git.garena.com/sea-labs-id/batch-02/natanael-benediktus/final-project-backend/internal/helpers"
	r "git.garena.com/sea-labs-id/batch-02/natanael-benediktus/final-project-backend/internal/repositories"
	s "git.garena.com/sea-labs-id/batch-02/natanael-benediktus/final-project-backend/internal/services"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	authService        s.AuthService
	cityService        s.CityService
	userService        s.UserService
	houseService       s.HouseService
	walletService      s.WalletService
	reservationService s.ReservationService
	pickupService      s.PickupService
	gameService        s.GameService
}

type HandlerConfig struct {
	AuthService        s.AuthService
	CityService        s.CityService
	UserService        s.UserService
	HouseService       s.HouseService
	WalletService      s.WalletService
	ReservationService s.ReservationService
	PickupService      s.PickupService
	GameService        s.GameService
}

func New(c *HandlerConfig) *Handler {
	return &Handler{
		authService:        c.AuthService,
		cityService:        c.CityService,
		userService:        c.UserService,
		houseService:       c.HouseService,
		walletService:      c.WalletService,
		reservationService: c.ReservationService,
		pickupService:      c.PickupService,
		gameService:        c.GameService,
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

	reservationRepo := r.NewReservationRepository(&r.RRConfig{
		DB: db,
	})

	pickupRepo := r.NewPickupRepository(&r.PRConfig{
		DB: db,
	})

	gameRepo := r.NewGameRepository(&r.GRConfig{
		DB: db,
	})

	hashHelper := helpers.NewHashHelper(&helpers.HHConfig{})

	coinGuessHelper := helpers.NewCoinGuessHelper(&helpers.CGHConfig{})

	jwtHelper := helpers.NewJwtHelper(&helpers.JHConfig{})

	authService := s.NewAuthService(&s.ASConfig{
		UserRepository:   userRepo,
		WalletRepository: walletRepo,
		GameRepository:   gameRepo,
		HashHelper:       hashHelper,
		JwtHelper:        jwtHelper,
	})

	cityService := s.NewCityService(&s.CSConfig{
		CityRepository: cityRepo,
	})

	userService := s.NewUserService(&s.USConfig{
		UserRepository: userRepo,
		HashHelper:     hashHelper,
	})

	houseService := s.NewHouseService(&s.HSConfig{
		HouseRepository:       houseRepo,
		UserRepository:        userRepo,
		ReservationRepository: reservationRepo,
	})

	walletService := s.NewWalletService(&s.WSConfig{
		WalletRepository: walletRepo,
		GameRepository:   gameRepo,
	})

	reservationService := s.NewReservationService(&s.RSConfig{
		ReservationRepository: reservationRepo,
		HouseRepository:       houseRepo,
		UserRepository:        userRepo,
		WalletRepository:      walletRepo,
	})

	pickupService := s.NewPickupService(&s.PSConfig{
		PickupRepository:      pickupRepo,
		ReservationRepository: reservationRepo,
		UserRepository:        userRepo,
		WalletRepository:      walletRepo,
	})

	gameService := s.NewGameService(&s.GSConfig{
		GameRepository:   gameRepo,
		WalletRepository: walletRepo,
		CoinGuessHelper:  coinGuessHelper,
	})

	h := New(&HandlerConfig{
		AuthService:        authService,
		CityService:        cityService,
		UserService:        userService,
		HouseService:       houseService,
		WalletService:      walletService,
		ReservationService: reservationService,
		PickupService:      pickupService,
		GameService:        gameService,
	})

	return h
}

var MockUser = dtos.JwtData{
	UserID: 1,
	Email:  "email123@email.com",
}

func MiddlewareMockUser(c *gin.Context) {
	c.Set("user", MockUser)
	c.Next()
}

var MockParamsHouseID = 1
var MockParamsHousePhotoID = 1
var MockParamsUserID = 1

func MiddlewareMockParams(c *gin.Context) {
	c.Params = []gin.Param{
		{Key: "houseId",
			Value: strconv.Itoa(MockParamsHouseID)},
		{Key: "housePhotoId",
			Value: strconv.Itoa(MockParamsHousePhotoID)},
		{Key: "userId",
			Value: strconv.Itoa(MockParamsUserID)},
	}
	c.Next()
}

var MockQueryParams = map[string]interface{}{
	"key": "value",
}

func MiddlewareMockQueryParams(c *gin.Context) {
	v := url.URL{
		RawQuery: "key=value",
	}
	c.Request.URL = &v
	c.Next()
}
