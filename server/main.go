package main

import (
	"server/core"
	"server/global"
	"server/initialize"
)

// @title Swagger Example API
// @version 0.0.1
// @description This is a sample Server pets
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name x-token
// @BasePath /
func main() {
	// 初始化Viper 无参数从config = utils.ConfigEnv = "GVA_CONFIG"= config.Server
	global.GVA_VP = core.Viper()
	global.GVA_LOG = core.Zap()              // 初始化zap日志库
	global.GVA_DB = initialize.Gorm()        // gorm连接数据库
	initialize.PostgresTables(global.GVA_DB) // 初始化表
	// 程序结束前关闭数据库链接
	db, _ := global.GVA_DB.DB()
	defer db.Close()

	//core.RunWindowsServer()
}
