package message_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/models"
	"gvb_server/models/res"
	"gvb_server/service/common"
)

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
