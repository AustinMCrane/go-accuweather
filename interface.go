package accuweather

import "time"

// TimeZone is a api defined time zone structure
type TimeZone struct {
	Code             string    `json:"Code"`
	Name             string    `json:"Name"`
	GmtOffset        float32   `json:"GmtOffset"`
	IsDaylightSaving bool      `json:"IsDaylightSaving"`
	NextOffsetChange time.Time `json:"NextOffsetChange"`
}

// Measure is used to contain a value and its units of measure
type Measure struct {
	Value    float32 `json:"Value"`
	Unit     string  `json:"Unit"`
	UnitType int     `json:"UnitType"`
}
