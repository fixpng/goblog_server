package es_ser

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"
	"gvb_server/global"
	"gvb_server/models"
)

type Option struct {
}

func CommList(key string, page, limit int) (list []models.ArticleModel, count int, err error) {
	// 列表查询
	boolSearch := elastic.NewBoolQuery()
	from := page
	if key != "" {
		boolSearch.Must(
			//elastic.NewPrefixQuery("title", key),
			elastic.NewMultiMatchQuery(key, "title", "abstract", "content"),
		)
	}
	// 默认值
	if limit == 0 {
		limit = 10
	}
	if from == 0 {
		from = 1
	}

	res, err := global.ESClient.
		Search(models.ArticleModel{}.Index()).
		Query(boolSearch).
		Highlight(elastic.NewHighlight().Field("title")).
		From((from - 1) * limit).
		Size(limit).
		Do(context.Background())
	if err != nil {
		logrus.Error(err.Error())
		return
	}

	count = int(res.Hits.TotalHits.Value) //搜索到的结果总条数
	demoList := []models.ArticleModel{}
	for _, hit := range res.Hits.Hits {
		var model models.ArticleModel
		data, err := hit.Source.MarshalJSON()
		if err != nil {
			logrus.Error(err.Error())
			continue
		}
		err = json.Unmarshal(data, &model)
		if err != nil {
			logrus.Error(err)
			continue
		}
		title, ok := hit.Highlight["title"]
		if ok {
			model.Title = title[0]
		}
		model.ID = hit.Id
		demoList = append(demoList, model)
	}
	return demoList, count, err

}

// CommeDetail 根据id查
func CommeDetail(id string) (model models.ArticleModel, err error) {
	res, err := global.ESClient.
		Get().
		Index(models.ArticleModel{}.Index()).
		Id(id).
		Do(context.Background())
	if err != nil {
		logrus.Error(err.Error())
		return
	}
	err = json.Unmarshal(res.Source, &model)
	if err != nil {
		return
	}
	model.ID = res.Id
	return
}

// CommeDetailByKeyword 根据keyword查
func CommeDetailByKeyword(key string) (model models.ArticleModel, err error) {
	res, err := global.ESClient.Search().
		Index(models.ArticleModel{}.Index()).
		Query(elastic.NewTermQuery("keyword", key)).
		Size(1).
		Do(context.Background())
	if err != nil {
		logrus.Error(err.Error())
		return
	}
	if res.Hits.TotalHits.Value == 0 {
		return model, errors.New("文章不存在")
	}
	hit := res.Hits.Hits[0]
	err = json.Unmarshal(hit.Source, &model)
	if err != nil {
		return
	}
	model.ID = hit.Id
	return
}
