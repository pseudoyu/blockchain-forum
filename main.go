package main

import (
	"blockchainguide_app/controller"
	"blockchainguide_app/dao/mysql"
	"blockchainguide_app/dao/redis"
	"blockchainguide_app/logger"
	"blockchainguide_app/pkg/snowflake"
	"blockchainguide_app/router"
	"blockchainguide_app/settings"
	"fmt"
)

// Go Web开发通用脚手架模板

func main() {
	//var confFile string
	//flag.StringVar(&confFile, "conf", "./conf/config.yaml", "配置文件")
	//flag.Parse()
	// 加载配置
	if err := settings.Init(); err != nil {
		fmt.Printf("load config failed, err:%v\n", err)
		return
	}
	if err := logger.Init(settings.Conf.LogConfig, settings.Conf.Mode); err != nil {
		fmt.Printf("init logger failed, err:%v\n", err)
		return
	}
	if err := mysql.Init(settings.Conf.MySQLConfig); err != nil {
		fmt.Printf("init mysql failed, err:%v\n", err)
		return
	}
	defer mysql.Close() // 程序退出关闭数据库连接
	if err := redis.Init(settings.Conf.RedisConfig); err != nil {
		fmt.Printf("init redis failed, err:%v\n", err)
		return
	}
	defer redis.Close()
	if err := snowflake.Init(1); err != nil {
		fmt.Printf("init snowflake failed, err:%v\n", err)
		return
	}

	// 初始化gin框架内置的翻译器
	if err := controller.InitTrans("zh"); err != nil {
		fmt.Printf("init trans failed, err:%v\n", err)
		return
	}
	// 注册路由
	r := router.Setup(settings.Conf.Mode)
	err := r.Run(fmt.Sprintf(":%d", settings.Conf.Port))
	if err != nil {
		fmt.Printf("run server failed, err:%v\n", err)
		return
	}
}
