package main

import (
	"fmt"
	"github.com/fatih/structs"
	"gvb_server/models"
)

type AdvertRequest struct {
	models.MODEL `structs:"-"`
	Title        string `json:"title" binding:"required" msg:"请输入标题" structs:"title"`       // 显示的标题
	Href         string `json:"href" binding:"required,url" msg:"跳转链接非法" structs:"-"`       // 跳转链接
	Images       string `json:"images" binding:"required,url" msg:"图片地址非法"`                 // 图片
	IsShow       bool   `json:"is_show" binding:"required" msg:"请选择是否展示" structs:"is_show"` // 是否展示
}

func main() {
	u1 := AdvertRequest{
		Title:  "xxxxxxx",
		Href:   "xxxxxxx",
		Images: "xxxxxxx",
		IsShow: true,
	}
	m3 := structs.Map(&u1)
	fmt.Println(m3)
}
