package routers

import (
	"gvb_server/api"
)

func (router RouterGroup) ChatRouter() {
	app := api.ApiGroupApp.ChatApi
	router.GET("chat_groups", app.ChatGroupView)
}
