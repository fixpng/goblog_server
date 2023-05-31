package log_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/models"
	"gvb_server/models/res"
	"gvb_server/plugins/log_stash"
	"gvb_server/service/common"
)

type LogRequest struct {
	models.PageInfo
	Level log_stash.Level `form:"level"`
}

// LogListView 日志列表
// @Tags 日志管理
// @Summary 日志列表
// @Description 日志列表
// @Param data query LogRequest    false  "查询参数"
// @Param level query int false      "日志等级"
// @Router /api/logs [get]
// @Produce json
// @Success 200 {object} res.Response{data=res.ListResponse[log_stash.LogStashModel]}
func (LogApi) LogListView(c *gin.Context) {
	var cr LogRequest
	err := c.ShouldBindQuery(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	if cr.Sort == "" {
		cr.Sort = "created_at desc"
	}

	list, count, _ := common.ComList(log_stash.LogStashModel{Level: cr.Level}, common.Option{
		PageInfo: cr.PageInfo,
		Debug:    false,
		Likes:    []string{"ip", "addr"},
	})
	res.OkWithList(list, count, c)
	return

}
