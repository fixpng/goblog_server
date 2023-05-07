package routers

import "gvb_server/api"

func (router RouterGroup) DateRouter() {
	app := api.ApiGroupApp.DateApi
	router.GET("date_login", app.SevenLogin)

}
