package octopusenergy

import "time"

// Consumption represents the consumption reported by a gas/electricity meter
// see https://developer.octopus.energy/docs/api/#consumption
type Consumption struct {
	Results []ConsumptionResult `json:"results,omitempty"`
}

// ConsumptionResult TODO
type ConsumptionResult struct {
	Consumption   float32   `json:"consumption"`
	IntervalStart time.Time `json:"interval_start"`
	IntervalEnd   time.Time `json:"interval_end"`
}
