package model

import (
	"github.com/jinzhu/gorm"
)

type Jobs struct {
	gorm.Model
	Deskripsi string `json:"deskripsi" form:"deskripsi"`
	Alamat    string `json:"alamat" form:"alamat"`
	Expire    string `json:"expire" form:"expire"`
	Status    string `json:"status" form:"status"`
	Gaji      int    `json:"gaji" form:"gaji"`

	PerusahaanID int
}
