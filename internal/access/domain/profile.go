package domain

import (
	"time"

	"gorm.io/gorm"
)

type Profile struct {
	UserID          uint `gorm:"not null;autoIncrement:false"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       gorm.DeletedAt `gorm:"index"`
	Description     string         `gorm:"size:500"`
	Skills          string         `gorm:"size:255"`
	Portfolio       string         `gorm:"size:255"`
	Specializations string         `gorm:"size:255"`
	Availability    bool
	Languages       string         `gorm:"size:100"`
	Location        string         `gorm:"size:255"`
}


