package model

import (
	"github.com/jinzhu/gorm"
)

// Transaction struct
type Transaction struct {
	gorm.Model
	Committer      string `gorm:"not null" json:"committer"`
	CompanyID      uint   `gorm:"not null" json:"companyID"`
	Company        Company
	PlantingAreaID uint `gorm:"not null" json:"plantingAreaID"`
	PlantingArea   PlantingArea
	NumberOfTrees  uint      `gorm:"not null" json:"numberOfTrees"`
}
