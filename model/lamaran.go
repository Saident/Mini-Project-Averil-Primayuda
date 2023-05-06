package model

import (
	"github.com/jinzhu/gorm"
)

type Lamaran struct {
	gorm.Model
	Lamaran_status string `json:"lamaran_status" form:"lamaran_status"`

	UserID int

	JobID int

	PerusahaanID int
}
