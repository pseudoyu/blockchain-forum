package controller

import (
	"blockchainguide_app/logic"
	"blockchainguide_app/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

func SignUpHandler(c *gin.Context) {
	// 1. 获取参数和参数校验
	p := new(models.ParamSignUp)
	if err := c.ShouldBindJSON(p); err != nil {
		// 请求参数有误，直接返回响应
		zap.L().Error("SignUp with invalid param", zap.Error(err))

		// 用于判断err是不是validator.ValidationErrors类型
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			c.JSON(http.StatusOK, gin.H{
				"msg": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"msg": errs.Translate(trans), // 翻译错误
		})
		return
	}

	// 手动请求对参数进行业务规则的校验
	//if len(p.Username) == 0 || len(p.Password) == 0 || p.RePassword != p.Password {
	//	zap.L().Error("SignUp with invalid param")
	//	c.JSON(http.StatusOK, gin.H{
	//		"msg": "请求参数有误",
	//	})
	//	return
	//}

	fmt.Println(p)
	// 2. 业务处理
	logic.SignUp(p)
	// 3. 返回响应
	c.JSON(http.StatusOK, "请求参数正确")

}
