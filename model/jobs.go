package model

import (
	"github.com/jinzhu/gorm"
)

type Jobs struct {
	gorm.Model
	ID        int    `gorm:"primaryKey"`
	Deskripsi string `json:"deskripsi" form:"deskripsi"`
	Alamat    string `json:"alamat" form:"alamat"`
	Expire    string `json:"expire" form:"expire"`
	Status    string `json:"status" form:"status"`
	Gaji      string `json:"gaji" form:"gaji"`

	PerusahaanID int
	Perusahaan   Perusahaan `gorm:"references:id"`
}
