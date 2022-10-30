package service

import (
	"myblog/conf"
	"myblog/errmsg"
)

func CheckLogin(username string, password string) int {
	if username != conf.Config.Username {
		return errmsg.ERROR_USER_NOT_EXIST
	}
	if password != conf.Config.Password {
		return errmsg.ERROR_PASSWORD_WRONG
	}
	return errmsg.SUCCESS
}
