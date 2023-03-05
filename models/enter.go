package models

import "time"

type MODEL struct {
	ID       uint      `gorm:"primarykey" json:"id"`
	CreateAt time.Time `json:"create_at"`
	UpdateAt time.Time `json:"-"`
}
