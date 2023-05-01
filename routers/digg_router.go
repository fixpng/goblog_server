package routers

import "gvb_server/api"

func (router RouterGroup) DiggRouter() {
	app := api.ApiGroupApp.DiggApi
	router.POST("digg/article", app.DiggArticleView)
}
