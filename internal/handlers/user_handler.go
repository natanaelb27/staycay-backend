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

func (h *Handler) GetUserProfileHandler(c *gin.Context) {
	var response *dtos.GetUserProfileResponseDTO
	userContext := c.MustGet("user").(dtos.JwtData)

	userDetails, err := h.userService.GetUserProfile(userContext.UserID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, helpers.BuildErrorResponse(http.StatusInternalServerError, err.Error()))
		return
	}

	response = &dtos.GetUserProfileResponseDTO{
		UserID:   userDetails.UserID,
		Name:     userDetails.Name,
		Email:    userDetails.Email,
		Address:  userDetails.Address,
		CityID:   userDetails.CityID,
		CityName: userDetails.CityName,
		Role:     userDetails.Role,
		WalletID: userDetails.WalletID,
		Balance:  userDetails.Balance,
	}

	c.JSON(http.StatusOK, helpers.BuildSuccessResponse(http.StatusOK, http.StatusText(http.StatusOK), response))
}

func (h *Handler) GetUserProfileByUserIDHandler(c *gin.Context) {
	var response *dtos.GetUserProfileByUserIDResponseDTO
	userID, _ := strconv.Atoi(c.Param("userId"))

	userDetails, err := h.userService.GetUserProfileByUserID(userID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, helpers.BuildErrorResponse(http.StatusInternalServerError, err.Error()))
		return
	}

	response = &dtos.GetUserProfileByUserIDResponseDTO{
		UserID:   userDetails.UserID,
		Name:     userDetails.Name,
		Email:    userDetails.Email,
		Address:  userDetails.Address,
		CityID:   userDetails.CityID,
		CityName: userDetails.CityName,
		Role:     userDetails.Role,
		WalletID: userDetails.WalletID,
		Balance:  userDetails.Balance,
	}

	c.JSON(http.StatusOK, helpers.BuildSuccessResponse(http.StatusOK, http.StatusText(http.StatusOK), response))
}

func (h *Handler) UpdateUserRoleHandler(c *gin.Context) {
	var response *dtos.UpdateUserRoleResponseDTO
	userContext := c.MustGet("user").(dtos.JwtData)
	userData, err := h.userService.UpdateUserRoleToHost(userContext.UserID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, helpers.BuildErrorResponse(http.StatusInternalServerError, err.Error()))
		return
	}

	response = &dtos.UpdateUserRoleResponseDTO{
		UserID: userData.ID,
		Role:   userData.Role,
	}

	c.JSON(http.StatusOK, helpers.BuildSuccessResponse(http.StatusOK, http.StatusText(http.StatusOK), response))
}

func (h *Handler) UpdateUserProfileHandler(c *gin.Context) {
	var request *dtos.UpdateUserProfileRequestDTO
	var response *dtos.UpdateUserProfileResponseDTO
	userContext := c.MustGet("user").(dtos.JwtData)

	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, helpers.BuildErrorResponse(http.StatusBadRequest, err.Error()))
		return
	}

	userData, err := h.userService.UpdateUserProfile(userContext.UserID, request)

	if err != nil {
		if errors.Is(err, e.ErrInvalidPasswordInputError) {
			c.JSON(http.StatusBadRequest, helpers.BuildErrorResponse(http.StatusBadRequest, e.ErrInvalidPasswordInputError.Error()))
			return
		}
		
		if errors.Is(err, e.ErrOldPasswordNotMatchError) {
			c.JSON(http.StatusBadRequest, helpers.BuildErrorResponse(http.StatusBadRequest, e.ErrOldPasswordNotMatchError.Error()))
			return
		}

		c.JSON(http.StatusInternalServerError, helpers.BuildErrorResponse(http.StatusInternalServerError, err.Error()))
		return
	}

	response = &dtos.UpdateUserProfileResponseDTO{
		UserID:  userData.ID,
		Name:    userData.Name,
		Address: userData.Address,
		CityID:  userData.CityID,
	}

	c.JSON(http.StatusOK, helpers.BuildSuccessResponse(http.StatusOK, http.StatusText(http.StatusOK), response))
}
