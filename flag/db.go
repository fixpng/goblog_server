package flag

import (
	"gvb_server/global"
	"gvb_server/models"
)

// Makemigrations 迁移表
func Makemigrations() {
	var err error
	//global.DB.SetupJoinTable(&models.UserModel{}, "CollectsModels", &models.UserCollectModel{})
	global.DB.SetupJoinTable(&models.MenuModel{}, "Banners", &models.MenuBannerModel{})
	err = global.DB.Set("gorm:table_options", "ENGINE=InnoDB").
		AutoMigrate(
			&models.BannerModel{},
			&models.TagModel{},
			&models.MessageModel{},
			&models.AdvertModel{},
			&models.UserModel{},
			&models.CommentModel{},
			//&models.ArticleModel{},
			&models.UserCollectModel{},
			&models.MenuModel{},
			&models.MenuBannerModel{},
			&models.FadeBackModel{},
			&models.LoginDataModel{},
		)
	if err != nil {
		global.Log.Error("[ error ] 生成数据库表结构失败！")
		return
	}
	global.Log.Infof("[ success ] 生成数据库表结构成功！")

}
