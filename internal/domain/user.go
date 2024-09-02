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
	FirstName      string         `json:"first_name"`
	LastName       string         `json:"last_name"`
	BirthDate      time.Time      `json:"birth_date"`
	TaxNumber      string         `json:"tax_number"`
	Email          string         `json:"email"`
	HashedPassword string         `json:"hashed_password"`
	Phone          string         `json:"phone"`
	Address        string         `json:"address"`
	UserType       string         `json:"user_type"`
	LastCode       string         `json:"last_code"`
	Profile        *Profile       `gorm:"foreignKey:UserID"`
}

type UserCreateRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type UserCompleteRequest struct {
	TaxNumber string    `json:"tax_number"`
	Phone     string    `json:"phone"`
	Address   string    `json:"address"`
	UserType  string    `json:"user_type"`
	BirthDate time.Time `json:"birth_date"`
}

type UserUpdateRequest struct {
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	TaxNumber string    `json:"tax_number"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Phone     string    `json:"phone"`
	Address   string    `json:"address"`
	UserType  string    `json:"user_type"`
	BirthDate time.Time `json:"birth_date"`
}

type UserPageResponse struct {
	ID        uint   `json:"id"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Address   string `json:"address"`
	UserType  string `json:"usertype"`
}

type UserPageFilter struct {
	Name      string `json:"name"`
	Address   string `json:"address"`
	UserType  string `json:"usertype"`
	TaxNumber string `json:"tax_number"`
}

func (f *UserPageFilter) Apply(db *gorm.DB) *gorm.DB {
	if f.Name != "" {
		db = db.Where("first_name LIKE ? OR last_name LIKE ?", "%"+f.Name+"%", "%"+f.Name+"%")
	}
	if f.Address != "" {
		db = db.Where("address LIKE ?", "%"+f.Address+"%")
	}
	if f.UserType != "" {
		db = db.Where("user_type = ?", f.UserType)
	}
	if f.TaxNumber != "" {
		db = db.Where("tax_number LIKE ?", "%"+f.TaxNumber+"%")
	}
	return db
}
