package models

import "gvb_server/models/ctype"

// UserModel 用户表
type UserModel struct {
	MODEL
	NickName   string           `gorm:"size:36" json:"nick_name,select(c|info)"`          // 昵称
	UserName   string           `gorm:"size:36" json:"user_name"`                         // 用户名
	Password   string           `gorm:"size:128" json:"-"`                                // 密码
	Avatar     string           `gorm:"size:256" json:"avatar,select(c|info)"`            // 头像
	Email      string           `gorm:"size:128" json:"email,select(info)"`               // 邮箱
	Tel        string           `gorm:"size:18" json:"tel"`                               // 手机号
	Addr       string           `gorm:"size:64" json:"addr,select(c|info)"`               // 地址
	Token      string           `gorm:"size:64" json:"token"`                             // 其他平台的唯一id
	IP         string           `gorm:"size:20" json:"ip,select(c)"`                      // ip地址
	Role       ctype.Role       `gorm:"size:4;default:1" json:"role,select(info)"`        // 权限  1 管理员  2 普通用户  3 游客
	SignStatus ctype.SignStatus `gorm:"type=smallint(6)" json:"sign_status,select(info)"` // 注册来源
	Integral   int              `gorm:"default:0" json:"integral,select(info)"`           // 积分
	Sign       string           `gorm:"size:128" json:"sign,select(info)"`                // 签名
	Link       string           `gorm:"size:128" json:"link,select(info)"`                // 链接地址
}
