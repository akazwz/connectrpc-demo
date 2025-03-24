package utils

import (
	"github.com/golang-jwt/jwt/v5"
	"os"
	"time"
)

var secret string

func init() {
	secret = os.Getenv("JWT_SECRET")
	if secret == "" {
		panic("JWT_SECRET is not set")
	}
}

type CustomClaims struct {
	UserID string `json:"user_id"`
	jwt.RegisteredClaims
}

func SignToken(userID string, expiresAt time.Time) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, CustomClaims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "opendrive",
			ExpiresAt: jwt.NewNumericDate(expiresAt),
		},
	})
	signedToken, _ := token.SignedString([]byte(secret))
	return signedToken
}

func VerifyToken(token string) (*CustomClaims, error) {
	var claims CustomClaims
	_, err := jwt.ParseWithClaims(token, &claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	return &claims, err
}
