package routes

import (
	"net/http"

	"git.garena.com/sea-labs-id/batch-02/natanael-benediktus/final-project-backend/internal/handlers"
	"git.garena.com/sea-labs-id/batch-02/natanael-benediktus/final-project-backend/internal/helpers"
	"git.garena.com/sea-labs-id/batch-02/natanael-benediktus/final-project-backend/internal/middlewares"
	"github.com/gin-gonic/gin"
)

func InitAllRoutes(r *gin.Engine, h *handlers.Handler) {
	r.GET("", func(c *gin.Context) {
		c.JSON(http.StatusOK, helpers.BuildSuccessResponse(http.StatusOK, "House Rental StayCay App", nil))
	})

	// Auth
	r.POST("/register", h.RegisterHandler)
	r.POST("/login", h.LoginHandler)

	// City
	r.GET("/city", h.GetAllCitiesHandler)

	// House
	r.GET("/house", h.GetAllHousesHandler)
	r.GET("/house/:houseId", h.GetHouseByIDHandler)
	// r.GET("/house/unscoped/:houseId", h.GetUnscopedHouseByIDHandler)
	r.GET("/house/status", h.GetAllHouseStatusesHandler)
	r.GET("/house/reservation/:houseId", h.GetHouseReservationsByHouseHandler)

	// User
	r.GET("/user/:userId", h.GetUserProfileByUserIDHandler)

	// Pickup
	r.GET("/pickup/status", h.GetAllPickupStatusesHandler)
	r.GET("/pickup", h.GetAllPickupDetailsHandler)

	// Need Authorizations
	r.Use(middlewares.AuthorizeJWT)

	// User
	r.GET("/user", h.GetUserProfileHandler)
	r.PATCH("/user", h.UpdateUserProfileHandler)
	r.PATCH("/user/role", h.UpdateUserRoleHandler)
	r.GET("/user/reservation", h.GetHouseReservationsByUserHandler)

	// Topup
	r.POST("/topup", h.TopUpWalletHandler)

	// Host
	r.GET("/host/house", h.GetHostHousesHandler)
	r.GET("/host/reservation", h.GetHouseReservationsByHostHandler)
	r.POST("/host/house", h.AddNewHouseHandler)
	r.PATCH("/host/house/:houseId", h.UpdateHouseHandler)
	r.DELETE("/host/house/:houseId", h.DeleteHouseHandler)
	r.POST("/host/house/photo/:houseId", h.AddHousePhotoHandler)
	r.DELETE("/host/house/photo/:housePhotoId", h.DeleteHousePhotoHandler)

	// Reservation
	r.GET("/reservation/:reservationId", h.GetHouseReservationsByIDHandler)
	r.POST("/reservation", h.AddReservationHandler)

	// Pickup
	r.GET("/pickup/reservation/:reservationId", h.GetPickupDetailsByReservationHandler)
	r.POST("/pickup/reservation/:reservationId", h.RequestPickupHandler)

	// Game
	r.GET("/game", h.GetGameHandler)
	r.POST("/game", h.PlayGameHandler)

	// Admin
	r.PATCH("admin/pickup/:pickupId", h.AdminUpdatePickupStatusHandler)

}
