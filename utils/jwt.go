package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)


const secretKey string = "$%MBxwbSd85d$3q47d&jQ@bm$Mx5X$3F&y"


func GenerateToken(username string, userId int64) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"userId": userId,
		"exp": time.Now().Add(time.Minute * 20).Unix(),
	})

	return token.SignedString([]byte(secretKey))
}

func VerifyToken(token string) (int64, error) {
	parsedToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("error with method jwt token")
		}

		return []byte(secretKey), nil
	})

	if err != nil {
		return 0, errors.New("erros with jwt token")
	}

	 if validToken := parsedToken.Valid; !validToken {
		return 0, errors.New("invalid jwt token")
	}

	claim := parsedToken.Claims.(jwt.MapClaims)

	userId := int64(claim["userId"].(float64))

	return userId, nil
}