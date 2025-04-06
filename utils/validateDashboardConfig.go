package utils

import (
	"fmt"
	"github.com/Nahom101-1/assignment-2/internal/models"
)

func ValidateDashboardConfig(config models.DashboardConfig) error {
	if config.Country == "" && config.IsoCode == "" {
		return fmt.Errorf("Either Country or IsoCode must be provided")
	}
	// Validate features
	f := config.Features
	if !f.Temperature && !f.Precipitation && !f.Capital && !f.Coordinates && !f.Population && !f.Area && len(f.TargetCurrencies) == 0 {
		return fmt.Errorf("At least one feature must be selected")
	}

	return nil
}
