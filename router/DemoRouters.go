package router

import (
	"GoGinStarter/controller/demo"
	"GoGinStarter/middleware"

	"github.com/gin-gonic/gin"
)

func BaseRegister(c *gin.Engine) {
	api := c.Group("/devops")
	{
		api.GET("/demo", middleware.JwtAuth(), demo.Get)
		api.POST("/demo", middleware.JwtAuth(), demo.Post)
	}
}
