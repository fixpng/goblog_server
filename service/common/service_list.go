package common

import (
	"gorm.io/gorm"
	"gvb_server/global"
	"gvb_server/models"
)

type Option struct {
	models.PageInfo
	Debug bool
}

func ComList[T any](model T, option Option) (list []T, count int64, err error) {

	DB := global.DB
	if option.Debug {
		DB = global.DB.Session(&gorm.Session{Logger: global.MysqlLog})
	}

	count = DB.Debug().Select("id").Find(&list).RowsAffected
	offset := (option.Page - 1) * option.Limit
	if offset < 0 {
		offset = 0
	}
	err = DB.Limit(option.Limit).Offset(offset).Find(&list).Error

	return list, count, err
}
