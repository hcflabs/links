package models

type LinkRequest struct {
	// LinkOptions
	ShortUrl    string `json:"short_url"`
	TargetUrl   string `json:"target_url"`
	Owner       string `json:owned_by`
	Description string `json:"description"`
}
