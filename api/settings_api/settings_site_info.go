package settings_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models/res"
)

// SettingsSiteInfoView 显示网站信息配置
// @Tags 配置管理
// @Summary 显示网站信息配置
// @Description 显示网站信息配置
// @Router /api/settings/site [get]
// @Produce json
// @Success 200 {object} res.Response{data=config.SiteInfo}
func (s SettingApi) SettingsSiteInfoView(c *gin.Context) {
	res.OkWithData(global.Config.SiteInfo, c)
}
