package tag_api

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

type TagResponse struct {
	Label string `json:"label"`
	Value string `json:"value"`
}

// TagNameListView 标签名称列表
// @Tags 标签管理
// @Summary 标签名称列表
// @Description 标签名称列表
// @Router /api/tag_names [get]
// @Produce json
// @Success 200 {object} res.Response{data=[]TagResponse}
func (TagApi) TagNameListView(c *gin.Context) {
	type T struct {
		DocCountErrorUpperBound int `json:"doc_count_error_upper_bound"`
		SumOtherDocCount        int `json:"sum_other_doc_count"`
		Buckets                 []struct {
			Key      string `json:"key"`
			DocCount int    `json:"doc_count"`
		} `json:"buckets"`
	}

	query := elastic.NewBoolQuery()
	agg := elastic.NewTermsAggregation().Field("tags")
	result, err := global.ESClient.
		Search(models.ArticleModel{}.Index()).
		Query(query).
		Aggregation("tags", agg).
		Size(0).
		Do(context.Background())
	if err != nil {
		logrus.Error(err)
		return
	}
	byteData := result.Aggregations["tags"]
	var tagType T
	json.Unmarshal(byteData, &tagType)

	var tagList = make([]TagResponse, 0)
	for _, bucket := range tagType.Buckets {
		tagList = append(tagList, TagResponse{
			Label: bucket.Key,
			Value: bucket.Key,
		})
	}
	res.OkWithData(tagList, c)

}
