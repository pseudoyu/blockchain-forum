package router

import (
	"blockchainguide_app/controller"
	"blockchainguide_app/middlewares"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Setup(mode string) *gin.Engine {
	if mode == gin.ReleaseMode {
		// gin设置成发布模式
		gin.SetMode(gin.ReleaseMode)
	}

	//r := gin.New()
	//r.Use(logger.GinLogger(), logger.GinRecovery(true))

	r := gin.Default()
	v1 := r.Group("/api/v1")

	// 注册业务路由
	v1.POST("/signup", controller.SignUpHandler)

	// 登录业务路由
	v1.POST("/login", controller.LoginHandler)

	// 应用JWT认证中间件
	v1.Use(middlewares.JWTAuthMiddleware())
	{
		v1.GET("/community", controller.CommunityHandler)
		v1.GET("/community/:id", controller.CommunityDetailHandler)

		v1.GET("/ping", func(c *gin.Context) {
			c.String(http.StatusOK, "pong")
		})

	}

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "404",
		})
	})
	return r
}
