package routers

import (
	"gvb_server/api"
	"gvb_server/middleware"
)

func (router RouterGroup) LogRouter() {
	app := api.ApiGroupApp.LogApi
	router.GET("logs", middleware.JwtAdmin(), app.LogListView)
}
