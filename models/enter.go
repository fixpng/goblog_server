package models

import "time"

type MODEL struct {
	ID        uint      `gorm:"primarykey" json:"id,select($any)"  structs:"-"`
	CreatedAt time.Time `json:"created_at,select($any)"  structs:"-"`
	UpdatedAt time.Time `json:"-"  structs:"-"`
}

type RemoveRequest struct {
	IDList []uint `json:"id_list" `
}

type ESIDRequest struct {
	ID string `json:"id" form:"id" uri:"id"`
}

type ESIDListRequest struct {
	IDList []string `json:"id_list" binding:"required"`
}

type PageInfo struct {
	Page  int    `form:"page"`
	Key   string `form:"key"`
	Limit int    `form:"limit"`
	Sort  string `form:"sort"`
}
