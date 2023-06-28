package routers

import (
	"gvb_server/api"
	"gvb_server/middleware"
)

func (router RouterGroup) ArticleRouter() {
	app := api.ApiGroupApp.ArticleApi
	router.POST("articles", middleware.JwtAdmin(), app.ArticleCreateView)                    // 创建文章
	router.GET("articles", app.ArticleListView)                                              // 文章列表
	router.GET("article_id_title", app.ArticleIDTitleListView)                               // 文章id-title列表
	router.GET("categorys", app.ArticleCategoryListView)                                     // 文章分类列表
	router.GET("articles/detail", app.ArticleDetailByTitleView)                              //文章标题查内容
	router.GET("articles/calendar", app.ArticleCalendarView)                                 // 文章时间聚合搜索
	router.GET("articles/tags", app.ArticleTagListView)                                      // 文章标签列表
	router.PUT("articles", middleware.JwtAdmin(), app.ArticleUpdateView)                     // 更新文章
	router.DELETE("articles", middleware.JwtAdmin(), app.ArticleRemoveView)                  // 批量删除文章
	router.POST("articles/collects", middleware.JwtAuth(), app.ArticleCollCreateView)        // 收藏/取消收藏文章
	router.GET("articles/collects", middleware.JwtAuth(), app.ArticleCollListView)           // 用户收藏的文章列表
	router.DELETE("articles/collects", middleware.JwtAuth(), app.ArticleCollBatchRemoveView) // 批量删除文章收藏
	router.GET("articles/text", app.FullTextContextView)                                     // 全文搜索
	router.POST("article/digg", app.ArticleDiggView)                                         // 文章点赞
	router.GET("articles/content/:id", app.ArticleContentView)                               // 文章正文
	router.GET("articles/:id", app.ArticleDetailView)                                        // id查询文章详情,放最后一个,避免覆盖其他路由
}
