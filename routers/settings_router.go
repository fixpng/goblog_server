package routers

import (
	"gvb_server/api"
	"gvb_server/middleware"
)

func (router RouterGroup) SettingsRouter() {
	settingsApi := api.ApiGroupApp.SettingApi
	router.GET("settings/site", settingsApi.SettingsSiteInfoView)
	router.PUT("settings/site", middleware.JwtAdmin(), settingsApi.SettingsSiteUpdateView)
	router.GET("settings/:name", middleware.JwtAdmin(), settingsApi.SettingsInfoView)
	router.PUT("settings/:name", middleware.JwtAdmin(), settingsApi.SettingsUpdateView)
}
