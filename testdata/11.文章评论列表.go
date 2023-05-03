package main

import (
	"fmt"
	"gvb_server/core"
	"gvb_server/global"
	"gvb_server/models"
)

func main() {
	core.InitConf()
	global.Log = core.InitLogger()
	global.DB = core.InitGorm()
	FindArticleCommentList("XNI61ocBvJkxiQvwtge8")
}

func FindArticleCommentList(articleID string) {
	// 先把文章下的根评论查出
	var RootCommentList []*models.CommentModel
	global.DB.Find(&RootCommentList, "article_id = ? and parent_comment_id is null", articleID)
	// 遍历根评论，递归查根评论下的所有子评论
	for _, model := range RootCommentList {
		var subCommentList []models.CommentModel
		FindSubComment(*model, &subCommentList)
		model.SubComments = subCommentList
	}
	fmt.Println(RootCommentList[0])
}

// FindSubComment 递归查某评论下的子评论
func FindSubComment(model models.CommentModel, subCommentList *[]models.CommentModel) {
	global.DB.Preload("SubComments").Take(&model)
	for _, sub := range model.SubComments {
		*subCommentList = append(*subCommentList, sub)
		FindSubComment(sub, subCommentList)
	}
	return
}
