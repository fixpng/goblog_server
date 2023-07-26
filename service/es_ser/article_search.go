package es_ser

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/service/redis_ser"
	"strings"
)

func CommList(option Option) (list []models.ArticleModel, count int, err error) {
	if option.Key != "" {
		option.Query.Must(
			elastic.NewMultiMatchQuery(option.Key, option.Fields...),
		)
	}
	// 根据标签搜
	if option.Tag != "" {
		option.Query.Must(
			elastic.NewMultiMatchQuery(option.Tag, "tags"),
		)
	}
	// 定义排序
	type SortField struct {
		Field     string
		Ascending bool
	}
	sortField := SortField{
		Field:     "created_at",
		Ascending: false, // true从小到大 false从大到小
	}
	if option.Sort != "" {
		_list := strings.Split(option.Sort, " ")
		if len(_list) == 2 && (_list[1] == "desc" || _list[1] == "asc") {
			sortField.Field = _list[0]
			if _list[1] == "desc" {
				sortField.Ascending = false
			}
			if _list[1] == "asc" {
				sortField.Ascending = true
			}
		}
	}

	res, err := global.ESClient.
		Search(models.ArticleModel{}.Index()).
		Query(option.Query).
		Highlight(elastic.NewHighlight().Field("title")).
		From(option.GetForm()).
		Sort(sortField.Field, sortField.Ascending).
		Size(option.Limit).
		Do(context.Background())
	if err != nil {
		logrus.Error(err.Error())
		return
	}

	count = int(res.Hits.TotalHits.Value) //搜索到的结果总条数
	demoList := []models.ArticleModel{}

	diggInfo := redis_ser.NewDigg().GetInfo()
	lookInfo := redis_ser.NewArticleLook().GetInfo()
	commentInfo := redis_ser.NewCommentCount().GetInfo()
	for _, hit := range res.Hits.Hits {
		var model models.ArticleModel
		err = json.Unmarshal(hit.Source, &model)
		if err != nil {
			logrus.Error(err)
			continue
		}
		title, ok := hit.Highlight["title"]
		if ok {
			model.Title = title[0]
		}
		model.ID = hit.Id
		digg := diggInfo[hit.Id]
		look := lookInfo[hit.Id]
		comment := commentInfo[hit.Id]

		model.DiggCount = model.DiggCount + digg
		model.LookCount = model.LookCount + look
		model.CommentCount = model.CommentCount + comment

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
	model.LookCount = model.LookCount + redis_ser.NewArticleLook().Get(res.Id)
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

// ArticleUpdate 更新点赞或收藏
func ArticleUpdate(id string, data map[string]any) error {
	_, err := global.ESClient.
		Update().
		Index(models.ArticleModel{}.Index()).
		Id(id).
		Doc(data).Refresh("true").
		Do(context.Background())
	return err
}
