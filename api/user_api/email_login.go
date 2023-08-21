package user_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/ctype"
	"gvb_server/models/res"
	"gvb_server/plugins/log_stash"
	"gvb_server/utils"
	"gvb_server/utils/jwts"
	"gvb_server/utils/pwd"
)

type EmailLoginRequest struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

// EmailLoginView 邮箱登录，返回token
// @Tags 用户管理
// @Summary 邮箱登录，返回token
// @Description 邮箱登录，返回token
// @Param data body EmailLoginRequest    true  "表示多个参数"
// @Router /api/email_login [post]
// @Produce json
// @Success 200 {object} res.Response{}
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
		log.Warn(fmt.Sprintf("%s 用户名密码错误", cr.UserName))
		res.FailWithMessage("用户名或密码错误", c)
		return
	}
	// 登录成功，生成token
	token, err := jwts.GenToken(jwts.JwtPayLoad{
		NickName: userModel.NickName,
		Role:     int(userModel.Role),
		UserID:   userModel.ID,
		Avatar:   userModel.Avatar,
	})
	if err != nil {
		global.Log.Error(err)
		log.Error(fmt.Sprintf("token生成失败 %s", err.Error()))
		res.FailWithMessage("token生成失败", c)
		return
	}
	ip, addr := utils.GetAddrByGin(c)
	log = log_stash.New(c.ClientIP(), token)
	log.Info("登录成功")
	global.DB.Create(&models.LoginDataModel{
		UserID:    userModel.ID,
		IP:        ip,
		NickName:  userModel.NickName,
		Token:     token,
		Device:    "",
		Addr:      addr,
		LoginType: ctype.SignEmail,
	})
	res.OkWithData(token, c)

}
