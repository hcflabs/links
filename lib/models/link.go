package models

import (
	"time"
)

// API Link Structure
type ExternalLinkMetaData struct {
	Owner       string    `json:owned_by gorm:"index"`
	Description string    `json:"text"`
	CreatedAt   time.Time `json:"created" time_format:"RFC3339"`
	UpdatedAt   time.Time `json:"modified" time_format:"RFC3339"`
	Protected   bool      `json:"protected"`
}

type ExternalLink struct {
	ShortUrl  string `json:"short_url"`
	TargetUrl string `json:"target_url"`
	Permanent bool   `json:"permanent"`
	MetaData  ExternalLinkMetaData
}

// Internal Link Structure
type InternalLink struct {
	ShortUrl    string
	TargetUrl   string
	Owner       string
	Description string
	CreatedAt   time.Time ` time_format:"RFC3339"`
	UpdatedAt   time.Time ` time_format:"RFC3339"`
	Permanent   bool    
	Protected   bool     
}
