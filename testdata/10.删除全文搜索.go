package main

import (
	"gvb_server/core"
	"gvb_server/global"
	"gvb_server/service/es_ser"
)

func main() {
	core.InitConf()
	global.ESClient = core.EsConnect()
	es_ser.DeleteFullTextByArticleID("nOaL0ocB_M76BrAWCR03")
}
