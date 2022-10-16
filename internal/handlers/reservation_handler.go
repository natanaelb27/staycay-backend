package handlers

import (
	"errors"
	"net/http"
	"strconv"

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
		if errors.Is(err, e.ErrInvalidReservationDateLaterError) {
			c.JSON(http.StatusBadRequest, helpers.BuildErrorResponse(http.StatusBadRequest, e.ErrInvalidReservationDateLaterError.Error()))
			return
		}

		if errors.Is(err, e.ErrInvalidReservationDateSameError) {
			c.JSON(http.StatusBadRequest, helpers.BuildErrorResponse(http.StatusBadRequest, e.ErrInvalidReservationDateSameError.Error()))
			return
		}

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

func (h *Handler) GetHouseReservationsByUserHandler(c *gin.Context) {
	var response *dtos.GetHouseReservationsByUserResponseDTO

	userContext := c.MustGet("user").(dtos.JwtData)

	queryParams := map[string]interface{}{}
	for key, values := range c.Request.URL.Query() {
		queryParams[key] = values[0]
	}

	data, err := h.reservationService.GetHouseReservationsByUser(userContext.UserID, queryParams)

	if err != nil {
		c.JSON(http.StatusInternalServerError, helpers.BuildErrorResponse(http.StatusInternalServerError, err.Error()))
		return
	}

	response = &dtos.GetHouseReservationsByUserResponseDTO{
		Reservation: data,
	}

	c.JSON(http.StatusOK, helpers.BuildSuccessResponse(http.StatusOK, http.StatusText(http.StatusOK), response))
}

func (h *Handler) GetHouseReservationsByHouseHandler(c *gin.Context) {
	var response *dtos.GetHouseReservationsByHouseResponseDTO

	houseID, _ := strconv.Atoi(c.Param("houseId"))

	userContext := c.MustGet("user").(dtos.JwtData)

	queryParams := map[string]interface{}{}
	for key, values := range c.Request.URL.Query() {
		queryParams[key] = values[0]
	}

	data, err := h.reservationService.GetHouseReservationsByHouse(houseID, userContext.UserID, queryParams)

	if err != nil {
		if errors.Is(err, e.ErrInvalidHostError) {
			c.JSON(http.StatusForbidden, helpers.BuildErrorResponse(http.StatusForbidden, e.ErrInvalidHostError.Error()))
			return
		}

		c.JSON(http.StatusInternalServerError, helpers.BuildErrorResponse(http.StatusInternalServerError, err.Error()))
		return
	}

	response = &dtos.GetHouseReservationsByHouseResponseDTO{
		Reservation: data,
	}

	c.JSON(http.StatusOK, helpers.BuildSuccessResponse(http.StatusOK, http.StatusText(http.StatusOK), response))

}
