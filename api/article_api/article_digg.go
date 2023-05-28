package article_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/models"
	"gvb_server/models/res"
	"gvb_server/service/redis_ser"
)

// ArticleDiggView 文章点赞
// @Tags 文章管理
// @Summary 文章点赞
// @Description 文章点赞
// @Param data body models.ESIDRequest    true  "表示多个参数"
// @Router /api/article/digg [post]
// @Produce json
// @Success 200 {object} res.Response{data=string}
func (ArticleApi) ArticleDiggView(c *gin.Context) {
	var cr models.ESIDRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	// id长度校验及查es文章是否存在（待补充）
	redis_ser.NewDigg().Set(cr.ID)
	res.OkWithMessage("文章点赞成功", c)
}
