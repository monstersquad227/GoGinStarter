package demo

import (
	"GoGinStarter/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Get(c *gin.Context) {
	response := model.Res{
		Code:    0,
		Message: "successful",
		Data:    "ok",
	}
	c.JSON(http.StatusOK, response)
}

func Post(c *gin.Context) {
	response := model.Res{
		Code:    0,
		Message: "successful",
		Data:    "ok",
	}
	c.JSON(http.StatusOK, response)
}
