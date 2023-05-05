package model

import (
	"github.com/jinzhu/gorm"
)

type Lamaran struct {
	gorm.Model
	ID             int    `gorm:"primaryKey"`
	Lamaran_status string `json:"lamaran_status" form:"lamaran_status"`

	UserID int
	User   User `gorm:"references:id"`

	JobID int
	Jobs  Jobs `gorm:"references:id"`

	PerusahaanID int
	Perusahaan   Perusahaan `gorm:"references:id"`
}
