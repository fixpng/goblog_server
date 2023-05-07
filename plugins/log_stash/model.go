package log_stash

import "time"

type LogStashModel struct {
	ID        uint      `gorm:"primarykey" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	IP        string    `gorm:"size:32" json:"ip"`
	Addr      string    `gorm:"size:64" json:"addr"`
	Level     Level     `gorm:"size:4" json:"level"`     // 日志等级
	Content   string    `gorm:"size:128" json:"content"` // 日志消息内容
	UserID    uint      `json:"user_id"`                 // 登录用户的用户id，需要自己在查询的时候做关联查询
}
