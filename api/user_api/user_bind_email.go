package user_api

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"
	"gvb_server/plugins/email"
	"gvb_server/utils/jwts"
	"gvb_server/utils/pwd"
	"gvb_server/utils/random"
)

type BindEmailRequest struct {
	Email    string  `json:"email" binding:"required,email" msg:"邮箱非法"`
	Code     *string `json:"code"`
	Password string  `json:"password"`
}

// UserBindEmailView 用户绑定邮箱
// @Tags 用户管理
// @Summary 用户绑定邮箱
// @Description 用户绑定邮箱
// @Param data body BindEmailRequest    true  "表示多个参数"
// @Param token header string true "token"
// @Router /api/user_bind_email [post]
// @Produce json
// @Success 200 {object} res.Response{}
func (UserApi) UserBindEmailView(c *gin.Context) {
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwts.CustomClaims)
	// 用户绑定邮箱，第一次输入是 邮箱
	// 后台会给这个邮箱发验证码
	var cr BindEmailRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}
	session := sessions.Default(c)
	if cr.Code == nil {
		// 第一次，后台发验证码
		// 生成4位验证码，将生成的验证码存入session
		code := random.Code(4)
		// 写入session
		session.Set("valid_code", code)
		err = session.Save()
		if err != nil {
			global.Log.Error(err)
			res.FailWithMessage("session错误", c)
			return
		}
		err = email.NewCode().Send(cr.Email, "您的验证码是 "+code)
		if err != nil {
			global.Log.Error(err)
		}
		res.OkWithMessage("验证码已发送，请查收", c)
		return
	}
	// 第二次，用户输入邮箱，验证码，密码
	code := session.Get("valid_code")
	// 校验验证码
	if code != *cr.Code {
		res.FailWithMessage("验证码错误", c)
		return
	}
	// 修改用户的邮箱
	var user models.UserModel
	err = global.DB.Take(&user, claims.UserID).Error
	if err != nil {
		res.FailWithMessage("用户不存在", c)
		return
	}
	if len(cr.Password) < 4 {
		res.FailWithMessage("密码强度过低", c)
		return
	}
	hashPwd := pwd.HashPwd(cr.Password)
	// 第一次的邮箱和第二次的邮箱也要做一致性校验
	err = global.DB.Model(&user).Updates(map[string]any{
		"email":    cr.Email,
		"password": hashPwd,
	}).Error
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("绑定邮箱失败", c)
		return
	}

	// 完成绑定
	res.OkWithMessage("邮箱绑定成功", c)
	err = email.NewNote().Send(cr.Email, fmt.Sprintf("账号更改绑定邮箱成功！用户昵称：%s", claims.NickName))
	if err != nil {
		global.Log.Error(err)
	}
	return

}
