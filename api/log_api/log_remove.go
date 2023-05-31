package log_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"
	"gvb_server/plugins/log_stash"
)

// LogRemoveListView 批量删除日志
// @Tags 日志管理
// @Summary 批量删除日志
// @Description 批量删除日志
// @Param data body models.RemoveRequest    true  "日志id列表"
// @Param token header string true "token"
// @Router /api/logs [delete]
// @Produce json
// @Success 200 {object} res.Response{data=string}
func (LogApi) LogRemoveListView(c *gin.Context) {
	var cr models.RemoveRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	var List []log_stash.LogStashModel
	count := global.DB.Where("id in (?)", cr.IDList).Find(&List).RowsAffected
	if count == 0 {
		res.FailWithMessage("日志不存在", c)
		return
	}
	global.DB.Delete(&List)
	res.OkWithMessage(fmt.Sprintf("共删除 %d 个日志", count), c)
}
