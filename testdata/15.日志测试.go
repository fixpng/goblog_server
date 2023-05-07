package main

import (
	"fmt"
	"gvb_server/core"
	"gvb_server/global"
	"gvb_server/plugins/log_stash"
)

func main() {
	core.InitConf()
	// 初始化日志
	global.Log = core.InitLogger()
	// 连接数据库
	global.DB = core.InitGorm()

	log := log_stash.New("192.168.100.101", "xxxxx")
	log.Error(fmt.Sprintf("%s 我日", "嘻嘻嘻"))
}
