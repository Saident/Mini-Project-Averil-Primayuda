package model

import (
	"github.com/jinzhu/gorm"
)

type Lampiran struct {
	gorm.Model
	ID               int
	Lampiran_tipe    int `json:"Lampiran_tipe" form:"Lampiran_tipe"`
	Lampiran_content int `json:"Lampiran_content" form:"Lampiran_content"`

	User_ID int
	User    User `gorm:"foreignKey:User_ID"`
}
