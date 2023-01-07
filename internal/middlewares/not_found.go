package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/natanaelb27/staycay-backend/internal/helpers"
)

func NotFoundHandler(c *gin.Context) {
	c.JSON(http.StatusNotFound, helpers.BuildErrorResponse(http.StatusNotFound, http.StatusText(http.StatusNotFound)))
}
