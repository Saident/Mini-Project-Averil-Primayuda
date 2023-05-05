package model

import (
	"github.com/jinzhu/gorm"
)

type Lamaran struct {
	gorm.Model
	ID             int    `gorm:"primaryKey"`
	Lamaran_status string `json:"lamaran_status" form:"lamaran_status"`

	User_ID int
	User    User `gorm:"references:id"`

	Job_ID int
	Jobs   Jobs `gorm:"references:id"`

	Perusahaan_ID int
	Perusahaan    Perusahaan `gorm:"references:id"`
}
