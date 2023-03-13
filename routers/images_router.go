package routers

import "gvb_server/api"

func (router RouterGroup) ImagesRouter() {
	app := api.ApiGroupApp.ImagesApi
	router.GET("images", app.ImageListView)
	router.POST("images", app.ImageUploadView)
}
