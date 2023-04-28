package article_api

import (
	"github.com/gin-gonic/gin"
	"github.com/liu-cn/json-filter/filter"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"
	"gvb_server/service/es_ser"
)

func (ArticleApi) ArticleListView(c *gin.Context) {
	var cr models.PageInfo
	if err := c.ShouldBindQuery(&cr); err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	list, count, err := es_ser.CommList(cr.Key, cr.Page, cr.Limit)
	if err != nil {
		global.Log.Error(err.Error())
		res.FailWithMessage("查询失败", c)
		return
	}
	// filter.Omit() 第三方库过滤content字段 "github.com/liu-cn/json-filter/filter"
	res.OkWithList(filter.Omit("list", list), int64(count), c)
}
