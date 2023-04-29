package article_api

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/olivere/elastic/v7"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"
)

func (ArticleApi) ArticleCalendarView(c *gin.Context) {

	// 按时间聚合
	agg := elastic.NewDateHistogramAggregation().Field("created_at").CalendarInterval("hour")

	result, err := global.ESClient.
		Search(models.ArticleModel{}.Index()).
		Query(elastic.NewBoolQuery()).
		Aggregation("calendar", agg).
		Size(0).
		Do(context.Background())
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("查询失败", c)
		return
	}

	fmt.Println(string(result.Aggregations["calendar"]))

}
