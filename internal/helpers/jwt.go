package helpers

import (
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/joho/godotenv"
)

type IdTokenClaims struct {
	jwt.RegisteredClaims
	Data interface{} `json:"data"`
}

func GenerateJwtToken(data interface{}) (*string, error) {
	godotenv.Load(".env")
	idExp, _ := strconv.ParseInt(os.Getenv("EXPIRED_TIME"), 10, 64)
	unixTime := time.Now().Unix()
	tokenExp := unixTime + idExp

	claims := IdTokenClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer: os.Getenv("ISSUER"),
			ExpiresAt: &jwt.NumericDate{
				Time: time.Unix(tokenExp, 0),
			},
			IssuedAt: &jwt.NumericDate{
				Time: time.Now(),
			},
		},
		Data: data,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRETKEY")))

	if err != nil {
		return nil, err
	}

	return &tokenString, nil
}
