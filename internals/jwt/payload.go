package jwt

import (
	"errors"
	"ezpz/config"
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"log"
	"strconv"
	"time"
)

type Payload struct {
	ID        uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	IssuedAt  time.Time `json:"issued_at"`
	ExpiredAt time.Time `json:"expired_at"`
}

func NewPayload(username string) *Payload {
	tokenID, _ := uuid.NewRandom()
	jwtExpireTime, err := strconv.Atoi(config.AppConfig()["jwt_expire"])
	if err != nil {
		log.Println(err)
	}
	duration := time.Now().Add(time.Minute * time.Duration(jwtExpireTime))
	payload := &Payload{
		ID:        tokenID,
		Username:  username,
		IssuedAt:  time.Now(),
		ExpiredAt: duration,
	}
	return payload
}

func keyFunc(token *jwt.Token) (interface{}, error) {
	_, ok := token.Method.(*jwt.SigningMethodHMAC)
	if !ok {
		return nil, errors.New("invalid token")
	}

	return []byte(config.AppConfig()["jwt_secret"]), nil
}
