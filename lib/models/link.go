package models

import (
	"gorm.io/gorm"
	"time"
)

type LinkOptions struct {
	Permanent bool `json:"permanent"`
}



type Link struct {
	gorm.Model
	LinkOptions
	ShortUrl    string    `json:"short_url" gorm:"primaryKey" gorm:"index"`
	TargetUrl   string    `json:"target_url"`
	Owner       string    `json:"owner" gorm:"index"`
	Description string    `json:"text"`
	CreatedAt   time.Time `json:"created" time_format:"RFC3339"`
	UpdatedAt   time.Time `json:"modified" time_format:"RFC3339"`
}

type Error struct {
	message string
}
