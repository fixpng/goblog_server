package routers

import "gvb_server/api"

func (router RouterGroup) NewsRouter() {
	app := api.ApiGroupApp.NewsApi
	router.POST("news", app.NewsListView)

}
