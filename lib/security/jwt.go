package security

import (
	"fmt"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/thoriqadillah/linktrim/lib/env"
)

var (
	key = env.Get("JWT_KEY").ToBytes()
	exp = env.Get("JWT_EXP").ToDuration()
)

func EncodeJWT(userID string) string {
	jwt := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp":        time.Now().Add(exp),
		"authorized": true,
		"user":       userID,
	})

	token, err := jwt.SignedString(key)
	if err != nil {
		log.Printf("Could not encode into token: %s", err.Error())
		return ""
	}

	return token
}

func DecodeJWT(token string) (userID uuid.UUID, err error) {
	jwtToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if t.Method != jwt.SigningMethodHS256 {
			return nil, fmt.Errorf("signing method invalid")
		}

		return key, nil
	})

	if err != nil {
		return uuid.Nil, fmt.Errorf("could not parse token: %s", err)
	}

	claims := jwtToken.Claims.(jwt.MapClaims)
	user := claims["user"].(string)
	if userID, err = uuid.Parse(user); err != nil {
		return uuid.Nil, err
	}

	return userID, nil
}
