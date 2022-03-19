package jwt

import (
	"fmt"

	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
)

func CreateToken(payload *Payload) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["id"] = payload.ID
	claims["issued_at"] = payload.IssuedAt
	claims["username"] = payload.Username
	claims["expired"] = payload.ExpiredAt
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	jwtSecret := fmt.Sprintf("value: %v", viper.GetString("JWT_SECRET"))

	return token.SignedString([]byte(jwtSecret))
}
