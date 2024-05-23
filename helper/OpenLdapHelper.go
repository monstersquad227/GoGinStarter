package helper

import (
	"GoGinStarter/configuration"
	"fmt"
	"github.com/go-ldap/ldap/v3"
	"log"
)

/*
	LDAP连接函数
*/

func connectLDAP() (*ldap.Conn, error) {
	l, err := ldap.Dial("tcp", configuration.Configs.LdapHost+":"+configuration.Configs.LdapPort)
	if err != nil {
		log.Printf("Error connecting to LDAP server: %v", err)
		return nil, err
	}
	return l, nil
}

/*
	验证LDAP用户
*/

func OpenldapVerify(username, password string) bool {
	l, err := connectLDAP()
	if err != nil {
		return false
	}
	defer l.Close()
	err = l.Bind(fmt.Sprintf("cn=%s,ou=技术部,dc=mojorycorp,dc=cn", username), password)
	if err != nil {
		log.Printf("Error binding with user's credentials: %v", err)
		return false
	}
	return true
}

/*
	修改LDAP用户密码
*/

func OpenldapModifyPassword(username, newPassword string) error {
	l, err := connectLDAP()
	if err != nil {
		return err
	}
	defer l.Close()

	err = l.Bind("cn=admin,dc=mojorycorp,dc=cn", "mojory@1q2w3e4r")
	if err != nil {
		log.Printf("Error binding with admin credentials: %v", err)
		return err
	}

	modifyRequest := ldap.NewModifyRequest(fmt.Sprintf("cn=%s,ou=技术部,dc=mojorycorp,dc=cn", username), []ldap.Control{})
	modifyRequest.Replace("userPassword", []string{newPassword})

	// 执行密码修改请求
	err = l.Modify(modifyRequest)
	if err != nil {
		log.Printf("Error modifying user password: %v", err)
		return err
	}
	return nil
}
