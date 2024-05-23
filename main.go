package main

import (
	"GoGinStarter/configuration"
	"GoGinStarter/helper"
	"GoGinStarter/middleware"
	"GoGinStarter/router"
	"fmt"

	"github.com/gin-gonic/gin"
)

func init() {
	configuration.LoadConfig("./config.ini")
	helper.MysqlConnect()
	helper.GitlabConnect()
	helper.JenkinsConnect()
}

func main() {

	app := gin.Default()

	// 跨域中间件
	app.Use(middleware.Cors())

	// 路由
	router.BaseRegister(app)
	router.LoginRegister(app)

	err := app.Run("0.0.0.0:" + configuration.Configs.ServerPort)
	if err != nil {
		fmt.Printf("Error starting server: %v\n", err)
		return
	}
}
