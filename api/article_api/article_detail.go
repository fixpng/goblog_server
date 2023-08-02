package article_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"
	"gvb_server/service/es_ser"
	"gvb_server/service/redis_ser"
	"gvb_server/utils/jwts"
)

type ArticleDetailResponse struct {
	models.ArticleModel
	IsCollect bool `json:"is_collect"` // 用户是否收藏文章
}

// ArticleDetailView 文章详情
// @Tags 文章管理
// @Summary 文章详情
// @Description 文章详情
// @Param id path string true "id"
// @Router /api/articles/{id} [get]
// @Produce json
// @Success 200 {object} res.Response{data=res.ListResponse[models.ArticleModel]}
func (ArticleApi) ArticleDetailView(c *gin.Context) {
	var cr models.ESIDRequest
	err := c.ShouldBindUri(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	// 用户浏览量
	redis_ser.NewArticleLook().Set(cr.ID)

	model, err := es_ser.CommeDetail(cr.ID)
	if err != nil {
		res.FailWithMessage(err.Error(), c)
		return
	}
	isCollect := IsUserArticleColl(c, model.ID)

	var articleDetail = ArticleDetailResponse{
		ArticleModel: model,
		IsCollect:    isCollect,
	}

	res.OkWithData(articleDetail, c)
}

func IsUserArticleColl(c *gin.Context, articleID string) (isCollect bool) {
	// 查询用户是否正常登录
	token := c.GetHeader("token")
	if token == "" {
		return
	}
	claims, err := jwts.ParseToken(token)
	if err != nil {
		return
	}
	// 判断是否在redis中
	if redis_ser.CheckLogout(token) {
		return
	}
	var count int64
	global.DB.Model(models.UserCollectModel{}).Where("user_id = ? and article_id = ?", claims.UserID, articleID).Count(&count)
	if count == 0 {
		return
	}
	return true

}

type ArticleDetailRequest struct {
	Title string `json:"title" form:"title"`
}

// ArticleDetailByTitleView 文章标题查内容
// @Tags 文章管理
// @Summary 文章标题查内容
// @Description 文章标题查内容
// @Param data query ArticleDetailRequest    true  "文章标题"
// @Router /api/articles/detail [get]
// @Produce json
// @Success 200 {object} res.Response{data=res.ListResponse[models.ArticleModel]}
func (ArticleApi) ArticleDetailByTitleView(c *gin.Context) {
	var cr ArticleDetailRequest
	err := c.ShouldBindQuery(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	model, err := es_ser.CommeDetailByKeyword(cr.Title)
	if err != nil {
		res.FailWithMessage(err.Error(), c)
		return
	}
	res.OkWithData(model, c)
}
