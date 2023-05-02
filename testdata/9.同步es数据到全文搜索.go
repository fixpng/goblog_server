package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"
	"gvb_server/core"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/service/es_ser"
)

func main() {
	core.InitConf()
	core.InitConf()
	global.ESClient = core.EsConnect()

	boolSearch := elastic.NewMatchAllQuery()
	result, _ := global.ESClient.Search(models.ArticleModel{}.Index()).
		Query(boolSearch).
		Size(1000).
		Do(context.Background())

	for _, hit := range result.Hits.Hits {
		var article models.ArticleModel
		_ = json.Unmarshal(hit.Source, &article)

		indexList := es_ser.GetSearchIndexDataByContent(hit.Id, article.Title, article.Content)

		// 批量添加
		bulk := global.ESClient.Bulk()
		for _, indexData := range indexList {
			req := elastic.NewBulkIndexRequest().Index(models.FullTextModel{}.Index()).Doc(indexData)
			bulk.Add(req)
		}
		result, err := bulk.Do(context.Background())
		if err != nil {
			fmt.Println(err)
			logrus.Error(err)
			continue
		}
		fmt.Println(article.Title, "添加成功", len(result.Succeeded()), "条")
	}
}
