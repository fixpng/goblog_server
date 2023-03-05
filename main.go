package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"gvb_server/core"
	"gvb_server/global"
)

func main() {
	// 读取配置文件
	core.InitConf()
	// 初始化日志
	global.Log = core.InitLogger()
	logrus.Warnf("xxxa")
	logrus.Errorf("xxxa")
	logrus.Infof("xxxa")
	// 连接数据库
	global.DB = core.InitGorm()
	fmt.Println(global.DB)
}
