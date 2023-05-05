package models

import "gvb_server/models/ctype"

type ChatModel struct {
	MODEL
	NickName string        `gorm:"size:15" json:"nick_name"` // 前端自己生成
	Avatar   string        `gorm:"size:128" json:"avatar"`   // 头像
	Content  string        `gorm:"size:256" json:"content"`  // 聊天的内容
	IP       string        `gorm:"size:32" json:"ip"`        // ip
	Addr     string        `gorm:"size:64" json:"addr"`      // 地址
	MsgType  ctype.MsgType `gorm:"size:4" json:"msg_type"`   // 聊天类型
}
