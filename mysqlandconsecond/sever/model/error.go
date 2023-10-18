package model

import "errors"

var (
	ERR_USER_NOTEX   = errors.New("用户不存在,请先进行注册")
	ERR_USER_EX      = errors.New("用户已经存在")
	ERR_USER_PWD     = errors.New("密码错误")
	ERR_USERNAME     = errors.New("用户名错误")
	ERR_USER_ONLINED = errors.New("该账户正在使用,请切换账户")
)
