package user_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/ctype"
	"gvb_server/models/res"
	"gvb_server/plugins/qq"
	"gvb_server/utils"
	"gvb_server/utils/jwts"
	"gvb_server/utils/pwd"
	"gvb_server/utils/random"
)

// QQLoginView QQ登录
// @Tags 用户管理
// @Summary QQ登录
// @Description QQ登录
// @Param limit query string false "表示单个参数"
// @Router /api/qq_login [post]
// @Produce json
// @Success 200 {object} res.Response{}
func (UserApi) QQLoginView(c *gin.Context) {
	code := c.Query("code")
	if code == "" {
		res.FailWithMessage("没有code", c)
		return
	}
	fmt.Println(code)
	qqInfo, err := qq.NewQQLogin(code)
	if err != nil {
		res.FailWithMessage(err.Error(), c)
		return
	}
	fmt.Println(qqInfo)
	res.OkWithData(qqInfo, c)

	openID := qqInfo.OpenID
	// 根据openID判断用户是否存在
	var user models.UserModel
	err = global.DB.Take(&user, "token = ?", openID).Error
	ip, addr := utils.GetAddrByGin(c)
	if err != nil {
		// 不存在，就注册
		hashPwd := pwd.HashPwd(random.RandString(16))
		user = models.UserModel{
			NickName:   qqInfo.Nickname,
			UserName:   openID,  // qq登录，邮箱+密码
			Password:   hashPwd, // 随机生成16位密码
			Avatar:     qqInfo.Avatar,
			Addr:       addr, // 根据ip算地址
			Token:      openID,
			IP:         ip,
			Role:       ctype.PermissionUser,
			SignStatus: ctype.SignQQ,
		}
		err = global.DB.Create(&user).Error
		if err != nil {
			global.Log.Error(err)
			res.FailWithMessage("注册失败", c)
			return
		}
	}

	// 登陆操作
	// 登录成功，生成token
	token, err := jwts.GenToken(jwts.JwtPayLoad{
		NickName: user.NickName,
		Role:     int(user.Role),
		UserID:   user.ID,
	})
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("token生成失败", c)
		return
	}
	global.DB.Create(&models.LoginDataModel{
		UserID:    user.ID,
		IP:        ip,
		NickName:  user.NickName,
		Token:     token,
		Device:    "",
		Addr:      addr,
		LoginType: ctype.SignQQ,
	})
	res.OkWithData(token, c)
}
