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
		if errors.Is(err, e.ErrIdNotFoundError) {
			c.JSON(http.StatusBadRequest, helpers.BuildErrorResponse(http.StatusBadRequest, e.ErrIdNotFoundError.Error()))
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
		CityID:         data.CityID,
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
		if errors.Is(err, e.ErrIdNotFoundError) {
			c.JSON(http.StatusBadRequest, helpers.BuildErrorResponse(http.StatusBadRequest, e.ErrIdNotFoundError.Error()))
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

	reservationID, _ := strconv.Atoi(c.Param("reservationId"))

	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, helpers.BuildErrorResponse(http.StatusBadRequest, err.Error()))
		return
	}

	pickup, err := h.pickupService.AdminUpdatePickupStatus(request.PickupStatusID, reservationID, userContext.UserID)

	if err != nil {
		if errors.Is(err, e.ErrInvalidRoleNotAdminError) {
			c.JSON(http.StatusForbidden, helpers.BuildErrorResponse(http.StatusForbidden, e.ErrInvalidRoleNotAdminError.Error()))
			return
		}

		if errors.Is(err, e.ErrIdNotFoundError) {
			c.JSON(http.StatusBadRequest, helpers.BuildErrorResponse(http.StatusBadRequest, e.ErrIdNotFoundError.Error()))
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

// func (h *Handler) GetAllPickupStatusesHandler(c *gin.Context) {
// 	var response *dtos.GetAllPickupStatusesResponseDTO

// 	data, err := h.pickupService.GetAllPickupStatuses()
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, helpers.BuildErrorResponse(http.StatusInternalServerError, err.Error()))
// 		return
// 	}

// 	response = &dtos.GetAllPickupStatusesResponseDTO{
// 		PickupStatus: data,
// 	}

// 	c.JSON(http.StatusOK, helpers.BuildSuccessResponse(http.StatusOK, http.StatusText(http.StatusOK), response))
// }
