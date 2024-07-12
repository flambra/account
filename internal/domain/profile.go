package domain

import (
	"time"

	"gorm.io/gorm"
)

type Profile struct {
	UserID             uint `gorm:"not null;autoIncrement:false"`
	CreatedAt          time.Time
	UpdatedAt          time.Time
	DeletedAt          gorm.DeletedAt `gorm:"index"`
	Description        string
	Skills             string
	Portfolio          string
	Specializations    string
	Availability       bool
	Languages          string
	Location           string
	PhoneType          string
	Rating             float32
	Level              int
	SuccessRate        float32
	ReturnRate         float32
	AvgDelivery        float32
	LastDelivery       time.Time
	Requests           int
	TiktokFollowers    int
	InstagramFollowers int
	YoutubeFollowers   int
}

type ProfileUpdateRequest struct {
	UserID          uint
	Description     string `json:"description"`
	Skills          string `json:"skills"`
	Portfolio       string `json:"portfolio"`
	Specializations string `json:"specializations"`
	Availability    bool   `json:"availability"`
	Languages       string `json:"languages"`
	Location        string `json:"location"`
	PhoneType       string `json:"phone_type"`
}
