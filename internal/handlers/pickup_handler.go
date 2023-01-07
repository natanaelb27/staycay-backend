package handlers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/natanaelb27/staycay-backend/internal/dtos"
	e "github.com/natanaelb27/staycay-backend/internal/errors"
	"github.com/natanaelb27/staycay-backend/internal/helpers"
)

func (h *Handler) RequestPickupHandler(c *gin.Context) {
	var request *dtos.RequestPickupRequestDTO
	var response *dtos.RequestPickupResponseDTO

	userContext := c.MustGet("user").(dtos.JwtData)

	reservationID, _ := strconv.Atoi(c.Param("reservationId"))

	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, helpers.BuildErrorResponse(http.StatusBadRequest, err.Error()))
		return
	}

	data, err := h.pickupService.RequestPickup(reservationID, userContext.UserID, request)

	if err != nil {
		if errors.Is(err, e.ErrReservationNotFoundError) {
			c.JSON(http.StatusNotFound, helpers.BuildErrorResponse(http.StatusNotFound, e.ErrReservationNotFoundError.Error()))
			return
		}

		if errors.Is(err, e.ErrUserNotFoundError) {
			c.JSON(http.StatusNotFound, helpers.BuildErrorResponse(http.StatusNotFound, e.ErrUserNotFoundError.Error()))
			return
		}

		if errors.Is(err, e.ErrDuplicatePickupReservationError) {
			c.JSON(http.StatusBadRequest, helpers.BuildErrorResponse(http.StatusBadRequest, e.ErrDuplicatePickupReservationError.Error()))
			return
		}

		if errors.Is(err, e.ErrInvalidUserReservationError) {
			c.JSON(http.StatusForbidden, helpers.BuildErrorResponse(http.StatusForbidden, e.ErrInvalidUserReservationError.Error()))
			return
		}

		if errors.Is(err, e.ErrPickupCompletedReservation) {
			c.JSON(http.StatusBadRequest, helpers.BuildErrorResponse(http.StatusBadRequest, e.ErrPickupCompletedReservation.Error()))
			return
		}

		if errors.Is(err, e.ErrPickupOngoingReservation) {
			c.JSON(http.StatusBadRequest, helpers.BuildErrorResponse(http.StatusBadRequest, e.ErrPickupOngoingReservation.Error()))
			return
		}

		if errors.Is(err, e.ErrInsufficientBalanceError) {
			c.JSON(http.StatusBadRequest, helpers.BuildErrorResponse(http.StatusBadRequest, e.ErrInsufficientBalanceError.Error()))
			return
		}

		c.JSON(http.StatusInternalServerError, helpers.BuildErrorResponse(http.StatusInternalServerError, err.Error()))
		return
	}

	response = &dtos.RequestPickupResponseDTO{
		PickupID:       data.ID,
		ReservationID:  data.ReservationID,
		PickupStatusID: data.PickupStatusID,
		PickupAddress:  data.PickupAddress,
		PickupCityID:   data.CityID,
		PickupPrice:    data.PickupPrice,
	}

	c.JSON(http.StatusOK, helpers.BuildSuccessResponse(http.StatusOK, http.StatusText(http.StatusOK), response))
}

func (h *Handler) GetPickupDetailsByReservationHandler(c *gin.Context) {
	var response *dtos.GetPickupDetailsByReservationResponseDTO

	userContext := c.MustGet("user").(dtos.JwtData)

	reservationID, _ := strconv.Atoi(c.Param("reservationId"))

	response, err := h.pickupService.GetPickupDetailsByReservation(reservationID, userContext.UserID)

	if err != nil {
		if errors.Is(err, e.ErrReservationNotFoundError) {
			c.JSON(http.StatusNotFound, helpers.BuildErrorResponse(http.StatusNotFound, e.ErrReservationNotFoundError.Error()))
			return
		}

		c.JSON(http.StatusInternalServerError, helpers.BuildErrorResponse(http.StatusInternalServerError, err.Error()))
		return
	}

	c.JSON(http.StatusOK, helpers.BuildSuccessResponse(http.StatusOK, http.StatusText(http.StatusOK), response))
}

