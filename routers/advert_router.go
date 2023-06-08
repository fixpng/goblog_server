package routers

import (
	"gvb_server/api"
	"gvb_server/middleware"
)

func (router RouterGroup) AdvertRouter() {
	app := api.ApiGroupApp.AdvertApi
	router.POST("adverts", middleware.JwtAdmin(), app.AdvertCreateView)
	router.GET("adverts", app.AdvertListView)
	router.PUT("adverts/:id", middleware.JwtAdmin(), app.AdvertUpdateView)
	router.DELETE("adverts", middleware.JwtAdmin(), app.AdvertRemoveView)
}
