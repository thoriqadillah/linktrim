package security

import (
	"encoding/base64"
	"fmt"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

var defaultkey = "default"

func EncodeJWT(userID string) string {
	jwt := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp":        time.Duration(24 * time.Hour),
		"authorized": true,
		"user":       userID,
	})

	//TODO: generate public key
	_, err := base64.URLEncoding.DecodeString("oRIU9Idp91hsO_taulZ8LRbxvwAYvoiteoj7prWj944=")
	if err != nil {
		log.Panicf("Could not decode key: %s", err.Error())
		return ""
	}

	token, err := jwt.SignedString([]byte(defaultkey))
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

		return []byte(defaultkey), nil
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
