package handlers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/natanaelb27/staycay-backend/internal/dtos"
	e "github.com/natanaelb27/staycay-backend/internal/errors"
	"github.com/natanaelb27/staycay-backend/internal/helpers"
)

func (h *Handler) RegisterHandler(c *gin.Context) {
	var request *dtos.RegisterRequestDTO
	var response *dtos.RegisterResponseDTO

	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, helpers.BuildErrorResponse(http.StatusBadRequest, err.Error()))
		return
	}

	newUser, token, err := h.authService.Register(request)
	if err != nil {
		if errors.Is(err, e.ErrInvalidEmailInputError) {
			c.JSON(http.StatusBadRequest, helpers.BuildErrorResponse(http.StatusBadRequest, e.ErrInvalidEmailInputError.Error()))
			return
		}

		if errors.Is(err, e.ErrInvalidPasswordInputError) {
			c.JSON(http.StatusBadRequest, helpers.BuildErrorResponse(http.StatusBadRequest, e.ErrInvalidPasswordInputError.Error()))
			return
		}

		if errors.Is(err, e.ErrDuplicateEmailError) {
			c.JSON(http.StatusConflict, helpers.BuildErrorResponse(http.StatusConflict, e.ErrDuplicateEmailError.Error()))
			return
		}

		c.JSON(http.StatusInternalServerError, helpers.BuildErrorResponse(http.StatusInternalServerError, err.Error()))
		return
	}

	response = &dtos.RegisterResponseDTO{
		UserID: newUser.ID,
		Email:  newUser.Email,
		Token:  *token,
	}

	c.JSON(http.StatusCreated, helpers.BuildSuccessResponse(http.StatusCreated, http.StatusText(http.StatusCreated), response))
}

func (h *Handler) LoginHandler(c *gin.Context) {
	var request *dtos.LoginRequestDTO
	var response *dtos.LoginResponseDTO

	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, helpers.BuildErrorResponse(http.StatusBadRequest, err.Error()))
		return
	}

	user, token, err := h.authService.Login(request)
	if err != nil {
		if errors.Is(err, e.ErrCredentialError) {
			c.JSON(http.StatusUnauthorized, helpers.BuildErrorResponse(http.StatusUnauthorized, e.ErrCredentialError.Error()))
			return
		}

		c.JSON(http.StatusInternalServerError, helpers.BuildErrorResponse(http.StatusInternalServerError, err.Error()))
		return
	}

	response = &dtos.LoginResponseDTO{
		UserID: user.ID,
		Email:  user.Email,
		Role:   user.Role,
		Token:  *token,
	}

	c.JSON(http.StatusOK, helpers.BuildSuccessResponse(http.StatusOK, http.StatusText(http.StatusOK), response))
}
