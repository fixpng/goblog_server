package comment_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"
	"gvb_server/service/redis_ser"
)

type CommentIDRequest struct {
	ID uint `json:"id" form:"id" uri:"id"`
}

// CommentDigg 评论点赞
// @Tags 评论管理
// @Summary 评论点赞
// @Description 评论点赞
// @Param data body CommentIDRequest    true  "表示多个参数"
// @Param id path int true "文章id"
// @Router /api/comments/{id} [post]
// @Produce json
// @Success 200 {object} res.Response{}
func (CommentApi) CommentDigg(c *gin.Context) {
	var cr CommentIDRequest
	err := c.ShouldBindUri(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	// 评论是否存在
	var commentModel models.CommentModel
	err = global.DB.Take(&commentModel, cr.ID).Error
	if err != nil {
		res.FailWithMessage("评论不存在", c)
		return
	}

	redis_ser.NewCommentDigg().Set(fmt.Sprintf("%d", cr.ID))
	res.OkWithMessage("评论点赞成功", c)
}
