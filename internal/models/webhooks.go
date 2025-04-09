package models

// Webhook struct for Registration of a webhook
type Webhook struct {
	ID      string `json:"id,omitempty"`
	URL     string `json:"url"`               // URL to be invoked when event occurs
	Country string `json:"country,omitempty"` // Country that is registered, or empty if all countries
	Event   string `json:"event"`             // Event on which it is invoked

}

// Notification struct for payload for post request to webhook
type Notification struct {
	ID      string `json:"id"`
	Country string `json:"country"`
	Event   string `json:"event"`
	Time    string `json:"time"`
}