func (h *Handler) AdminUpdatePickupStatusHandler(c *gin.Context) {
	var request *dtos.AdminUpdatePickupStatusRequestDTO
	var response *dtos.AdminUpdatePickupStatusResponseDTO

	userContext := c.MustGet("user").(dtos.JwtData)

	pickupID, _ := strconv.Atoi(c.Param("pickupId"))

	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, helpers.BuildErrorResponse(http.StatusBadRequest, err.Error()))
		return
	}

	pickup, err := h.pickupService.AdminUpdatePickupStatus(request.PickupStatusID, pickupID, userContext.UserID)

	if err != nil {
		if errors.Is(err, e.ErrUserNotFoundError) {
			c.JSON(http.StatusNotFound, helpers.BuildErrorResponse(http.StatusNotFound, e.ErrUserNotFoundError.Error()))
			return
		}

		if errors.Is(err, e.ErrReservationNotFoundError) {
			c.JSON(http.StatusNotFound, helpers.BuildErrorResponse(http.StatusNotFound, e.ErrReservationNotFoundError.Error()))
			return
		}

		if errors.Is(err, e.ErrInvalidCheckInDateUpdatePickupStatusError) {
			c.JSON(http.StatusBadRequest, helpers.BuildErrorResponse(http.StatusBadRequest, e.ErrInvalidCheckInDateUpdatePickupStatusError.Error()))
			return
		}

		if errors.Is(err, e.ErrInvalidRoleNotAdminError) {
			c.JSON(http.StatusForbidden, helpers.BuildErrorResponse(http.StatusForbidden, e.ErrInvalidRoleNotAdminError.Error()))
			return
		}

		if errors.Is(err, e.ErrPickupNotFoundError) {
			c.JSON(http.StatusNotFound, helpers.BuildErrorResponse(http.StatusNotFound, e.ErrPickupNotFoundError.Error()))
			return
		}

		c.JSON(http.StatusInternalServerError, helpers.BuildErrorResponse(http.StatusInternalServerError, err.Error()))
		return
	}

	response = &dtos.AdminUpdatePickupStatusResponseDTO{
		ReservationID:  pickup.ReservationID,
		PickupStatusID: pickup.PickupStatusID,
	}

	c.JSON(http.StatusOK, helpers.BuildSuccessResponse(http.StatusOK, http.StatusText(http.StatusOK), response))
}

func (h *Handler) GetAllPickupStatusesHandler(c *gin.Context) {
	var response *dtos.GetAllPickupStatusesResponseDTO

	data, err := h.pickupService.GetAllPickupStatuses()
	if err != nil {
		c.JSON(http.StatusInternalServerError, helpers.BuildErrorResponse(http.StatusInternalServerError, err.Error()))
		return
	}

	response = &dtos.GetAllPickupStatusesResponseDTO{
		PickupStatus: data,
	}

	c.JSON(http.StatusOK, helpers.BuildSuccessResponse(http.StatusOK, http.StatusText(http.StatusOK), response))
}

func (h *Handler) GetAllPickupDetailsHandler(c *gin.Context) {
	var response *dtos.GetAllPickupDetailsResponseDTO

	queryParams := map[string]interface{}{}
	for key, values := range c.Request.URL.Query() {
		queryParams[key] = values[0]
	}

	data, pagination, err := h.pickupService.GetAllPickupDetails(queryParams)
	if err != nil {
		c.JSON(http.StatusInternalServerError, helpers.BuildErrorResponse(http.StatusInternalServerError, err.Error()))
		return
	}

	response = &dtos.GetAllPickupDetailsResponseDTO{
		Pagination: pagination,
		Pickup:     data,
	}

	c.JSON(http.StatusOK, helpers.BuildSuccessResponse(http.StatusOK, http.StatusText(http.StatusOK), response))
}
