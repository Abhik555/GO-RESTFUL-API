package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const secretKey = "ultratopsecrect"

func GenerateToken(email string, id int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userID": id,
		"expiry": time.Now().Add(time.Hour * 2).Unix(),
	})

	return token.SignedString([]byte(secretKey))
}

func VerifyToken(token string) (int64,error) {
	ptoken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("Invalid Signature.")
		}
		return []byte(secretKey), nil
	})

	if err != nil {
		return 0,errors.New("Could not parse token.")
	}

	if !ptoken.Valid {
		return 0,errors.New("Invalid token.")
	}

	claims , ok := ptoken.Claims.(jwt.MapClaims)

	if !ok {
		return 0,errors.New("invalid claims.")
	}

	// email := claims["email"].(string)
	userid := int64(claims["userID"].(float64))

	return userid ,nil
}
