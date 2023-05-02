package flag

import "gvb_server/models"

func EsCreateIndex() {
	//models.ArticleModel{}.CreateIndex()
	models.FullTextModel{}.CreateIndex()
}
