package settings_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models/res"
)

type SettingsUri struct {
	Name string `uri:"name"`
}

// SettingsInfoView 显示配置信息
// @Tags 配置管理
// @Summary 显示配置信息
// @Description 显示配置信息
// @Param data query SettingsUri    false  "查询参数"
// @Router /api/settings/site [get]
// @Produce json
// @Success 200 {object} res.Response{data=string}
func (s SettingApi) SettingsInfoView(c *gin.Context) {
	var cr SettingsUri
	err := c.ShouldBindUri(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	switch cr.Name {
	case "site":
		res.OkWithData(global.Config.SiteInfo, c)
	case "email":
		res.OkWithData(global.Config.Email, c)
	case "qq":
		res.OkWithData(global.Config.QQ, c)
	case "qiniu":
		res.OkWithData(global.Config.QiNiu, c)
	case "jwts":
		res.OkWithData(global.Config.Jwy, c)
	default:
		res.FailWithMessage("没有对应的配置信息", c)
	}

}
