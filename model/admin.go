package model

import (
	"github.com/jinzhu/gorm"
)

type Admin struct {
	gorm.Model
	ID       int
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}
