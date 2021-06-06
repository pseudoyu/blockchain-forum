package logic

import (
	"blockchainguide_app/dao/mysql"
	"blockchainguide_app/models"
	"blockchainguide_app/pkg/jwt"
	"blockchainguide_app/pkg/snowflake"
)

// 存放业务逻辑的代码

// SignUp 用户注册业务逻辑
func SignUp(p *models.ParamSignUp) (err error) {
	// 1. 判断用户存不存在
	if err = mysql.CheckUserExist(p.Username); err != nil {
		// 数据库查询出错
		return err
	}
	// 2. 生成UID
	userID, _ := snowflake.GenID()
	// 构造一个User实例
	user := &models.User{
		UserID:   userID,
		Username: p.Username,
		Password: p.Password,
	}
	// 3. 保存进数据库
	return mysql.InsertUser(user)
}

// Login 用户登录业务逻辑
func Login(p *models.ParamLogin) (token string, err error) {
	user := &models.User{
		Username: p.Username,
		Password: p.Password,
	}

	// 传递指针，拿到UserID
	if err := mysql.Login(user); err != nil {
		return "", err
	}

	// 生成JWT
	return jwt.GenToken(user.UserID, user.Username)
}
