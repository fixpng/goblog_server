package tag_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"
)

type TagRequest struct {
	Title string `json:"title" binding:"required" msg:"请输入标题" structs:"title"` // 显示的标题
}

// TagCreateView 创建标签
// @Tags 标签管理
// @Summary 创建标签
// @Description 创建标签
// @Param data body TagRequest    true  "表示多个参数"
// @Router /api/Tags [post]
// @Produce json
// @Success 200 {object} res.Response{}
func (TagApi) TagCreateView(c *gin.Context) {
	var cr TagRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}
	// 重复的判断
	var tag models.TagModel
	err = global.DB.Take(&tag, "title = ?", cr.Title).Error
	if err == nil {
		res.FailWithMessage("该标签已存在", c)
		return
	}

	err = global.DB.Create(&models.TagModel{
		Title: cr.Title,
	}).Error

	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("添加标签失败", c)
		return
	}

	res.OkWithMessage("添加标签成功", c)
}
