package user_api

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models/res"
	"gvb_server/plugins/email"
	"gvb_server/service/user_ser"
	"gvb_server/utils/random"
)

type UserRegisterRequest struct {
	NickName string  `json:"nick_name" binding:"required" msg:"请输入昵称"`  // 昵称
	UserName string  `json:"user_name" binding:"required" msg:"请输入用户名"` // 用户名
	Password string  `json:"password"`                                  // 密码
	Email    string  `json:"email" binding:"required,email" msg:"邮箱非法"`
	Code     *string `json:"code"`
}

// UserRegisterView 用户注册
// @Tags 用户管理
// @Summary 用户注册
// @Description 用户注册
// @Param data body UserRegisterRequest    true  "表示多个参数"
// @Router /api/user_register [post]
// @Produce json
// @Success 200 {object} res.Response{}
func (UserApi) UserRegisterView(c *gin.Context) {
	var cr UserRegisterRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}
	// 用户绑定邮箱，第一次输入是 邮箱
	// 后台会给这个邮箱发验证码
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

	if len(cr.Password) == 0 {
		res.FailWithMessage("请输入密码", c)
		return
	} else if len(cr.Password) < 4 {
		res.FailWithMessage("密码强度过低", c)
		return
	}

	// 权限  1 管理员  2 普通用户  3 游客
	err = user_ser.UserService{}.CreateUser(cr.UserName, cr.NickName, cr.Password, 2, cr.Email, c.ClientIP())
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage(err.Error(), c)
		return
	}

	res.OkWithMessage(fmt.Sprintf("用户%s创建成功!", cr.UserName), c)
	err = email.NewCode().Send(cr.Email, fmt.Sprintf("账号创建成功！用户昵称：%s，用户名：%s，密码：%s，请妥善保存账密", cr.NickName, cr.UserName, cr.Password))
	if err != nil {
		global.Log.Error(err)
	}
	return
}
