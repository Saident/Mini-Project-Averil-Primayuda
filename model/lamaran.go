package model

import (
	"github.com/jinzhu/gorm"
)

type Lamaran struct {
	gorm.Model
	ID             int
	Lamaran_status string `json:"lamaran_status" form:"lamaran_status"`

	User_ID int
	User    User `gorm:"foreignKey:User_ID"`

	Job_ID int
	Jobs   Jobs `gorm:"foreignKey:Perusahaan_ID"`

	Perusahaan_ID int
	Perusahaan    Perusahaan `gorm:"foreignKey:Perusahaan_ID"`
}
