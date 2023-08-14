package cron_ser

import (
	"gorm.io/gorm"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/service/redis_ser"
)

// SyncCommentData 同步redis评论数据到数据库
func SyncCommentData() {
	commentDiggInfo := redis_ser.NewCommentDigg().GetInfo()
	for key, count := range commentDiggInfo {
		var comment models.CommentModel
		err := global.DB.Take(&comment, key).Error
		if err != nil {
			global.Log.Error(err)
			continue
		}
		err = global.DB.Model(&comment).
			Update("digg_count", gorm.Expr("digg + ?", count)).Error
		if err != nil {
			global.Log.Error(err)
			continue
		}
		global.Log.Infof("%s 更新成功 点赞数为：%d", comment.Content, comment.CommentCount)
	}
	redis_ser.NewCommentDigg().Clear()
}
