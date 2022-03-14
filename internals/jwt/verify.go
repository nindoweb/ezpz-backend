package jwt

import (
	"github.com/dgrijalva/jwt-go"
)

func VerifyToken(t string) (*jwt.Token, error) {
	return jwt.ParseWithClaims(t, jwt.MapClaims{}, keyFunc)
}
