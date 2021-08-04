package model

import (
	"github.com/jinzhu/gorm"
)

// Company struct
type Company struct {
	gorm.Model
	Name        string `gorm:"not null" json:"name"`
}
