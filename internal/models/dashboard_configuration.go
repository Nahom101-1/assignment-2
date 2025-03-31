package models

// Features represents the optional data fields that a user wants

type Features struct {
	Temperature      bool     `json:"temperature"`      // Indicates whether temperature in degree Celsius is shown
	Precipitation    bool     `json:"precipitation"`    // Indicates whether precipitation (rain, showers and snow) is shown
	Capital          bool     `json:"capital"`          // Indicates whether the name of the capital is shown
	Coordinates      bool     `json:"coordinates"`      // Indicates whether population is shown
	Population       bool     `json:"population"`       // Indicates whether population is shown
	Area             bool     `json:"area"`             // Indicates whether land area size is shown
	TargetCurrencies []string `json:"targetCurrencies"` // Indicates which exchange rates (to target currencies) relative to the base currency of the registered country (in this case NOK for Norway) are shown
}

// DashboardConfig represents the configuration a user submits
type DashboardConfig struct {
	Country  string   `json:"country"`  // Indicates country name (alternatively to ISO code, i.e., country name can be empty if ISO code field is filled and vice versa)
	IsoCode  string   `json:"isoCode"`  // Indicates two-letter ISO code for country (alternatively to country name)
	Features Features `json:"features"` // Set of dashboard data options
}
