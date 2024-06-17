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

type UserCreateRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	TaxNumber string `json:"tax_number"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Phone     string `json:"phone"`
	Address   string `json:"address"`
	UserType  string `json:"user_type"`
}

type UserUpdateRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	TaxNumber string `json:"tax_number"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Phone     string `json:"phone"`
	Address   string `json:"address"`
	UserType  string `json:"user_type"`
}

type UserPageResponse struct {
	Name     string `json:"name"`
	Address  string `json:"address"`
	UserType string `json:"user_type"`
}

type UserPageFilter struct {
	Name     string `json:"name"`
	Address  string `json:"address"`
	UserType string `json:"user_type"`
}

func (f *UserPageFilter) Apply(db *gorm.DB) *gorm.DB {
	return db.Where("first_name LIKE ? OR last_name LIKE ? OR address LIKE ? OR user_type = ?", "%"+f.Name+"%", "%"+f.Name+"%", "%"+f.Address+"%", "%"+f.UserType+"%")
}
