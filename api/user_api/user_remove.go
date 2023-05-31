package user_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"
)

// UserRemove 批量删除用户
// @Tags 用户管理
// @Summary 批量删除用户
// @Description 批量删除用户
// @Param data body models.RemoveRequest    true  "用户id列表"
// @Param token header string true "token"
// @Router /api/tags [delete]
// @Produce json
// @Success 200 {object} res.Response{data=string}
func (UserApi) UserRemove(c *gin.Context) {
	var cr models.RemoveRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	var UserList []models.UserModel
	count := global.DB.Find(&UserList, cr.IDList).RowsAffected
	if count == 0 {
		res.FailWithMessage("用户不存在", c)
		return
	}

	// 事务
	err = global.DB.Transaction(func(tx *gorm.DB) error {
		// TODO: 删除用户，消息表，评论表，用户收藏的文章，用户发布的文章
		/*
			err = global.DB.Model(&UserList).Association("Banners").Clear()
			if err != nil {
				global.Log.Error(err)
				return err
			}*/
		err = global.DB.Delete(&UserList).Error
		if err != nil {
			global.Log.Error(err)
			return err
		}
		return nil
	})
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("删除用户失败", c)
		return
	}

	res.OkWithMessage(fmt.Sprintf("共删除 %d 个用户", count), c)
}
