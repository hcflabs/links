package models 

type LinkRequest struct {
	LinkOptions
	ShortUrl    string `json:"short_url"`
	TargetUrl   string `json:"target_url"`
	Owner       string `json:"owner"`
	Description string `json:"description"`
}
