package article_api

import (
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"
)

// FullTextContextView 全文搜索
// @Tags 文章管理
// @Summary 全文搜索
// @Description 全文搜索
// @Param data query models.PageInfo    false  "查询参数"
// @Router /api/articles/text [get]
// @Produce json
// @Success 200 {object} res.Response{data=res.ListResponse[models.FullTextModel]}
func (ArticleApi) FullTextContextView(c *gin.Context) {
	var cr models.PageInfo
	_ = c.ShouldBindQuery(&cr)

	boolQuery := elastic.NewBoolQuery()
	if cr.Key != "" {
		boolQuery.Must(elastic.NewMultiMatchQuery(cr.Key, "title", "body"))
	}
	result, err := global.ESClient.
		Search(models.FullTextModel{}.Index()).
		Query(boolQuery).
		Highlight(elastic.NewHighlight().Field("body")).
		Size(100).
		Do(context.Background())
	if err != nil {
		logrus.Error(err.Error())
		return
	}

	count := result.Hits.TotalHits.Value //搜索到的结果总条数
	fullTextList := make([]models.FullTextModel, 0)
	for _, hit := range result.Hits.Hits {
		var model models.FullTextModel
		json.Unmarshal(hit.Source, &model)
		body, ok := hit.Highlight["body"]
		if ok {
			model.Body = body[0]
		}
		fullTextList = append(fullTextList, model)
	}

	res.OkWithList(fullTextList, count, c)

}
