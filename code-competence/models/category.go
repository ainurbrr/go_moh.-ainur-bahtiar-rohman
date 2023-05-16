package models

import (
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	ID   uint   `json:"id"`
	Name string `json:"name" form:"name"`
}
