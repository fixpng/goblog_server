package log_stash

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gvb_server/global"
	"gvb_server/utils"
	"gvb_server/utils/jwts"
)

type Log struct {
	IP     string `json:"ip"`
	Addr   string `json:"addr"`
	UserID uint   `json:"user_id"`
}

func New(ip string, token string) *Log {
	// 解析Token
	claims, err := jwts.ParseToken(token)
	var userID uint
	if err == nil {
		userID = claims.UserID
	}
	addr := utils.GetAddr(ip)
	// 拿到用户id
	return &Log{
		IP:     ip,
		Addr:   addr,
		UserID: userID,
	}
}

func NewLogByGin(c *gin.Context) *Log {
	ip := c.ClientIP()
	token := c.Request.Header.Get("token")
	return New(ip, token)
}

func (l Log) Debug(content string) {
	l.send(DebugLevel, content)
}
func (l Log) Info(content string) {
	l.send(InfoLevel, content)
}
func (l Log) Warn(content string) {
	l.send(WarnLevel, content)
}
func (l Log) Error(content string) {
	l.send(ErrorLevel, content)
}

func (l Log) send(level Level, content string) {
	err := global.DB.Create(&LogStashModel{
		IP:      l.IP,
		Addr:    l.Addr,
		Level:   level,
		Content: content,
		UserID:  l.UserID,
	}).Error
	if err != nil {
		logrus.Error(err)
	}
	fmt.Println(l.IP, l.UserID, l.Addr, content, level)
}

//
//func Debug(ip string, content string) {
//	std.Debug(ip, content)
//}
//func Info(ip string, content string) {
//	std.Info(ip, content)
//}
//func Warn(ip string, content string) {
//	std.Warn(ip, content)
//}
//func Error(ip string, content string) {
//	std.Error(ip, content)
//}
