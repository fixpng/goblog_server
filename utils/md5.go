package utils

import (
	"crypto/md5"
	"fmt"
)

func MD5(data []byte) string {
	has := md5.Sum(data)             //切片
	md5str := fmt.Sprintf("%x", has) //将[]byte转成16进制
	return md5str
}
