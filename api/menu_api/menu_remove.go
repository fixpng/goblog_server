package menu_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"
)

// MenuRemoveView 批量删除菜单
// @Tags 菜单管理
// @Summary 批量删除菜单
// @Description 批量删除菜单
// @Param data body models.RemoveRequest    true  "菜单id列表"
// @Router /api/Menus [delete]
// @Produce json
// @Success 200 {object} res.Response{data=string}
func (MenuApi) MenuRemoveView(c *gin.Context) {
	var cr models.RemoveRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	var MenuList []models.MenuModel
	count := global.DB.Find(&MenuList, cr.IDList).RowsAffected
	if count == 0 {
		res.FailWithMessage("菜单不存在", c)
		return
	}

	// 事务
	err = global.DB.Transaction(func(tx *gorm.DB) error {
		err = global.DB.Model(&MenuList).Association("Banners").Clear()
		if err != nil {
			global.Log.Error(err)
			return err
		}
		err = global.DB.Delete(&MenuList).Error
		if err != nil {
			global.Log.Error(err)
			return err
		}
		return nil
	})
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("删除菜单失败", c)
		return
	}

	res.OkWithMessage(fmt.Sprintf("共删除 %d 个菜单", count), c)
}
