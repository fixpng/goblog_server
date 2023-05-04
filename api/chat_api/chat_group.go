package chat_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"gvb_server/models/res"
	"net/http"
)

func (ChatApi) ChatGroupView(c *gin.Context) {
	var upGrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			// 鉴权 true表示放行，false表示拦截
			return true
		},
	}
	// 将http升级至websocket
	conn, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	fmt.Println(err)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	for {
		// 消息类型，消息，错误
		_, p, err := conn.ReadMessage()
		if err != nil {
			// 用户断开聊天
			break
		}
		fmt.Println(string(p))
		// 发送消息
		conn.WriteMessage(websocket.TextMessage, []byte("xxx"))
	}
	defer conn.Close()

}
