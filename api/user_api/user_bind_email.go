package user_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/models/res"
	"gvb_server/plugins/email"
	"gvb_server/utils/random"
)

type BindEmailRequest struct {
	Email    string  `json:"email" binding:"required,email" msg:"邮箱非法"`
	Code     *string `json:"code"`
	Password string  `json:"password"`
}

func (UserApi) UserBindEmailView(c *gin.Context) {
	// 用户绑定邮箱，第一次输入是 邮箱
	// 后台会给这个邮箱发验证码
	var cr BindEmailRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}
	if cr.Code == nil {
		// 第一次，后台发验证码
		// 生成4位验证码，将生成的验证码存入session
		code := random.Code(4)
		// 写入session
		email.NewCode().Send(cr.Email, "你的验证码是 "+code)
	}

	// 第二次，用户输入邮箱，验证码，密码
	// 完成绑定

}
