package model

import (
	"github.com/jinzhu/gorm"
)

type Lampiran struct {
	gorm.Model
	ID               int    `gorm:"primaryKey"`
	Lampiran_tipe    int    `json:"Lampiran_tipe" form:"Lampiran_tipe"`
	Lampiran_content string `json:"Lampiran_content" form:"Lampiran_content"`

	UserID int
	User   User `gorm:"references:id"`
}
