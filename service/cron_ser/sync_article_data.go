package cron_ser

import (
	"context"
	"encoding/json"
	"github.com/olivere/elastic/v7"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/service/redis_ser"
)

// SyncArticleData 同步redis文章数据到es
func SyncArticleData() {
	// 1.查询es中的全部数据，为后面数据更新做准备
	result, err := global.ESClient.
		Search(models.ArticleModel{}.Index()).
		Query(elastic.NewMatchAllQuery()).
		Size(10000).
		Do(context.Background())
	if err != nil {
		global.Log.Error(err)
		return
	}
	// 2.拿到redis中的缓存数据
	diggInfo := redis_ser.NewDigg().GetInfo()
	lookInfo := redis_ser.NewArticleLook().GetInfo()
	commentInfo := redis_ser.NewCommentCount().GetInfo()

	for _, hit := range result.Hits.Hits {
		var article models.ArticleModel
		err = json.Unmarshal(hit.Source, &article)
		if err != nil {
			global.Log.Error(err)
			continue
		}
		digg := diggInfo[hit.Id]
		look := lookInfo[hit.Id]
		comment := commentInfo[hit.Id]
		// 3.计算新的数据 旧数据+缓存中数据
		newDigg := article.DiggCount + digg
		newLook := article.LookCount + look
		newComment := article.CommentCount + comment

		// 4.判断数据量是否有变化，如果三个变化之后的数据和之前一样
		if digg == 0 && look == 0 && comment == 0 {
			//global.Log.Infof("%s 无变化", article.Title)
			continue
		}

		// 5.更新
		_, err := global.ESClient.Update().
			Index(models.ArticleModel{}.Index()).
			Id(hit.Id).
			Doc(map[string]int{
				"digg_count":    newDigg,
				"look_count":    newLook,
				"comment_count": newComment,
			}).Do(context.Background())
		if err != nil {
			global.Log.Error(err)
			continue
		}
		global.Log.Infof("%s 更新成功 点赞数：%d,评论数：%d,浏览量：%d",
			article.Title, newDigg, newComment, newLook)
	}

	// 6.清除redis中的数据
	redis_ser.NewDigg().Clear()
	redis_ser.NewArticleLook().Clear()
	redis_ser.NewCommentCount().Clear()

}
