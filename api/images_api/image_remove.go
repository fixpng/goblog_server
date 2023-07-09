package images_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"
)

// ImageRemoveView 批量删除图片
// @Tags 图片管理
// @Summary 批量删除图片
// @Description 批量删除图片
// @Param data body models.RemoveRequest    true  "图片id列表"
// @Router /api/images [delete]
// @Produce json
// @Success 200 {object} res.Response{data=string}
func (ImagesApi) ImageRemoveView(c *gin.Context) {
	var cr models.RemoveRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	var imageList []models.BannerModel
	count := global.DB.Where("id in (?)", cr.IDList).Find(&imageList, cr.IDList).RowsAffected
	if count == 0 {
		res.FailWithMessage("文件不存在", c)
		return
	}
	global.DB.Where("id in (?)", cr.IDList).Delete(&imageList)
	res.OkWithMessage(fmt.Sprintf("共删除 %d 张图片", count), c)
}
