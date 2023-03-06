package settings_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models/res"
)

func (s SettingApi) SettingsInfoView(c *gin.Context) {
	//res.Ok(map[string]string{}, "xxx", c)
	//res.OkWithData(map[string]string{"id": "xxx"}, c)
	//res.FailWithCode(res.SettingsError, c)
	res.OkWithData(global.Config.SiteInfo, c)
}
