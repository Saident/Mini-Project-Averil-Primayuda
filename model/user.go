package model

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	ID          int    `gorm:"primaryKey"`
	Nama        string `json:"name" form:"name"`
	Tgl_lahir   string `json:"tgl_lahir" form:"tgl_lahir"`
	Alamat      string `json:"alamat" form:"alamat"`
	Disabilitas string `json:"disabilitas" form:"disabilitas"`
	Kelamin     string `json:"kelamin" form:"kelamin"`
	Email       string `json:"email" form:"email"`
	Password    string `json:"password" form:"password"`
}
