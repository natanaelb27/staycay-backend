package middlewares

import (
	"net/http"

	"git.garena.com/sea-labs-id/batch-02/natanael-benediktus/final-project-backend/internal/helpers"
	"github.com/gin-gonic/gin"
)

func NotFoundHandler(c *gin.Context) {
	c.JSON(http.StatusNotFound, helpers.BuildErrorResponse(http.StatusNotFound, http.StatusText(http.StatusNotFound)))
}
