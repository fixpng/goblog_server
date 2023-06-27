package routers

import (
	"gvb_server/api"
	"gvb_server/middleware"
)

func (router RouterGroup) CommentRouter() {
	app := api.ApiGroupApp.CommentApi
	router.POST("comments", middleware.JwtAuth(), app.CommentCreateView)
	router.GET("comments_all", app.CommentListAllView)
	router.POST("comments/:id", app.CommentDigg)
	router.DELETE("comments/:id", middleware.JwtAuth(), app.CommentRemoveView)
	router.GET("comments/:id", app.CommentListView)
}
