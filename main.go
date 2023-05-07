package main

import (
	"gvb_server/core"
	_ "gvb_server/docs"
	"gvb_server/flag"
	"gvb_server/global"
	"gvb_server/routers"
	"gvb_server/utils"
)

// @title gvb_server API文档
// @version 1.0
// @description API文档
// @host 127.0.0.01:8080
// @BasePath /
func main() {
	// 读取配置文件
	core.InitConf()
	// 初始化日志
	global.Log = core.InitLogger()
	// 连接数据库
	global.DB = core.InitGorm()
	// ip地址数据库加载
	core.InitAddrDB()
	defer global.AddrDB.Close()

	// 命令行参数绑定
	option := flag.Parse()
	if flag.IsWebStop(option) {
		flag.SwitchOption(option)
		return
	}

	// 连接redis
	global.Redis = core.ConnectRedis()
	// 连接es
	global.ESClient = core.EsConnect()

	router := routers.InitRouter()

	addr := global.Config.System.Addr()
	utils.PrintSystem()
	err := router.Run(addr)
	if err != nil {
		global.Log.Fatalf(err.Error())
	}
}
