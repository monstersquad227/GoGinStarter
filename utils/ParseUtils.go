package utils

import (
	"GoGinStarter/middleware"
	"errors"
	"github.com/dgrijalva/jwt-go"
)

func ParseJwtUsername(tokenStr string) (string, error) {
	parse, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return middleware.JwtKey, nil
	})
	if err != nil {
		return "", err
	}
	claims, ok := parse.Claims.(jwt.MapClaims)
	if !ok || !parse.Valid {
		return "", errors.New("invalid token")
	}
	username, ok := claims["username"].(string)
	if !ok {
		return "", errors.New("username field not found or not a string")
	}
	return username, nil
}
