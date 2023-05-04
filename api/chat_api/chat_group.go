package chat_api

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
	"gvb_server/models/res"
	"net/http"
	"strings"
	"time"
)

var ConnGroupMap = map[string]*websocket.Conn{}

type MsgType int

const (
	TextMsg    MsgType = 1
	ImageMsg   MsgType = 2
	SystemMsg  MsgType = 3
	InRoomMsg  MsgType = 4
	OutRoomMsg MsgType = 5
)

// GroupRequest 群聊入参
type GroupRequest struct {
	NickName string  `json:"nick_name"` // 前端自己生成
	Avatar   string  `json:"avatar"`    // 头像
	Content  string  `json:"content"`   // 聊天的内容
	MsgType  MsgType `json:"msg_type"`  // 聊天类型
}

// GroupResponse 群聊出参
type GroupResponse struct {
	GroupRequest
	Date time.Time `json:"date"` // 消息发送时间
}

func (ChatApi) ChatGroupView(c *gin.Context) {
	var upGrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			// 鉴权 true表示放行，false表示拦截
			return true
		},
	}
	// 将http升级至websocket
	conn, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	addr := conn.RemoteAddr().String()
	ConnGroupMap[addr] = conn
	logrus.Infof("%s 链接成功", addr)
	for {
		// 消息类型，消息，错误
		_, p, err := conn.ReadMessage()
		if err != nil {
			// 用户断开聊天
			SendGroupMsg(GroupResponse{
				GroupRequest: GroupRequest{
					Content: fmt.Sprintf("%s 离开聊天室", addr),
				},
				Date: time.Now(),
			})
			break
		}
		// 请求参数绑定
		var request GroupRequest
		err = json.Unmarshal(p, &request)
		if err != nil {
			// 参数绑定失败
			continue
		}
		// 内容不能为空
		if strings.TrimSpace(request.Avatar) == "" || strings.TrimSpace(request.NickName) == "" {
			continue
		}

		// 判断类型，分发逻辑
		switch request.MsgType {
		case TextMsg:
			if strings.TrimSpace(request.Content) == "" {
				continue
			}
			SendGroupMsg(GroupResponse{
				GroupRequest: request,
				Date:         time.Now(),
			})
		case InRoomMsg:
			request.Content = fmt.Sprintf("%s 进入聊天室", request.NickName)
			SendGroupMsg(GroupResponse{
				GroupRequest: request,
				Date:         time.Now(),
			})
		}
	}
	defer conn.Close()
	delete(ConnGroupMap, addr)
}

// SendGroupMsg 消息发送
func SendGroupMsg(response GroupResponse) {
	byteData, _ := json.Marshal(response)
	for _, conn := range ConnGroupMap {
		conn.WriteMessage(websocket.TextMessage, byteData)
	}
}
