package domain

import (
	"gorm.io/gorm"
)

type Profile struct {
	gorm.Model
	UserID          uint   `gorm:"not null;unique;foreignKey:ID;references:ID"`
	Description     string `gorm:"size:500"`
	Skills          string `gorm:"size:255"`
	Portfolio       string `gorm:"size:255"`
	Specializations string `gorm:"size:255"`
	Availability    bool   `gorm:"size:255"`
	Languages       string `gorm:"size:100"`
	Location        string `gorm:"size:255"`
}
