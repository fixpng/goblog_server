package advert_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"
)

func (AdvertApi) AdvertRemoveView(c *gin.Context) {
	var cr models.RemoveRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	var advertList []models.AdvertModel
	count := global.DB.Where("id in (?)", cr.IDList).Find(&advertList).RowsAffected
	if count == 0 {
		res.FailWithMessage("广告不存在", c)
		return
	}
	global.DB.Delete(&advertList)
	res.OkWithMessage(fmt.Sprintf("共删除 %d 个广告", count), c)
}
