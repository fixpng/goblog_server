package redis_ser

import (
	"gvb_server/global"
	"strconv"
)

const diggPrefix = "digg"

// Digg 用户点赞某篇文章
func Digg(id string) error {
	num, _ := global.Redis.HGet(diggPrefix, id).Int()
	num++
	err := global.Redis.HSet(diggPrefix, id, num).Err()
	return err
}

// GetDigg 获取某一篇文章下的点赞数
func GetDigg(id string) int {
	num, _ := global.Redis.HGet(diggPrefix, id).Int()
	return num
}

// GetDiggInfo 获取点赞数据
func GetDiggInfo() map[string]int {
	var diggInfo = map[string]int{}
	maps := global.Redis.HGetAll(diggPrefix).Val()
	for id, val := range maps {
		num, _ := strconv.Atoi(val)
		diggInfo[id] = num
	}
	return diggInfo
}

func DiggClear() {
	global.Redis.Del(diggPrefix)
}
