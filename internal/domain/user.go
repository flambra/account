package domain

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID             uint `gorm:"primarykey"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      gorm.DeletedAt `gorm:"index"`
	FirstName      string         `gorm:"size:100;not null"`
	LastName       string         `gorm:"size:100;not null"`
	TaxNumber      string         `gorm:"size:100;unique;not null"`
	Email          string         `gorm:"size:100;unique;not null"`
	HashedPassword string         `gorm:"size:255;not null"`
	Phone          string         `gorm:"size:15;unique"`
	Address        string         `gorm:"size:255;not null"`
	UserType       string         `gorm:"size:50;not null"`
	LastCode       string
	Profile        Profile `gorm:"foreignKey:UserID"`
}
