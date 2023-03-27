package models

import "time"

type MODEL struct {
	ID        uint      `gorm:"primarykey" json:"id"  structs:"-"`
	CreatedAt time.Time `json:"created_at"  structs:"-"`
	UpdatedAt time.Time `json:"-"  structs:"-"`
}

type RemoveRequest struct {
	IDList []uint `json:"id_list" `
}

type PageInfo struct {
	Page  int    `form:"page"`
	Key   string `form:"key"`
	Limit int    `form:"limit"`
	Sort  string `form:"sort"`
}

var (
	ModelCreate = MODEL{CreatedAt: time.Now(), UpdatedAt: time.Now()}
	ModelUpdate = MODEL{UpdatedAt: time.Now()}
)
