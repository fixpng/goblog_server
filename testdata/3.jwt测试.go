package main

import (
	"fmt"
	"gvb_server/core"
	"gvb_server/global"
	"gvb_server/utils/jwts"
)

func main() {
	core.SetYaml()
	global.Log = core.InitLogger()

	token, err := jwts.GenToken(jwts.JwtPayLoad{
		Username: "xixi",
		NickName: "xxx",
		Role:     1,
		UserID:   1,
	})
	fmt.Println(token, err)
}
