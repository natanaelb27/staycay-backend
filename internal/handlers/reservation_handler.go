package handlers

import (
	"errors"
	"net/http"

	"git.garena.com/sea-labs-id/batch-02/natanael-benediktus/final-project-backend/internal/dtos"
	e "git.garena.com/sea-labs-id/batch-02/natanael-benediktus/final-project-backend/internal/errors"
	"git.garena.com/sea-labs-id/batch-02/natanael-benediktus/final-project-backend/internal/helpers"
	"github.com/gin-gonic/gin"
)

func (h *Handler) AddReservationHandler(c *gin.Context) {
	var request *dtos.AddReservationRequestDTO
	var response *dtos.AddReservationResponseDTO

	userContext := c.MustGet("user").(dtos.JwtData)

	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, helpers.BuildErrorResponse(http.StatusBadRequest, err.Error()))
		return
	}

	data, err := h.reservationService.AddReservation(userContext.UserID, request)

	if err != nil {
		if errors.Is(err, e.ErrNotValidReservationDateConflict) {
			c.JSON(http.StatusBadRequest, helpers.BuildErrorResponse(http.StatusBadRequest, e.ErrNotValidReservationDateConflict.Error()))
			return
		}

		c.JSON(http.StatusInternalServerError, helpers.BuildErrorResponse(http.StatusInternalServerError, err.Error()))
		return
	}

	response = &dtos.AddReservationResponseDTO{
		ReservationID:   data.ID,
		HouseID:         data.HouseID,
		UserID:          data.UserID,
		ReservationCode: data.ReservationCode,
		CheckInDate:     data.CheckInDate,
		CheckOutDate:    data.CheckOutDate,
		TotalPrice:      data.TotalPrice,
	}

	c.JSON(http.StatusOK, helpers.BuildSuccessResponse(http.StatusOK, http.StatusText(http.StatusOK), response))
}
