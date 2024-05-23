package router

import (
	"GoGinStarter/controller/login"
	"GoGinStarter/middleware"

	"github.com/gin-gonic/gin"
)

func LoginRegister(c *gin.Engine) {
	api := c.Group("/devops")
	{
		api.POST("/user/login", login.Auth)
		api.POST("/user/password", middleware.JwtAuth(), login.Password)
	}
}
