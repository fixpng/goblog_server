package main

import (
	"context"
	"github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"
	"gvb_server/core"
)

var client *elastic.Client

func EsConnect() *elastic.Client {
	var err error
	sniffOpt := elastic.SetSniff(false)
	host := "http://localhost:9200"
	c, err := elastic.NewClient(
		elastic.SetURL(host),
		sniffOpt,
		elastic.SetBasicAuth("", ""),
	)
	if err != nil {
		logrus.Fatalf("es连接失败 %s", err.Error())
	}
	return c
}

func init() {
	core.InitConf()
	core.InitLogger()
	client = EsConnect()
}

type DemoModel struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	UserID    uint   `json:"user_id"`
	CreatedAt string `json:"created_at"`
}

func (DemoModel) Index() string {
	return "demo_index"
}

func Create(data *DemoModel) (err error) {
	indexResponse, err := client.Index().
		Index(data.Index()).
		BodyJson(data).Do(context.Background())
	if err != nil {
		logrus.Error(err.Error())
		return err
	}
	data.ID = indexResponse.Id
	return nil
}

func main() {
	//DemoModel{}.CreateIndex()
	//Create(&DemoModel{Title: "python基础", UserID: 1, CreatedAt: time.Now().Format("2006-01-02 15:04:05")})
}
