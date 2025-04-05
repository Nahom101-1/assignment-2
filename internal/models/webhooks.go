package models

// Webhook struct for Registration of a webhook
type Webhook struct {
	ID      string `json:"id"`
	URL     string `json:"url"`               // URL to be invoked when event occurs
	Country string `json:"country,omitempty"` // Country that is registered, or empty if all countries
	Event   string `json:"event"`             // Event on which it is invoked

}
