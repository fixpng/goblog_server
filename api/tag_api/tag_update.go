package tag_api

import (
	"fmt"
	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"
)

// TagUpdateView 更新标签
// @Tags 标签管理
// @Summary 更新标签
// @Description 更新标签
// @Param data body TagRequest    true  "标签的一些参数"
// @Param token header string    true  "token"
// @Router /api/tags/{id} [put]
// @Produce json
// @Success 200 {object} res.Response{data=string}
func (TagApi) TagUpdateView(c *gin.Context) {

	id := c.Param("id")
	var cr TagRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}
	// 存在的判断
	var tag models.TagModel
	err = global.DB.Take(&tag, id).Error
	if err != nil {
		res.FailWithMessage("标签不存在", c)
		return
	}

	// 结构体转map的第三方包
	maps := structs.Map(&cr)
	fmt.Println(maps)
	err = global.DB.Model(&tag).Updates(maps).Error

	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("修改标签失败", c)
		return
	}

	res.OkWithMessage("修改标签成功", c)
}
