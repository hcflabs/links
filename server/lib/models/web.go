package models

// JSON API Link Structure
type NewLink struct {
	ShortUrl  string `json:"short_url"`
	TargetUrl string `json:"target_url"`
	Permanent bool   `json:"permanent"`
	// Metadata and Control Options
	Owner       string `json:"owner"`
	Description string `json:"description"`
	Protected   bool   `json:"protected"`
	// Extra Features
	//TODO: Unimplemented
	// QueryParam []string
}
