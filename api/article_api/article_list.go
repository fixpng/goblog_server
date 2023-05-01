package article_api

import (
	"github.com/gin-gonic/gin"
	"github.com/liu-cn/json-filter/filter"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"
	"gvb_server/service/es_ser"
)

type ArticleSearchRequest struct {
	models.PageInfo
	Tag string `json:"tag" form:"tag"`
}

func (ArticleApi) ArticleListView(c *gin.Context) {
	var cr ArticleSearchRequest
	if err := c.ShouldBindQuery(&cr); err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	list, count, err := es_ser.CommList(
		es_ser.Option{
			PageInfo: cr.PageInfo,
			Fields:   []string{"title", "content"},
			Tag:      cr.Tag,
		})
	if err != nil {
		global.Log.Error(err.Error())
		res.FailWithMessage("查询失败", c)
		return
	}
	data := filter.Omit("list", list)

	// 判断是否为空 json-filter空值问题
	_list, _ := data.(filter.Filter)
	if string(_list.MustMarshalJSON()) == "{}" {
		list = make([]models.ArticleModel, 0)
		res.OkWithList(list, int64(count), c)
		return
	}

	// filter.Omit() 第三方库过滤content字段 "github.com/liu-cn/json-filter/filter"
	res.OkWithList(data, int64(count), c)
}
