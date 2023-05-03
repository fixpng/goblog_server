package es_ser

import (
	"context"
	"github.com/PuerkitoBio/goquery"
	"github.com/olivere/elastic/v7"
	"github.com/russross/blackfriday"
	"github.com/sirupsen/logrus"
	"gvb_server/global"
	"gvb_server/models"
	"strings"
)

type SearchData struct {
	Key   string `json:"key"`   // 文章关联的id
	Body  string `json:"body"`  // 正文
	Slug  string `json:"slug"`  // 跳转地址，包含文章的id
	Title string `json:"title"` // 标题
}

// AsyncArticleByFullText 同步文章数据到全文搜索
func AsyncArticleByFullText(id, title, content string) {
	indexList := GetSearchIndexDataByContent(id, title, content)
	// 批量添加
	bulk := global.ESClient.Bulk()
	for _, indexData := range indexList {
		req := elastic.NewBulkIndexRequest().Index(models.FullTextModel{}.Index()).Doc(indexData)
		bulk.Add(req)
	}
	result, err := bulk.Do(context.Background())
	if err != nil {
		logrus.Error(err)
		return
	}
	logrus.Infof("%s 添加成功，共 %d 条", title, len(result.Succeeded()))
}

// DeleteFullTextByArticleID 删除文章搜索数据
func DeleteFullTextByArticleID(id string) {
	boolSearch := elastic.NewTermQuery("key", id)
	result, _ := global.ESClient.
		DeleteByQuery().
		Index(models.FullTextModel{}.Index()).
		Query(boolSearch).
		Size(1000).
		Do(context.Background())
	logrus.Infof("成功删除 %d 条记录", result.Deleted)
}

// GetSearchIndexDataByContent 全文搜索
func GetSearchIndexDataByContent(id, title, content string) (searchDataList []SearchData) {
	dataList := strings.Split(content, "\n")
	var isCode bool = false
	var headList, bodyList []string
	var body string
	headList = append(headList, getHeader(title))
	for _, s := range dataList {
		// #{1,6}
		// 判断一下是否是代码块
		if strings.HasPrefix(s, "```") {
			isCode = !isCode
		}
		if strings.HasPrefix(s, "#") && !isCode {
			headList = append(headList, getHeader(s))
			//if strings.TrimSpace(s) != "" {
			bodyList = append(bodyList, getBody(s))
			//}
			body = ""
			continue
		}
		body += s
	}
	bodyList = append(bodyList, getBody(body))
	ln := len(headList)
	for i := 0; i < ln; i++ {
		searchDataList = append(searchDataList, SearchData{
			Body:  bodyList[i],
			Title: headList[i],
			Slug:  id + getSlug(headList[i]),
			Key:   id,
		})
	}
	return searchDataList
}

// getHeader 获取搜索文章标题
func getHeader(head string) string {
	head = strings.ReplaceAll(head, "#", "")
	head = strings.ReplaceAll(head, " ", "")
	return head
}

// getBody 获取搜索文章内容
func getBody(body string) string {
	// 处理content
	unsafe := blackfriday.MarkdownCommon([]byte(body))
	// 是否又script标签
	doc, _ := goquery.NewDocumentFromReader(strings.NewReader(string(unsafe)))
	//fmt.Println(doc.Text(), "aaa")
	return doc.Text()
}

func getSlug(slug string) string {
	return "#" + slug
}
