package advert_api

import (
	"fmt"
	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"
)

// AdvertUpdateView 更新广告
// @Tags 广告管理
// @Summary 更新广告
// @Description 更新广告
// @Param data body AdvertRequest    true  "广告的一些参数"
// @Router /api/adverts/:id [put]
// @Produce json
// @Success 200 {object} res.Response{data=string}
func (AdvertApi) AdvertUpdateView(c *gin.Context) {

	id := c.Param("id")
	var cr AdvertRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}
	// 存在的判断
	var advert models.AdvertModel
	err = global.DB.Take(&advert, id).Error
	if err != nil {
		res.FailWithMessage("广告不存在", c)
		return
	}

	// 结构体转map的第三方包
	maps := structs.Map(&cr)
	fmt.Println(maps)
	err = global.DB.Model(&advert).Updates(maps).Error

	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("修改广告失败", c)
		return
	}

	res.OkWithMessage("修改广告成功", c)
}
