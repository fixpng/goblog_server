package routers

import (
	"gvb_server/api"
	"gvb_server/middleware"
)

func (router RouterGroup) UserRouter() {
	app := api.ApiGroupApp.UserApi
	router.POST("email_login", app.EmailLoginView)
	router.GET("users", middleware.JwtAuth(), app.UserListView)
	router.PUT("user_role", middleware.JwtAuth(), app.UserUpdateRoleView)

}
