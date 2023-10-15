package models

type Link struct {
    shortUrl     string  `json:"short_url"`
    TargetUrl  string  `json:"target_url"`
    Owner string  `json:"owner"`
    Description  string `json:"description"`
}