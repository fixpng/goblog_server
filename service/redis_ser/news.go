package redis_ser

import (
	"encoding/json"
	"fmt"
	"gvb_server/global"
	"time"
)

const newsIndex = "news_index"

type NewData struct {
	Index    string `json:"index"`
	Title    string `json:"title"`
	HotValue string `json:"hotValue"`
	Link     string `json:"link"`
}

// SetNews 新闻存入缓存
func SetNews(key string, newData []NewData) error {
	byteData, _ := json.Marshal(newData)
	// 缓存一小时过期
	err := global.Redis.Set(fmt.Sprintf("%s_%s", newsIndex, key), byteData, 1*time.Hour).Err()
	//err := global.Redis.HSet(newsIndex, key, byteData).Err()
	return err
}

func GetNews(key string) (newData []NewData, err error) {
	res := global.Redis.Get(fmt.Sprintf("%s_%s", newsIndex, key)).Val()
	err = json.Unmarshal([]byte(res), &newData)
	return
}
