package mysql

import "errors"

var (
	ErrorInvalidID       = errors.New("无效的ID")
	ErrorQueryFailed     = errors.New("查询数据失败")
	ErrorUserExist       = errors.New("用户已存在")
	ErrorUserNotExist    = errors.New("用户不存在")
	ErrorInvalidPassword = errors.New("用户名或密码错误")
)
