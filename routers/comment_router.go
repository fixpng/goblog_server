package routers

import (
	"gvb_server/api"
	"gvb_server/middleware"
)

func (router RouterGroup) CommentRouter() {
	app := api.ApiGroupApp.CommentApi
	router.POST("comments", middleware.JwtAuth(), app.CommentCreateView)
	router.GET("comments", app.CommentListView)
	router.GET("comments/:id", app.CommentDigg)
	router.DELETE("comments/:id", middleware.JwtAuth(), app.CommentRemoveView)
}
