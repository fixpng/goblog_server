package api

import (
	"gvb_server/api/advert_api"
	"gvb_server/api/images_api"
	"gvb_server/api/menu_api"
	"gvb_server/api/settings_api"
)

type ApiGroup struct {
	SettingApi settings_api.SettingApi
	ImagesApi  images_api.ImagesApi
	AdvertApi  advert_api.AdvertApi
	MenuApi    menu_api.MenuApi
}

var ApiGroupApp = new(ApiGroup)
