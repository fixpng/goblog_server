package article_api

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/olivere/elastic/v7"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"
	"gvb_server/service/es_ser"
)

type IDListRequest struct {
	IDList []string `json:"id_list"`
}

// ArticleRemoveView 批量删除文章
// @Tags 文章管理
// @Summary 批量删除文章
// @Description 批量删除文章
// @Param data body IDListRequest    true  "文章id列表"
// @Router /api/articles [delete]
// @Param token header string false "token"
// @Produce json
// @Success 200 {object} res.Response{data=string}
func (ArticleApi) ArticleRemoveView(c *gin.Context) {
	var cr IDListRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		global.Log.Error(err)
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	// 文章删除后，用户收藏过这篇文章如何处理?（代码待完善）
	// 1.删除时把文章关联的收藏也删除
	// 2.用户收藏表，新增一个字段，表示文章是否删除，用户可以删除这个收藏记录，但是找不到文章改收藏数
	bulkService := global.ESClient.Bulk().Index(models.ArticleModel{}.Index()).Refresh("true")
	for _, id := range cr.IDList {
		req := elastic.NewBulkDeleteRequest().Id(id)
		bulkService.Add(req)
		// 异步全文搜索删除
		go es_ser.DeleteFullTextByArticleID(id)
	}
	result, err := bulkService.Do(context.Background())
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("删除失败", c)
		return
	}
	res.OkWithMessage(fmt.Sprintf("成功删除 %d 篇文章", len(result.Succeeded())), c)
}
