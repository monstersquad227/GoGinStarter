package model

import "GoGinStarter/helper"

type Permission struct {
	Id     int    `json:"id"`
	Role   string `json:"role"`
	Remark string `json:"remark"`
}

type PermissionUser struct {
	Id           int    `json:"id"`
	PermissionId int    `json:"permission_id"`
	Name         string `json:"name"`
	Remark       string `json:"remark"`
}

func (p *Permission) Verify(nameStr string) (string, error) {
	var result string
	query := "SELECT role" +
		" FROM permission" +
		" WHERE id = (SELECT permission_id FROM permission_user WHERE name = ?)"
	mysqlEngin := helper.SqlContext
	err := mysqlEngin.QueryRow(query, nameStr).Scan(&result)
	if err != nil {
		return "", err
	}
	return result, nil
}
