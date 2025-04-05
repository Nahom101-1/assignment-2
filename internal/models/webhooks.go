package models

// Webhook struct for Registration of a webhook
type Webhook struct {
	ID      string `json:"id"`
	URL     string `json:"url"`
	Country string `json:"country,omitempty"`
	Event   string `json:"event"`
}
