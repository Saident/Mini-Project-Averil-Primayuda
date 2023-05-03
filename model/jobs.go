package model

import (
	"github.com/jinzhu/gorm"
)

type Jobs struct {
	gorm.Model
	ID        int
	Deskripsi string `json:"nama" form:"nama"`
	Alamat    string `json:"alamat" form:"alamat"`
	Expire    string `json:"expire" form:"expire"`
	Status    string `json:"status" form:"status"`
	Gaji      string `json:"gaji" form:"gaji"`

	Perusahaan_ID int
	Perusahaan    Perusahaan `gorm:"foreignKey:Perusahaan_ID"`
}
