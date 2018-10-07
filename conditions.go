package accuweather

// CurrentCondition has data for a locations temperature, weather icons, phrase and time
type CurrentCondition struct {
	LocalObservationDateTime string `json:"LocalObservationDateTime"`
	EpochTime                int    `json:"EpochTime"`
	WeatherText              string `json:"WeatherText"`
	WeatherIcon              int    `json:"WeatherIcon"`
	IsDayTime                bool   `json:"IsDayTime"`
	Temperature              struct {
		Metric   Measure `json:"Metric"`
		Imperial Measure `json:"Imperial"`
	} `json:"Temperature"`
	MobileLink string `json:"MobileLink"`
	Link       string `json:"Link"`
}
