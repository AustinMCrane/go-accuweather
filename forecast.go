package accuweather

type DailyForecastType string
type HourlyForecastType string

func (f DailyForecastType) String() string {
	return string(f)
}

func (f HourlyForecastType) String() string {
	return string(f)
}

const (
	OneDay     DailyForecastType = "1day"
	FiveDay    DailyForecastType = "5day"
	TenDay     DailyForecastType = "10day"
	FifteenDay DailyForecastType = "15day"

	OneHour              HourlyForecastType = "1hour"
	TwelveHour           HourlyForecastType = "12hour"
	TwentyFourHour       HourlyForecastType = "24hour"
	SeventyTwoHour       HourlyForecastType = "72hour"
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
