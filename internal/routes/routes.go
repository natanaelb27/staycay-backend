package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/natanaelb27/staycay-backend/internal/handlers"
	"github.com/natanaelb27/staycay-backend/internal/helpers"
	"github.com/natanaelb27/staycay-backend/internal/middlewares"
)

func InitAllRoutes(r *gin.Engine, h *handlers.Handler) {
	r.GET("", func(c *gin.Context) {
		c.JSON(http.StatusOK, helpers.BuildSuccessResponse(http.StatusOK, "House Rental StayCay App", nil))
	})

	r.POST("/register", h.RegisterHandler)
	r.POST("/login", h.LoginHandler)

	city := r.Group("/city")
	city.GET("", h.GetAllCitiesHandler)

	house := r.Group("/house")
	house.GET("", h.GetAllHousesHandler)
	house.GET("/:houseId", h.GetHouseByIDHandler)
	house.GET("/status", h.GetAllHouseStatusesHandler)
	house.GET("/reservation/:houseId", h.GetHouseReservationsByHouseHandler)

	user := r.Group("/user")
	user.GET("/:userId", h.GetUserProfileByUserIDHandler)
	user.Use(middlewares.AuthorizeJWT)
	user.GET("", h.GetUserProfileHandler)
	user.PATCH("", h.UpdateUserProfileHandler)
	user.PATCH("/role", h.UpdateUserRoleHandler)
	user.GET("/reservation", h.GetHouseReservationsByUserHandler)

	pickup := r.Group("/pickup")
	pickup.GET("/status", h.GetAllPickupStatusesHandler)
	pickup.GET("", h.GetAllPickupDetailsHandler)
	pickup.Use(middlewares.AuthorizeJWT)
	pickup.GET("/reservation/:reservationId", h.GetPickupDetailsByReservationHandler)
	pickup.POST("/reservation/:reservationId", h.RequestPickupHandler)

	topup := r.Group("/topup")
	topup.Use(middlewares.AuthorizeJWT)
	topup.POST("", h.TopUpWalletHandler)

	host := r.Group("/host")
	host.Use(middlewares.AuthorizeJWT)
	host.GET("house", h.GetHostHousesHandler)
	host.GET("/reservation", h.GetHouseReservationsByHostHandler)
	host.POST("/house", h.AddNewHouseHandler)
	host.PATCH("/house/:houseId", h.UpdateHouseHandler)
	host.DELETE("/house/:houseId", h.DeleteHouseHandler)
	host.POST("/house/photo/:houseId", h.AddHousePhotoHandler)
	host.DELETE("/house/photo/:housePhotoId", h.DeleteHousePhotoHandler)

	reservation := r.Group("/reservation")
	reservation.Use(middlewares.AuthorizeJWT)
	reservation.GET("/:reservationId", h.GetHouseReservationsByIDHandler)
	reservation.POST("", h.AddReservationHandler)

	game := r.Group("/game")
	game.Use(middlewares.AuthorizeJWT)
	game.GET("", h.GetGameHandler)
	game.POST("", h.PlayGameHandler)

	admin := r.Group("/admin")
	admin.Use(middlewares.AuthorizeJWT)
	admin.PATCH("/pickup/:pickupId", h.AdminUpdatePickupStatusHandler)
}
