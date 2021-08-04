package model

import (
"github.com/jinzhu/gorm"
)

// Balance struct
type Balance struct {
	gorm.Model
	Amount      uint `gorm:"not null" json:"amount"`
	UserID      uint `gorm:"not null" json:"user_id"`
	User        User
}
