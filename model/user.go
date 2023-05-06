package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Nama        string    `json:"nama" form:"nama"`
	Tgl_lahir   time.Time `json:"tgl_lahir" form:"tgl_lahir"`
	Alamat      string    `json:"alamat" form:"alamat"`
	Disabilitas string    `json:"disabilitas" form:"disabilitas"`
	Kelamin     string    `json:"kelamin" form:"kelamin"`
	Email       string    `json:"email" form:"email"`
	Password    string    `json:"password" form:"password"`
}
