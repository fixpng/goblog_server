package models

// TagModel 标签表
type TagModel struct {
	MODEL
	Title    string         `gorm:"size:16" json:"title"`           // 标签的名称
	Articles []ArticleModel `gorm:"many2many:article_tag" json:"-"` // 关联该标签的文章列表
}
