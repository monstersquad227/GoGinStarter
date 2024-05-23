package model

import (
	"GoGinStarter/helper"
	"database/sql"
)

type User struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func (u *User) CheckUsername() (int, error) {
	var total int
	query := "SELECT count(*) " +
		"FROM user " +
		"WHERE username = ?"
	mysqlEngine := helper.SqlContext
	err := mysqlEngine.QueryRow(query, u.Username).Scan(&total)
	if err != nil {
		return 0, err
	}
	return total, err
}

func (u *User) CheckPasswordByUsername() (string, error) {
	var password string
	query := "SELECT password " +
		"FROM user " +
		"WHERE username = ?"
	mysqlEngine := helper.SqlContext
	err := mysqlEngine.QueryRow(query, u.Username).Scan(&password)
	if err != nil {
		return "", err
	}
	return password, nil
}

func (u *User) ModifyPassword(newPassword string) (sql.Result, error) {
	query := "UPDATE user " +
		"SET password = ? WHERE username = ?"
	mysqlEngine := helper.SqlContext
	data, err := mysqlEngine.Exec(query, newPassword, u.Username)
	if err != nil {
		return nil, err
	}
	return data, nil
}
