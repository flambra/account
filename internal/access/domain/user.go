package domain

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FirstName      string `gorm:"size:100;not null"`
	LastName       string `gorm:"size:100;not null"`
	TaxNumber      string `gorm:"size:100;unique;not null"`
	Email          string `gorm:"size:100;unique;not null"`
	HashedPassword string `gorm:"size:255;not null"`
	Phone          string `gorm:"size:15;unique"`
	Address        string `gorm:"size:255"`
	UserType       string `gorm:"size:50;not null"` // "freelancer" or "contractor"
	Profile        Profile
}
