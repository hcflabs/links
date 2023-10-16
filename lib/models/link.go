package models

type LinkOptions struct {
	Permanent bool `json:"permanent"`
}

type LinkRequest struct {
	LinkOptions
	ShortUrl    string `json:"short_url"`
	TargetUrl   string `json:"target_url"`
	Owner       string `json:"owner"`
	Description string `json:"description"`
}

type Link struct {
	LinkOptions
	ShortUrl          string `json:"short_url"`
	TargetUrl         string `json:"target_url"`
	Owner             string `json:"owner"`
	Description       string `json:"description"`
	CreationTimestamp int    `json:"created"`
	ModifiedTimestamp int    `json:"modified"`
}

type Error struct {
	message string
}
