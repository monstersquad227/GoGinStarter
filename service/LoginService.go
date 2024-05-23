package service

import (
	"GoGinStarter/configuration"
	"GoGinStarter/helper"
	"GoGinStarter/middleware"
	"GoGinStarter/model"
	"GoGinStarter/utils"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func LoginAuth(username, password string) model.Res {
	response := model.Res{
		Code:    0,
		Message: "Successful",
		Data:    nil,
	}

	// 生成加密密码
	EnPassword, err := utils.Encrypt(password, configuration.Configs.EncryptionKey)
	if err != nil {
		response.Message = err.Error()
		return response
	}

	// 验证用户是否正确
	u := model.User{
		Username: username,
		Password: EnPassword,
	}
	row, err := u.CheckUsername()
	if err != nil {
		response.Code = 40001
		response.Message = err.Error()
		return response
	}
	if row != 1 {
		response.Code = 40001
		response.Message = "用户不存在"
		return response
	}

	// 验证密码是否正确
	passwordStr, err := u.CheckPasswordByUsername()
	if err != nil {
		response.Code = 40001
		response.Message = err.Error()
		return response
	}
	if passwordStr != u.Password {
		response.Code = 40001
		response.Message = "密码错误"
		return response
	}
	// 生成Jwt token
	expirationTime := time.Now().Add(time.Duration(configuration.Configs.ExpireTime) * time.Hour)
	claims := &model.JwtClaims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(middleware.JwtKey)
	if err != nil {
		response.Message = err.Error()
		return response
	}

	// 返回数据
	type responseBody struct {
		Token string `json:"token"`
	}
	response.Data = responseBody{tokenString}
	return response
}

func LoginPassword(username, password, newPassword string) model.Res {
	response := model.Res{
		Code:    0,
		Message: "Successful",
		Data:    nil,
	}

	// 加密
	EnPassword, err := utils.Encrypt(password, configuration.Configs.EncryptionKey)
	if err != nil {
		response.Message = err.Error()
		response.Code = 50000
		return response
	}
	EnNewPassword, err := utils.Encrypt(newPassword, configuration.Configs.EncryptionKey)
	if err != nil {
		response.Message = err.Error()
		response.Code = 50000
		return response
	}

	u := model.User{
		Username: username,
		Password: EnPassword,
	}
	passwordStr, err := u.CheckPasswordByUsername()
	if err != nil {
		response.Code = 50000
		response.Message = err.Error()
		return response
	}
	if passwordStr != EnPassword {
		response.Code = 50000
		response.Message = "当前密码不正确"
		return response
	}

	result, err := u.ModifyPassword(EnNewPassword)
	if err != nil {
		response.Code = 50000
		response.Message = err.Error()
		return response
	}

	response.Data = result
	return response
}

func LoginAuthLdap(username, password string) model.Res {
	response := model.Res{
		Code:    0,
		Message: "Login Successful",
		Data:    nil,
	}
	if ok := helper.OpenldapVerify(username, password); ok == false {
		response.Code = 50000
		response.Message = "用户名或密码错误"
		return response
	}

	expirationTime := time.Now().Add(time.Duration(configuration.Configs.ExpireTime) * time.Hour)
	claims := &model.JwtClaims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(middleware.JwtKey)
	if err != nil {
		response.Code = 50000
		response.Message = err.Error()
		return response
	}
	type responseBody struct {
		Token string `json:"token"`
	}
	response.Data = responseBody{tokenString}
	return response
}

func LoginPasswordLdap(username, password, newPassword string) model.Res {
	response := model.Res{
		Code:    0,
		Message: "Successful",
		Data:    nil,
	}
	if err := helper.OpenldapVerify(username, password); err == false {
		response.Code = 50000
		response.Message = "当前密码不正确"
		return response
	}
	if err := helper.OpenldapModifyPassword(username, newPassword); err != nil {
		response.Code = 50000
		response.Message = err.Error()
		return response
	}

	return response
}
