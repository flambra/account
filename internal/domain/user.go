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
	FirstName      string
	LastName       string
	TaxNumber      string
	Email          string
	HashedPassword string
	Phone          string
	Address        string
	UserType       string
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
	ID        uint   `json:"id"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Address   string `json:"address"`
	UserType  string `json:"usertype"`
}

type UserPageFilter struct {
	Name     string `json:"name"`
	Address  string `json:"address"`
	UserType string `json:"usertype"`
}

func (f *UserPageFilter) Apply(db *gorm.DB) *gorm.DB {
	return db.Where("first_name LIKE ? OR last_name LIKE ? OR address LIKE ? OR user_type = ?", "%"+f.Name+"%", "%"+f.Name+"%", "%"+f.Address+"%", f.UserType)
}
