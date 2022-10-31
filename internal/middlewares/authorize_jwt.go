package middlewares

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"

	"git.garena.com/sea-labs-id/batch-02/natanael-benediktus/final-project-backend/internal/dtos"
	"git.garena.com/sea-labs-id/batch-02/natanael-benediktus/final-project-backend/internal/helpers"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func validateToken(encodedToken string) (*jwt.Token, error) {
	// godotenv.Load(".env")
	return jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		hmacSampleSecret := []byte(os.Getenv("SECRETKEY"))
		return hmacSampleSecret, nil
	})
}

func AuthorizeJWT(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")

	if authHeader == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, helpers.BuildErrorResponse(http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized)))
		return
	}

	tokenString := strings.Split(authHeader, " ")
	if len(tokenString) != 2 {
		c.AbortWithStatusJSON(http.StatusUnauthorized, helpers.BuildErrorResponse(http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized)))
		return
	}

	token, err := validateToken(tokenString[1])

	if err != nil || !token.Valid {
		c.AbortWithStatusJSON(http.StatusUnauthorized, helpers.BuildErrorResponse(http.StatusUnauthorized, err.Error()))
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok {
		c.AbortWithStatusJSON(http.StatusUnauthorized, helpers.BuildErrorResponse(http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized)))
		return
	}

	userJson, _ := json.Marshal(claims["data"])
	user := dtos.JwtData{}
	err = json.Unmarshal(userJson, &user)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, helpers.BuildErrorResponse(http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized)))
		return
	}

	c.Set("user", user)

}
