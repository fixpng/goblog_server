package models

import "time"

// UserCollectModel 自定义第三张表 记录用户什么时候收藏了什么文章
type UserCollectModel struct {
	UserID       uint         `gorm:"primaryKey" json:"user_id"`
	UserModel    UserModel    `gorm:"foreignKey:UserID"`
	ArticleID    uint         `gorm:"primaryKey" json:"article_id"`
	ArticleModel ArticleModel `gorm:"foreignKey:ArticleID"`
	CreatedAt    time.Time
}
