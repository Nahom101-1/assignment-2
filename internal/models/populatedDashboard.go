package models

// PopulatedDashboard represents the final dashboard data returned to client
type PopulatedDashboard struct {
	Country       string            `json:"country"`
	IsoCode       string            `json:"isoCode"`
	Features      PopulatedFeatures `json:"features"`
	LastRetrieval string            `json:"lastRetrieval"`
}

// PopulatedFeatures holds all the optional features of the dashboard
type PopulatedFeatures struct {
	Temperature      *float32           `json:"temperature,omitempty"`
	Precipitation    *float32           `json:"precipitation,omitempty"`
	Capital          *string            `json:"capital,omitempty"`
	Coordinates      *Coordinates       `json:"coordinates,omitempty"`
	Population       *int32             `json:"population,omitempty"`
	Area             *float32           `json:"area,omitempty"`
	TargetCurrencies map[string]float32 `json:"targetCurrencies,omitempty"`
}

// Coordinates struct for lang and long
type Coordinates struct {
	Latitude  float32 `json:"latitude"`
	Longitude float32 `json:"longitude"`
}

// CountryCoordinatesResponse struct lang and long  data from the RestCountries API,
// also currency but only used to get ISO3
type CountryCoordinatesResponse struct {
	LatLng     []float32              `json:"latlng"`
	Currencies map[string]interface{} `json:"currencies"`
}

// WeatherData struct for weather data from the Open-Meteo API
type WeatherData struct {
	Current struct {
		Temperature2M float32 `json:"temperature_2m"`
		Precipitation float32 `json:"precipitation"`
	} `json:"current"`
}

// RatesResponse struct for currency rates
type RatesResponse struct {
	Rates map[string]float32 `json:"rates"`
}

// GeneralData Struct for capital, area and population
type GeneralData struct {
	Capital    []string `json:"capital"`
	Area       float32  `json:"area"`
	Population int32    `json:"population"`
}

// GeneralDataResponse holds capital, population, and area fields
type GeneralDataResponse struct {
	Capital    []string `json:"capital"`
	Population int32    `json:"population"`
	Area       float32  `json:"area"`
}
