package message_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/models"
	"gvb_server/models/res"
	"gvb_server/service/common"
)

// MessageListAllView 消息列表
// @Tags 消息管理
// @Summary 消息列表
// @Description 消息列表
// @Router /api/messages_all [get]
// @Produce json
// @Success 200 {object} res.Response{data=res.ListResponse[MessageGroup]}
func (MessageApi) MessageListAllView(c *gin.Context) {
	var cr models.PageInfo
	if err := c.ShouldBindQuery(&cr); err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	list, count, _ := common.ComList(models.MessageModel{}, common.Option{
		PageInfo: cr,
	})
	// 需要展示这个标签下文章的数量
	res.OkWithList(list, count, c)
}
