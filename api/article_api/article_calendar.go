package article_api

import (
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/olivere/elastic/v7"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"
	"time"
)

type CalendarDateResponse struct {
	Date  string `json:"date"`
	Count int    `json:"count"`
}

type BucketsType struct {
	Buckets []struct {
		KeyAsString string `json:"key_as_string"`
		Key         int64  `json:"key"`
		DocCount    int    `json:"doc_count"`
	} `json:"buckets"`
}

var DataCount = map[string]int{}

func (ArticleApi) ArticleCalendarView(c *gin.Context) {

	// 按时间聚合
	agg := elastic.NewDateHistogramAggregation().Field("created_at").CalendarInterval("hour")

	// 时间段搜索
	// 从今天开始，到去年的今天
	now := time.Now()
	aYearAgo := now.AddDate(-1, 0, 0)
	format := "2006-01-02 15:04:05"
	// lt 小于 gt 大于
	query := elastic.NewRangeQuery("created_at").
		Gte(aYearAgo.Format(format)).
		Lte(now.Format(format))

	result, err := global.ESClient.
		Search(models.ArticleModel{}.Index()).
		Query(query).
		Aggregation("calendar", agg).
		Size(0).
		Do(context.Background())
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("查询失败", c)
		return
	}

	var data BucketsType
	_ = json.Unmarshal(result.Aggregations["calendar"], &data)

	var resList = make([]CalendarDateResponse, 0)
	for _, bucket := range data.Buckets {
		Time, _ := time.Parse(format, bucket.KeyAsString)
		DataCount[Time.Format("2006-01-02")] = bucket.DocCount

	}

	days := int(now.Sub(aYearAgo).Hours() / 24)
	for i := 0; i < days; i++ {
		day := aYearAgo.AddDate(0, 0, i).Format("2006-01-02")
		count, _ := DataCount[day]
		resList = append(resList, CalendarDateResponse{
			Date:  day,
			Count: count,
		})
	}

	res.OkWithData(resList, c)

}
