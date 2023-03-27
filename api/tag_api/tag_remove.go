package tag_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"
)

// TagRemoveView 批量删除标签
// @Tags 标签管理
// @Summary 批量删除标签
// @Description 批量删除标签
// @Param data body models.RemoveRequest    true  "标签id列表"
// @Router /api/Tags [delete]
// @Produce json
// @Success 200 {object} res.Response{data=string}
func (TagApi) TagRemoveView(c *gin.Context) {
	var cr models.RemoveRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	var tagList []models.TagModel
	count := global.DB.Where("id in (?)", cr.IDList).Find(&tagList).RowsAffected
	if count == 0 {
		res.FailWithMessage("标签不存在", c)
		return
	}
	// 如果这个标签下有文章，怎么办？
	global.DB.Delete(&tagList)
	res.OkWithMessage(fmt.Sprintf("共删除 %d 个标签", count), c)
}
