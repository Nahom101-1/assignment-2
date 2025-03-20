package models

type Status struct {
	RestCountriesAPI string `json:"countries_api"`
	OpenMeteoAPI     string `json:"meteo_api"`
	CurrencyAPI      string `json:"currency_api"`
	NotificationDB   string `json:"notification_db"`
	Webhooks         int    `json:"webhooks"`
	Version          string `json:"version"`
	Uptime           int    `json:"uptime"`
}
