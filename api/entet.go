package api

import "gvb_server/api/settings_api"

type ApiGroup struct {
	SettingApi settings_api.SettingApi
}

var ApiGroupApp = new(ApiGroup)
