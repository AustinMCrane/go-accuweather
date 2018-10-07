package accuweather

import "time"

type TimeZone struct {
	Code             string    `json:"Code"`
	Name             string    `json:"Name"`
	GmtOffset        float32   `json:"GmtOffset"`
	IsDaylightSaving bool      `json:"IsDaylightSaving"`
	NextOffsetChange time.Time `json:"NextOffsetChange"`
}

type Measure struct {
	Value    float32 `json:"Value"`
	Unit     string  `json:"Unit"`
	UnitType int     `json:"UnitType"`
}
