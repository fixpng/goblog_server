package main

import (
	"fmt"
	"gvb_server/core"
	"gvb_server/global"
	"gvb_server/utils/jwts"
)

func main() {
	core.InitConf()
	global.Log = core.InitLogger()

	fmt.Println(global.Config.Jwt.Secret)
	token, err := jwts.GenToken(jwts.JwtPayLoad{
		//Username: "xixi",
		NickName: "xxx",
		Role:     1,
		UserID:   1,
		Avatar:   "xxx",
	})
	fmt.Println(token, err)

	claims, err := jwts.ParseToken(token)
	fmt.Println(claims, err)
}
