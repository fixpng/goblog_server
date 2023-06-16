package article_api

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"
	"gvb_server/service/redis_ser"
)

// ArticleContentView 获取文章正文
// @Tags 文章管理
// @Summary 获取文章正文
// @Description 获取文章正文
// @Param id path string true "id"
// @Router /api/articles/content/{id} [get]
// @Produce json
// @Success 200 {object} res.Response{}
func (ArticleApi) ArticleContentView(c *gin.Context) {
	var cr models.ESIDRequest
	err := c.ShouldBindUri(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	// 用户浏览量
	redis_ser.NewArticleLook().Set(cr.ID)

	result, err := global.ESClient.
		Get().
		Index(models.ArticleModel{}.Index()).
		Id(cr.ID).
		Do(context.Background())
	if err != nil {
		res.FailWithMessage("查询失败", c)
		return
	}
	var model models.ArticleModel
	err = json.Unmarshal(result.Source, &model)
	if err != nil {
		return
	}
	res.OkWithData(model.Content, c)
}
