package routers

import "gvb_server/api"

func (router RouterGroup) DataRouter() {
	app := api.ApiGroupApp.DataApi
	router.GET("data_seven_login", app.DataSevenLogin)
	router.GET("data_sum", app.DataSumView)

}
