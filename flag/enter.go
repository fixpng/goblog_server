package flag

import (
	sys_flag "flag"
	"github.com/fatih/structs"
	"gvb_server/core"
	"gvb_server/global"
)

type Option struct {
	DB   bool
	User string // -u admin -u user
	ES   string // -es create -es delete
}

// Parse 解析命令行参数
func Parse() Option {
	db := sys_flag.Bool("db", false, "初始化数据库")
	user := sys_flag.String("u", "", "创建用户")
	es := sys_flag.String("es", "", "es操作")
	// 解析命令行参数写入注册的flag里
	sys_flag.Parse()
	return Option{
		DB:   *db,
		User: *user,
		ES:   *es,
	}
}

// IsWebStop 是否停止Web项目
func IsWebStop(option Option) (f bool) {
	maps := structs.Map(&option)
	for _, v := range maps {
		switch val := v.(type) {
		case string:
			if v != "" {
				f = true
			}
		case bool:
			if val == true {
				f = true
			}
		}
	}
	return f
}

// SwitchOption 根据命令执行不同的函数
func SwitchOption(option Option) {
	if option.DB {
		Makemigrations()
		return
	}

	if option.User == "admin" || option.User == "user" {
		CreateUser(option.User)
		return
	}

	if option.ES == "create" {
		// 连接es
		global.ESClient = core.EsConnect()
		EsCreateIndex()
	}
	//fmt.Printf("%#v\n", option.User)
	//sys_flag.Usage()

}
