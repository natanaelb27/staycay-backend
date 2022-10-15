package routes

import (
	"git.garena.com/sea-labs-id/batch-02/natanael-benediktus/final-project-backend/internal/handlers"
	"git.garena.com/sea-labs-id/batch-02/natanael-benediktus/final-project-backend/internal/middlewares"
	"github.com/gin-gonic/gin"
)

func InitAllRoutes(r *gin.Engine, h *handlers.Handler) {
	// Auth
	r.POST("/register", h.RegisterHandler)
	r.POST("/login", h.LoginHandler)

	// City
	r.GET("/city", h.GetAllCitiesHandler)

	// House
	r.GET("/house", h.GetAllHousesHandler)

	r.Use(middlewares.AuthorizeJWT)

	// User
	r.GET("/user", h.GetUserProfileHandler)
	r.PATCH("/user/", h.UpdateUserProfileHandler)
	r.PATCH("/user/role", h.UpdateUserRoleHandler)

	// Topup
	r.POST("/topup", h.TopUpWalletHandler)

	// Host
	r.GET("/host/house", h.GetHostHousesHandler)
	r.POST("/host/house", h.AddNewHouseHandler)
	r.PATCH("/host/house/:houseId", h.UpdateHouseHandler)
	r.DELETE("/host/house/photo/:housePhotoId", h.DeleteHousePhotoHandler)
}
