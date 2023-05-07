package utils

import "gvb_server/global"

func PrintSystem() {
	ip := global.Config.System.Host
	port := global.Config.System.Port
	if ip == "0.0.0.0" {
		ipList := GetIPList()
		for _, i := range ipList {
			global.Log.Infof("gvb_server 运行在：http://%s:%d/api", i, port)
			global.Log.Infof("gvb_server api文档 运行在：http://%s:%d/swagger/index.html#", i, port)
		}
	} else {
		global.Log.Infof("gvb_server 运行在：http://%s:%d/api", ip, port)
		global.Log.Infof("gvb_server api文档 运行在：http://%s:%d/swagger/index.html#", ip, port)
	}

}
