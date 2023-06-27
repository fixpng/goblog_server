package user_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models/res"
)

// QQLoginLinkView 获取qq登录的跳转地址
// @Tags 用户管理
// @Summary 获取qq登录的跳转地址
// @Description 获取qq登录的跳转地址，data就是qq的跳转地址
// @Router /api/qq_login_path [get]
// @Produce json
// @Success 200 {object} res.Response{}
func (UserApi) QQLoginLinkView(c *gin.Context) {
	path := global.Config.QQ.GetPath()
	if path == "" {
		res.FailWithMessage("未配置qq登录地址", c)
		return
	}

	res.OkWithData(path, c)
	return
}
