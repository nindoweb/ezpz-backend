package jwt

import (
	"errors"
	"log"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"github.com/spf13/viper"
)

type Payload struct {
	ID        uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	IssuedAt  time.Time `json:"issued_at"`
	ExpiredAt time.Time `json:"expired_at"`
}

func NewPayload(username string) *Payload {
	tokenID, _ := uuid.NewRandom()
	jwtExpireTime, err := strconv.Atoi(viper.GetString("JWT_EXPIRE_TIME"))
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

	return []byte(viper.GetString("JWT_SECRET")), nil
}
