package models

import (
	"context"
	"github.com/sirupsen/logrus"
	"gvb_server/global"
	"gvb_server/models/ctype"
)

// ArticleModel 文章表
type ArticleModel struct {
	ID        string `json:"id" `
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at" `

	Title    string `json:"title"`    // 文章标题
	Abstract string `json:"abstract"` // 文章简介
	Content  string `json:"content"`  // 文章内容

	LookCount     int `json:"look_count"`     // 浏览量
	CommentCount  int `json:"comment_count"`  // 评论量
	DiggCount     int `json:"digg_count"`     // 点赞量
	CollectsCount int `json:"collects_count"` // 收藏量

	UserID       uint   `json:"user_id"`        // 用户id
	UserNickName string `json:"user_nick_name"` // 用户昵称
	UserAvatar   string `json:"user_avatar"`    // 用户头像

	Category string `json:"category"` // 文章分类
	Source   string `json:"source"`   // 文章来源
	Link     string `json:"link"`     // 原文链接

	BannerID  uint   `json:"banner_id"`  // 文章封面ID
	BannerUrl string `json:"banner_url"` // 文章封面

	Tags ctype.Array `json:"tags"` // 文章标签
}

func (ArticleModel) Index() string {
	return "article_index"
}

func (ArticleModel) Mapping() string {
	return `
{
  "settings": {
    "index":{
      "max_result_window": "100000"
    }
  }, 
  "mappings": {
    "properties": {
      "title": { 
        "type": "text"
      },
      "abstract": { 
        "type": "text"
      },
      "content": { 
        "type": "text"
      },
      "look_count": {
        "type": "integer"
      },
      "comment_count": {
        "type": "integer"
      },
      "digg_count": {
        "type": "integer"
      },
      "collects_count": {
        "type": "integer"
      },
      "user_id": {
        "type": "integer"
      },
      "user_nick_name": { 
        "type": "text"
      },
      "user_avatar": { 
        "type": "text"
      },
      "category": { 
        "type": "text"
      },
      "source": { 
        "type": "text"
      },
      "link": { 
        "type": "text"
      },
      "banner_id": {
        "type": "integer"
      },
      "banner_url": { 
        "type": "text"
      },
      "created_at":{
        "type": "date",
        "null_value": "null",
        "format": "[yyyy-MM-dd HH:mm:ss]"
      },
      "updated_at":{
        "type": "date",
        "null_value": "null",
        "format": "[yyyy-MM-dd HH:mm:ss]"
      }
    }
  }
}
`
}

// IndexExists 查索引是否存在
func (a ArticleModel) IndexExists() bool {
	context.Background()
	exists, err := global.ESClient.
		IndexExists(a.Index()).
		Do(context.Background())
	if err != nil {
		logrus.Error(err.Error())
		return exists
	}
	return exists
}

// CreateIndex 创建索引
func (a ArticleModel) CreateIndex() error {
	if a.IndexExists() {
		// 有索引
		a.RemoveIndex()
	}
	// 无索引
	// 创建索引
	createIndex, err := global.ESClient.
		CreateIndex(a.Index()).
		BodyString(a.Mapping()).
		Do(context.Background())
	if err != nil {
		logrus.Error("创建索引失败")
		logrus.Error(err.Error())
		return err
	}
	if !createIndex.Acknowledged {
		global.Log.Error("创建失败")
		return err
	}
	logrus.Infof("索引 %s 创建成功", a.Index())
	return nil

}

// RemoveIndex 删除索引
func (a ArticleModel) RemoveIndex() error {
	logrus.Info("索引存在，删除索引")
	// 删除索引
	indexDelete, err := global.ESClient.DeleteIndex(a.Index()).Do(context.Background())
	if err != nil {
		logrus.Error("删除索引失败")
		logrus.Error(err.Error())
		return err
	}
	if !indexDelete.Acknowledged {
		logrus.Error("删除索引失败")
		return err
	}
	logrus.Info("索引删除成功")
	return nil
}
