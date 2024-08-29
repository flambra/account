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
	Description        string         `json:"description"`
	Skills             string         `json:"skills"`
	Portfolio          string         `json:"portfolio"`
	Specializations    string         `json:"specializations"`
	Availability       bool           `json:"availability"`
	Languages          string         `json:"languages"`
	Location           string         `json:"location"`
	PhoneType          string         `json:"phone_type"`
	Rating             float32        `json:"rating"`
	Level              int            `json:"level"`
	SuccessRate        float32        `json:"success_rate"`
	ReturnRate         float32        `json:"return_rate"`
	AvgDelivery        float32        `json:"avg_delivery"`
	LastDelivery       time.Time      `json:"last_delivery"`
	Requests           int            `json:"requests"`
	TiktokFollowers    int            `json:"tiktok_followers"`
	InstagramFollowers int            `json:"instagram_followers"`
	YoutubeFollowers   int            `json:"youtube_followers"`
}

type ProfileUpdateRequest struct {
	UserID          uint
	Description     string `json:"description,omitempty"`
	Skills          string `json:"skills,omitempty"`
	Portfolio       string `json:"portfolio,omitempty"`
	Specializations string `json:"specializations,omitempty"`
	Availability    bool   `json:"availability,omitempty"`
	Languages       string `json:"languages,omitempty"`
	Location        string `json:"location,omitempty"`
	PhoneType       string `json:"phone_type,omitempty"`
}
