package redis_ser

import (
	"gvb_server/global"
	"strconv"
)

// CountDB redis构造方法
type CountDB struct {
	Index string // 索引前缀
}

// Set 设置某一个数据，重复执行，重复累加
func (c CountDB) Set(id string) error {
	num, _ := global.Redis.HGet(c.Index, id).Int()
	num++
	err := global.Redis.HSet(c.Index, id, num).Err()
	return err
}

// SetCount 在原有基础上增加多少
func (c CountDB) SetCount(id string, num int) error {
	oldNum, _ := global.Redis.HGet(c.Index, id).Int()
	newNum := oldNum + num
	err := global.Redis.HSet(c.Index, id, newNum).Err()
	return err
}

// Get 获取某一数据
func (c CountDB) Get(id string) int {
	num, _ := global.Redis.HGet(c.Index, id).Int()
	return num
}

// GetInfo 获取数据
func (c CountDB) GetInfo() map[string]int {
	var diggInfo = map[string]int{}
	maps := global.Redis.HGetAll(c.Index).Val()
	for id, val := range maps {
		num, _ := strconv.Atoi(val)
		diggInfo[id] = num
	}
	return diggInfo
}

// Clear 删除数据
func (c CountDB) Clear() {
	global.Redis.Del(c.Index)
}
