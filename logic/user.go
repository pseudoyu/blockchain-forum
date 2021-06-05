package logic

import (
	"blockchainguide_app/dao/mysql"
	"blockchainguide_app/models"
	"blockchainguide_app/pkg/snowflake"
)

// 存放业务逻辑的代码

func SignUp(p *models.ParamSignUp) {
	// 1. 判断用户存不存在
	mysql.QueryUserByUsername()
	// 2. 生成UID
	snowflake.GenID()
	// 3. 保存进数据库
	mysql.InsertUser()
}
