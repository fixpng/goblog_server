package user_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/ctype"
	"gvb_server/models/res"
	"gvb_server/plugins/log_stash"
	"gvb_server/utils/jwts"
	"gvb_server/utils/pwd"
)

type EmailLoginRequest struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

func (UserApi) EmailLoginView(c *gin.Context) {
	var cr EmailLoginRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}

	log := log_stash.NewLogByGin(c)

	var userModel models.UserModel
	err = global.DB.Take(&userModel, "user_name = ? or email = ?", cr.UserName, cr.UserName).Error
	if err != nil {
		// 没找到
		global.Log.Warn("用户名不存在")
		log.Warn(fmt.Sprintf("%s 用户名不存在", cr.UserName))
		res.FailWithMessage("用户名或密码错误", c)
		return
	}
	// 校验密码
	isCheck := pwd.CheckPwd(userModel.Password, cr.Password)
	if !isCheck {
		global.Log.Warn("用户名密码错误")
		log.Warn(fmt.Sprintf("用户名密码错误 %s %s", cr.UserName, cr.Password))
		res.FailWithMessage("用户名或密码错误", c)
		return
	}
	// 登录成功，生成token
	token, err := jwts.GenToken(jwts.JwtPayLoad{
		NickName: userModel.NickName,
		Role:     int(userModel.Role),
		UserID:   userModel.ID,
	})
	if err != nil {
		global.Log.Error(err)
		log.Error(fmt.Sprintf("token生成失败 %s", err.Error()))
		res.FailWithMessage("token生成失败", c)
		return
	}
	log = log_stash.New(c.ClientIP(), token)
	log.Info("登录成功")
	global.DB.Create(&models.LoginDataModel{
		UserID:    userModel.ID,
		IP:        c.ClientIP(),
		NickName:  userModel.NickName,
		Token:     token,
		Device:    "",
		Addr:      "内网",
		LoginType: ctype.SignEmail,
	})
	res.OkWithData(token, c)

}
