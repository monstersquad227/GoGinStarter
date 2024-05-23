package middleware

import (
	"GoGinStarter/configuration"
	"GoGinStarter/model"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
)

var JwtKey = []byte(configuration.Configs.JwtSecretKey)

func JwtAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		response := model.Res{
			Code:    http.StatusOK,
			Message: "successful",
			Data:    nil,
		}

		// 验证请求头
		tokenString := c.Request.Header.Get("Authorization")
		if tokenString == "" {
			response.Code = http.StatusUnauthorized
			response.Message = "Authorization header required"
			c.JSON(http.StatusUnauthorized, response)
			c.Abort()
			return
		}

		token, err := jwt.ParseWithClaims(tokenString, &model.JwtClaims{}, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrSignatureInvalid
			}
			return JwtKey, nil
		})

		if err != nil {
			response.Code = http.StatusUnauthorized
			response.Message = "Invalid token"
			c.JSON(http.StatusUnauthorized, response)
			c.Abort()
			return
		}

		if claims, ok := token.Claims.(*model.JwtClaims); ok && token.Valid {
			c.Set("username", claims.Username)
			c.Next()
		} else {
			response.Code = http.StatusUnauthorized
			response.Message = "Invalid token"
			c.JSON(http.StatusUnauthorized, response)
			c.Abort()
			return
		}
	}
}
