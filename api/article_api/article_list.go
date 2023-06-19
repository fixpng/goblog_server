package article_api

import (
	"github.com/gin-gonic/gin"
	"github.com/liu-cn/json-filter/filter"
	"github.com/olivere/elastic/v7"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"
	"gvb_server/service/es_ser"
	"gvb_server/service/redis_ser"
	"gvb_server/utils/jwts"
)

type ArticleSearchRequest struct {
	models.PageInfo
	Tag    string `json:"tag" form:"tag"`
	IsUser bool   `json:"is_user" form:"is_user"` // 根据这个参数判断是否显示我收藏的文章列表
}

// ArticleListView 文章列表
// @Tags 文章管理
// @Summary 文章列表
// @Description 文章列表
// @Param data query ArticleSearchRequest    false  "查询参数"
// @Param token header string false "token"
// @Router /api/articles [get]
// @Produce json
// @Success 200 {object} res.Response{data=res.ListResponse[models.ArticleModel]}
func (ArticleApi) ArticleListView(c *gin.Context) {
	var cr ArticleSearchRequest
	if err := c.ShouldBindQuery(&cr); err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	// 列表查询
	boolSearch := elastic.NewBoolQuery()

	// 带了token
	if cr.IsUser {
		token := c.GetHeader("token")
		claims, err := jwts.ParseToken(token)

		if err == nil && !redis_ser.CheckLogout(token) {
			boolSearch.Must(elastic.NewTermsQuery("user_id", claims.UserID))
		}
	}

	list, count, err := es_ser.CommList(
		es_ser.Option{
			PageInfo: cr.PageInfo,
			Fields:   []string{"title", "content", "category"},
			Tag:      cr.Tag,
			Query:    boolSearch,
		})
	if err != nil {
		global.Log.Error(err.Error())
		res.FailWithMessage("查询失败", c)
		return
	}

	// 判断是否为空 json-filter空值问题
	data := filter.Omit("list", list)
	_list, _ := data.(filter.Filter)
	if string(_list.MustMarshalJSON()) == "{}" {
		list = make([]models.ArticleModel, 0)
		res.OkWithList(list, int64(count), c)
		return
	}

	// filter.Omit() 第三方库过滤content字段 "github.com/liu-cn/json-filter/filter"
	res.OkWithList(data, int64(count), c)
}
