package main

import (
	"go-admin/config"
	"go-admin/entery"
	"go-admin/router"
)

func main() {

	// 创建必须的缓存目录
	entery.EnteryCheckDir()

	// 检查启动参数
	entery.EnteryCheckEnv()

	// 初始化连接数据库
	config.InitSQLite()

	// 初始化路由
	router.InitRouterIndex()
}
