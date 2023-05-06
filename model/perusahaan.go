package model

import (
	"github.com/jinzhu/gorm"
)

type Perusahaan struct {
	gorm.Model
	Nama     string `json:"nama" form:"nama"`
	Sektor   string `json:"Sektor" form:"Sektor"`
	Alamat   string `json:"alamat" form:"alamat"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}
