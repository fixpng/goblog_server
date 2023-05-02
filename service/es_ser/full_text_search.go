package es_ser

import (
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/russross/blackfriday"
	"strings"
)

type SearchData struct {
	Body  string `json:"body"`  // 正文
	Slug  string `json:"slug"`  // 跳转地址，包含文章的id
	Title string `json:"title"` // 标题
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
			//if strings.TrimSpace(body) != "" {
			bodyList = append(bodyList, getBody(body))
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
		})
	}
	// 遍历次数
	b, _ := json.Marshal(searchDataList)
	fmt.Println(string(b))
	return searchDataList
}

// getHeader 获取搜索文章标题
func getHeader(head string) string {
	head = strings.ReplaceAll(head, "#", "")
	head = strings.ReplaceAll(head, " ", "")
	return "#" + head
}

// getBody 获取搜索文章内容
func getBody(body string) string {
	// 处理content
	unsafe := blackfriday.MarkdownCommon([]byte(body))
	// 是否又script标签
	doc, _ := goquery.NewDocumentFromReader(strings.NewReader(string(unsafe)))
	//fmt.Println(doc.Text())
	return doc.Text()
}

func getSlug(slug string) string {
	return "#" + slug
}
