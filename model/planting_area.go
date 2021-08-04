package model

import (
	"github.com/jinzhu/gorm"
)

// PlantingArea struct
type PlantingArea struct {
	gorm.Model
	Name        string `gorm:"not null" json:"name"`
}
