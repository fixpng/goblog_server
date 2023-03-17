package routers

import "gvb_server/api"

func (router RouterGroup) MenuRouter() {
	app := api.ApiGroupApp.MenuApi
	router.POST("menus", app.MenuCreateView)

}
