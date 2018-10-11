package accuweather

// DailyForecastType forecast type for daily forecasts
type DailyForecastType string

// HourlyForecastType forecast type for hourly forecasts
type HourlyForecastType string

func (f DailyForecastType) String() string {
	return string(f)
}

func (f HourlyForecastType) String() string {
	return string(f)
}

const (
	// OneDay daily forecast type for 1 days
	OneDay DailyForecastType = "1day"
	// FiveDay daily forecast type for 5 days
	FiveDay DailyForecastType = "5day"
	// TenDay daily forecast type for 10 days
	TenDay DailyForecastType = "10day"
	// FifteenDay daily forecast type for 15 days
	FifteenDay DailyForecastType = "15day"

	// OneHour hourly forecast type for 1 hour
	OneHour HourlyForecastType = "1hour"
	// TwelveHour hourly forecast type for 12 hours
	TwelveHour HourlyForecastType = "12hour"
	// TwentyFourHour hourly forecast type for 24 hours
	TwentyFourHour HourlyForecastType = "24hour"
	// SeventyTwoHour hourly forecast type for 72 hours
	SeventyTwoHour HourlyForecastType = "72hour"
	// OneHundredTwentyHour hourly forecast type for 120 hours
	OneHundredTwentyHour HourlyForecastType = "120hour"
)

// DailyForecast contains data for one day of forecast
// also is split into day and night forecast icons
type DailyForecast struct {
	Headline struct {
		EffectiveDate      string `json:"EffectiveDate"`
		EffectiveEpochDate int    `json:"EffectiveEpochDate"`
		Severity           int    `json:"Severity"`
		Text               string `json:"Text"`
		Category           string `json:"Category"`
		EndDate            string `json:"EndDate"`
		EndEpochDate       int    `json:"EndEpochDate"`
		MobileLink         string `json:"MobileLink"`
		Link               string `json:"Link"`
	} `json:"Headline"`
	DailyForecasts []struct {
		Date        string `json:"Date"`
		EpochDate   int    `json:"EpochDate"`
		Temperature struct {
			Minimum Measure `json:"Minimum"`
			Maximum Measure `json:"Maximum"`
		} `json:"Temperature"`
		Day struct {
			Icon       int    `json:"Icon"`
			IconPhrase string `json:"IconPhrase"`
		} `json:"Day"`
		Night struct {
			Icon       int    `json:"Icon"`
			IconPhrase string `json:"IconPhrase"`
		} `json:"Night"`
		Sources    []string `json:"Sources"`
		MobileLink string   `json:"MobileLink"`
		Link       string   `json:"Link"`
	} `json:"DailyForecasts"`
}

// HourlyForecast contains data for an hour of forcast
type HourlyForecast struct {
	DateTime                 string  `json:"DateTime"`
	EpochDateTime            int     `json:"EpochDateTime"`
	WeatherIcon              int     `json:"WeatherIcon"`
	IconPhrase               string  `json:"IconPhrase"`
	IsDaylight               bool    `json:"IsDaylight"`
	Temperature              Measure `json:"Temperature"`
	PrecipitationProbability int     `json:"PrecipitationProbability"`
	MobileLink               string  `json:"MobileLink"`
	Link                     string  `json:"Link"`
}
