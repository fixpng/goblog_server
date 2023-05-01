package routers

import (
	"gvb_server/api"
	"gvb_server/middleware"
)

func (router RouterGroup) ArticleRouter() {
	app := api.ApiGroupApp.ArticleApi
	router.POST("articles", middleware.JwtAuth(), app.ArticleCreateView)
	router.GET("articles", app.ArticleListView)
	router.GET("articles/detail", app.ArticleDetailByTitleView)
	router.GET("articles/calendar", app.ArticleCalendarView)
	router.GET("articles/tags", app.ArticleTagListView)
	router.PUT("articles", app.ArticleUpdateView)
	router.DELETE("articles", app.ArticleRemoveView)
	router.POST("articles/collects", middleware.JwtAuth(), app.ArticleCollCreateView)
	router.GET("articles/collects", middleware.JwtAuth(), app.ArticleCollListView)
	// id查询放最后一个
	router.GET("articles/:id", app.ArticleDetailView)
}
