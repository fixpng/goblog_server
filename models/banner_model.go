package models

import (
	"gorm.io/gorm"
	"gvb_server/global"
	"gvb_server/models/ctype"
	"os"
)

// BannerModel 图片表
type BannerModel struct {
	MODEL
	Path      string          `json:"path"`                        // 图片路径
	Hash      string          `json:"hash"`                        // 图片的hash值，用于判断重复图片
	Name      string          `gorm:"size:38" json:"name"`         // 图片名称
	ImageType ctype.ImageType `gorm:"default:1" json:"image_type"` // 图片的类型，本地还是七牛
}

func (u *BannerModel) BeforeDelete(tx *gorm.DB) (err error) {
	if u.ImageType == ctype.Local {
		// 本地图片，删除，还要删除本地的存储
		err = os.Remove(u.Path[1:])
		if err != nil {
			global.Log.Error(err)
			return err
		}
	}
	return nil
}
