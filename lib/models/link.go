package models

import (
	"time"
)

type LinkOptions struct {
	Permanent bool `json:"permanent"`
}

type LinkMetadata struct {
	Owner       string    `json:"owner" gorm:"index"`
	Description string    `json:"text"`
	CreatedAt   time.Time `json:"created" time_format:"RFC3339"`
	UpdatedAt   time.Time `json:"modified" time_format:"RFC3339"`
}

type Link struct {
	ShortUrl    string    `json:"short_url" gorm:"primaryKey" gorm:"index"`
	TargetUrl   string    `json:"target_url"`
    LinkMetadata LinkMetadata
    LinkOptions LinkOptions

}

type Error struct {
	message string
}
