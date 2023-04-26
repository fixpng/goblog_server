package common

import (
	"gorm.io/gorm"
	"gvb_server/global"
	"gvb_server/models"
)

type Option struct {
	models.PageInfo
	Debug bool
	Likes []string // 待优化
}

func ComList[T any](model T, option Option) (list []T, count int64, err error) {

	DB := global.DB

	if option.Debug {
		DB = global.DB.Session(&gorm.Session{Logger: global.MysqlLog})
	}
	if option.Sort == "" {
		option.Sort = "created_at desc" // 默认按照创建时间从大到小
	}

	likes := option.Likes[0]
	global.Log.Infof(likes) // 待优化
	query := DB.Where(model)
	count = query.Find(&list).RowsAffected
	offset := (option.Page - 1) * option.Limit
	if offset < 0 {
		offset = 0
	}

	err = query.Limit(option.Limit).Offset(offset).Order(option.Sort).Find(&list).Error

	return list, count, err
}
