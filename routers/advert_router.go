package routers

import "gvb_server/api"

func (router RouterGroup) AdvertRouter() {
	app := api.ApiGroupApp.AdvertApi
	router.POST("adverts", app.AdvertCreateView)
}
