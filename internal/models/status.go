package models

type Status struct {
	RestCountriesAPI int    `json:"countries_api"`
	OpenMeteoAPI     int    `json:"meteo_api"`
	CurrencyAPI      int    `json:"currency_api"`
	NotificationDB   int    `json:"notification_db"`
	Webhooks         int    `json:"webhooks"`
	Version          string `json:"version"`
	Uptime           int    `json:"uptime"`
}
