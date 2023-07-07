package routers

import (
	"gvb_server/api"
	"gvb_server/middleware"
)

func (router RouterGroup) MenuRouter() {
	app := api.ApiGroupApp.MenuApi
	router.POST("menus", middleware.JwtAuth(), app.MenuCreateView)
	router.GET("menus", app.MenuListView)
	router.GET("menu_names", app.MenuNameList)
	router.PUT("menus/:id", middleware.JwtAuth(), app.MenuUpdateView)
	router.DELETE("menus", middleware.JwtAuth(), app.MenuRemoveView)
	router.GET("menus/detail", app.MenuDetailByPathView)
	router.GET("menus/:id", app.MenuDetailView)

}
