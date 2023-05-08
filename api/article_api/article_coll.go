package article_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"
	"gvb_server/service/es_ser"
	"gvb_server/utils/jwts"
)

// ArticleCollCreateView 收藏/取消收藏文章
// @Tags 文章管理
// @Summary 收藏/取消收藏文章
// @Description 收藏/取消收藏文章
// @Param data body models.ESIDRequest    true  "表示多个参数"
// @Router /api/articles/collects [post]
// @Produce json
// @Success 200 {object} res.Response{}
func (ArticleApi) ArticleCollCreateView(c *gin.Context) {
	var cr models.ESIDRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	// 登录的用户
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwts.CustomClaims)

	model, err := es_ser.CommeDetail(cr.ID)
	if err != nil {
		res.FailWithMessage("文章不存在", c)
		return
	}

	// 文章收藏数
	var num = -1
	var coll models.UserCollectModel
	err = global.DB.Take(&coll, "user_id = ? and article_id = ?", claims.UserID, cr.ID).Error

	if err != nil {
		// 没有找到记录，则收藏，文章收藏数 +1
		global.DB.Create(&models.UserCollectModel{
			UserID:    claims.UserID,
			ArticleID: cr.ID,
		})
		num = 1
	} else {
		// 找到记录，则取消收藏，文章收藏数 -1
		global.DB.Delete(&coll)
	}

	// 更新文章收藏数
	err = es_ser.ArticleUpdate(cr.ID, map[string]any{
		"collects_count": model.CollectsCount + num,
	})
	if num == 1 {
		res.OkWithMessage("收藏文章成功", c)
	} else {
		res.OkWithMessage("取消收藏成功", c)
	}

}
