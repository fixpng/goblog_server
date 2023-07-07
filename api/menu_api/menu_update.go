package menu_api

import (
	"fmt"
	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"
)

// MenuUpdateView 更新菜单
// @Tags 菜单管理
// @Summary 更新菜单
// @Description 更新菜单
// @Param data body MenuRequest    true  "菜单的一些参数"
// @Param id path int true "id"
// @Router /api/menus/{id} [put]
// @Produce json
// @Success 200 {object} res.Response{data=string}
func (MenuApi) MenuUpdateView(c *gin.Context) {
	var cr MenuRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}

	// 判断菜单是否存在
	id := c.Param("id")
	var menuModel models.MenuModel
	err = global.DB.Take(&menuModel, id).Error
	if err != nil {
		res.FailWithMessage("菜单不存在", c)
		return
	}
	// 先把之前的banner清空
	global.DB.Model(&menuModel).Association("Banners").Clear()

	// 如果选择了banner,那就添加
	if len(cr.ImageSortList) > 0 {
		// 操作第三张表
		var bannerList []models.MenuBannerModel
		for _, sort := range cr.ImageSortList {
			bannerList = append(bannerList, models.MenuBannerModel{
				MenuID:   menuModel.ID,
				BannerID: sort.ImageID,
				Sort:     sort.Sort,
			})
		}
		err = global.DB.Create(&bannerList).Error
		if err != nil {
			global.Log.Error(err)
			res.FailWithMessage("创建菜单图片失败", c)
			return
		}
	}

	// 普通更新
	// 结构体转map的第三方包
	maps := structs.Map(&cr)
	fmt.Println(maps)
	err = global.DB.Model(&menuModel).Updates(maps).Error

	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("修改菜单失败", c)
		return
	}

	res.OkWithMessage("修改菜单成功", c)

}
